package schedule

import (
	"github.com/robfig/cron/v3"
	"gocms/pkg/schedule/log"
)

type CronWorker struct {
	cron           *cron.Cron
	tasks          map[cron.EntryID]*Task
	taskHandleFunc *TaskHandleFuncMap
}

func (that *CronWorker) Process(task *Task) error {
	log.D("worker", "process task: "+task.String())
	handleFunc := that.taskHandleFunc.GetHandleFunc(task.Name)
	entryId, err := that.cron.AddFunc(task.CronExpr, func() {
		handleFunc(nil)
	})
	if err != nil {
		return err
	}
	//that.tasks[entryId] = task
	task.TaskId = int(entryId)
	err = task.ChangeState(TaskStateRunning)
	return err
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
	that.tasks = map[cron.EntryID]*Task{}
	go that.cron.Run()
}
