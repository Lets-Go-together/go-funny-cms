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
	idEntryIdMap   map[int]*cron.EntryID
	taskHandleFunc *TaskHandleFuncMap
	idTaskMap      map[int]*Task
}

func (that *CronWorker) Process(task *Task) error {

	log.D("worker/Process", "process task: "+task.String())

	if task.NeedStart() {
		return that.startTask(task)

	} else if task.NeedStop() {
		if err := that.removeTask(task); err != nil {
			return err
		}
	} else {
		// task 状态不变, 信息更新的情况
		that.idTaskMap[task.Id] = task
	}
	return nil
}

func (that *CronWorker) startTask(task *Task) error {
	handleFunc := that.taskHandleFunc.GetHandleFunc(task.Name)
	if handleFunc == nil || len(handleFunc) == 0 {
		return errors.New(fmt.Sprintf("no TaskHandleFunc for task '%s'", task.Name))
	}

	entryId, err := that.cron.AddFunc(task.CronExpr, that.getCmd(task.Id))
	if err != nil {
		return err
	}

	that.idTaskMap[task.Id] = task
	that.idEntryIdMap[task.Id] = &entryId
	//task.TaskId = int(entryId)
	err = task.ChangeState(TaskStateRunning)
	return err
}

func (that *CronWorker) getCmd(id int) func() {

	return func() {
		// 运行中的 task 可能会更新, 每次运行都要重新获取一下
		// TODO 2021年6月24日16:01:33 优化一下更新 task 时 worker 的同步方式
		task := that.idTaskMap[id]
		handleFunc := that.taskHandleFunc.GetHandleFunc(task.Name)

		retry := 0
		ctx := &Context{
			Task: *task,
		}
	Retry:
		retry++
		var err error
		task.ExecuteInfo.ExecuteNow()
		for _, taskHandleFunc := range handleFunc {
			err = taskHandleFunc(ctx)
			if err != nil {
				break
			}
		}
		ctx.Retry = retry
		ctx.RetryRemaining = task.RetryTimes - retry
		if err == nil {
			task.ExecuteInfo.SuccessNow()
		} else {
			task.ExecuteInfo.FailNow()
			if ctx.RetryRemaining > 0 {
				log.D("Worker/getCmd", "Task ", task.Name, "execute failed. retrying, remain:", ctx.RetryRemaining, "reason:", err.Error())
				goto Retry
			}
		}
		task.ExecuteInfo.Update()
	}
}

func (that *CronWorker) removeTask(task *Task) (err error) {

	entryId := that.idEntryIdMap[task.Id]
	if entryId != nil {
		that.cron.Remove(*entryId)
		delete(that.idEntryIdMap, task.Id)
		delete(that.idTaskMap, task.Id)
		task.ExecuteInfo.StopNow()
		log.D("worker/removeTask", "task stop success, name:", task.Name, ", id:", task.Id)
	} else if task.State != TaskStateDeleting && task.State != TaskStateRebooting {
		err = errors.New("task not in cron: " + task.String())
		return
	}
	if task.State == TaskStateDeleting {
		//err = task.ChangeState(TaskStateDeleted)
		err = task.Delete()
	} else if task.State == TaskStateRebooting {
		err = task.ChangeState(TaskStateStarting)
	} else {
		err = task.ChangeState(TaskStateStopped)
	}
	return
}

func (that *CronWorker) Start() {
	that.cron.Run()
}

func (that *CronWorker) Stop() {
	that.cron.Stop()
}

func (that *CronWorker) Launch(funcMap *TaskHandleFuncMap) {
	that.taskHandleFunc = funcMap
	that.cron = cron.New(cron.WithSeconds())
	that.idEntryIdMap = map[int]*cron.EntryID{}
	that.idTaskMap = map[int]*Task{}
	go that.cron.Run()
}
