package dispatcher

import (
	"errors"
	"github.com/robfig/cron/v3"
)

type CronWorker struct {
	cron  *cron.Cron
	tasks map[cron.EntryID]*CronTask
}

func (that *CronWorker) NewTask(task Task) error {
	value, ok := task.(*CronTask)
	if !ok {
		return errors.New("tpe assertion failed")
	}
	entryId, err := that.cron.AddFunc(value.CronExpr, func() {
		task.Execute()
	})
	that.tasks[entryId] = value
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

func (that *CronWorker) Initialize() {
	that.cron = cron.New()
}
