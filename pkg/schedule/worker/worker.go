package worker

import "gocms/pkg/schedule/task"

type Worker struct {
	MaxTask   int
	taskQueue task.BlockingQueue
}

func (that *Worker) Handle(task *task.Task) {
	// blocking
	that.taskQueue.Put(task)
}
