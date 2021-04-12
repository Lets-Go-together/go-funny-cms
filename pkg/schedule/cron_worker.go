package schedule

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"gocms/pkg/schedule/log"
)

type CronWorker struct {
	cron           *cron.Cron
	tasks          map[int]*cron.EntryID
	taskHandleFunc *TaskHandleFuncMap
}

func (that *CronWorker) Process(task *Task) error {

	log.D("worker", "process task: "+task.String())
	if task.State == TaskStateStarting || task.State == TaskStateRebooting || task.State == TaskStateInitialize {
		return that.startTask(task)
	} else if task.State == TaskSateStopping || task.State == TaskStateDeleting {
		return that.stopTask(task)
	}
	return nil
}

func (that *CronWorker) startTask(task *Task) error {
	handleFunc := that.taskHandleFunc.GetHandleFunc(task.Name)
	if handleFunc == nil {
		return errors.New(fmt.Sprintf("no TaskHandleFunc for task '%s'", task.Name))
	}

	entryId, err := that.cron.AddFunc(task.CronExpr, func() {
		task.executeInfo.ExecuteNow()
		retry := uint8(0)
		ctx := &Context{
			Task: *task,
		}
	Retry:
		retry++
		err := handleFunc(ctx)
		ctx.Retry = retry
		ctx.RetryRemaining = task.RetryTimes - retry
		if err == nil {
			task.executeInfo.SuccessNow()
		} else {
			if ctx.RetryRemaining > 0 {
				goto Retry
			}
			task.executeInfo.FailNow()
		}
		task.executeInfo.TotalExecute++
	})
	if err != nil {
		return err
	}

	that.tasks[task.Id] = &entryId
	//task.TaskId = int(entryId)
	err = task.ChangeState(TaskStateRunning)
	return err
}

func (that *CronWorker) stopTask(task *Task) error {
	entryId := that.tasks[task.Id]
	if entryId == nil {
		return errors.New("task not in cron: " + task.String())
	}
	that.cron.Remove(*entryId)
	return nil
}

func (that *CronWorker) Start() {
	that.cron.Run()
}

func (that *CronWorker) Stop() {
	that.cron.Stop()
}

func (that *CronWorker) Initialize(funcMap *TaskHandleFuncMap) {
	that.taskHandleFunc = funcMap
	that.cron = cron.New(cron.WithSeconds())
	that.tasks = map[int]*cron.EntryID{}
	go that.cron.Run()
}
