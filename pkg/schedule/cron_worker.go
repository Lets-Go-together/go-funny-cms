package schedule

import (
	"errors"
	"github.com/robfig/cron/v3"
)

type CronWorker struct {
	cron           *cron.Cron
	tasks          map[cron.EntryID]*CronTask
	taskHandleFunc *TaskHandleFuncMap
}

func (that *CronWorker) NewTask(task Task) error {
	value, ok := task.(*CronTask)
	if !ok {
		return errors.New("tpe assertion failed")
	}
	handleFunc := that.taskHandleFunc.GetHandleFunc(task.GetInfo().Name)
	entryId, err := that.cron.AddFunc(value.CronExpr, func() {
		handleFunc(nil)
	})
	if err != nil {
		return err
	}
	that.tasks[entryId] = value
	err = task.ChangeState(TaskStateRunning)
	return err
}

func (that *CronWorker) Start() {
	that.cron.Run()
}

func (that *CronWorker) Stop() {
	that.cron.Stop()
}

func (that *CronWorker) StopNow() {

}

func (that *CronWorker) Initialize(funcMap *TaskHandleFuncMap) {
	that.taskHandleFunc = funcMap
	that.cron = cron.New()
	that.tasks = map[cron.EntryID]*CronTask{}
}
