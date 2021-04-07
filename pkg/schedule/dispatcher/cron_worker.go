package dispatcher

import "github.com/robfig/cron/v3"

type CronWorker struct {
	Worker
	cron *cron.Cron
}

func (that *CronWorker) handle(task Task) bool {
	return true
}

func (that *CronWorker) Stop() {

}

func (that *CronWorker) StopNow() {

}

func (that *CronWorker) Initialize() {

}
