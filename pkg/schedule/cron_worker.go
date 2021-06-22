package schedule

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"gocms/pkg/schedule/log"
)

// CronWorker 为 cron.Cron 现实的一个执行器
type CronWorker struct {
	cron           *cron.Cron
	tasks          map[int]*cron.EntryID
	taskHandleFunc *TaskHandleFuncMap
}

func (that *CronWorker) Process(task *Task) error {

	log.D("worker/Process", "process task: "+task.String())

	if task.NeedStart() {
		// 当任务状态为更变为执行时
		return that.startTask(task)

	} else if task.NeedStop() {
		// 需要停止任务
		return that.removeTask(task)
	}
	return nil
}

func (that *CronWorker) startTask(task *Task) error {
	handleFunc := that.taskHandleFunc.GetHandleFunc(task.Name)
	if handleFunc == nil || len(handleFunc) == 0 {
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
		var err error
		for _, taskHandleFunc := range handleFunc {
			err = taskHandleFunc(ctx)
			if err != nil {
				break
			}
		}
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

func (that *CronWorker) removeTask(task *Task) (err error) {

	entryId := that.tasks[task.Id]
	if entryId != nil {
		that.cron.Remove(*entryId)
		log.D("worker/removeTask", "task stop success, name:", task.Name, ", id:", task.Id)
	} else if task.State != TaskStateDeleting {
		err = errors.New("task not in cron: " + task.String())
		return
	}
	if task.State == TaskStateDeleting {
		//err = task.ChangeState(TaskStateDeleted)
		err = task.Delete()
	} else {
		err = task.ChangeState(TaskStateStopped)
		task.executeInfo.StopNow()
	}
	return
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
