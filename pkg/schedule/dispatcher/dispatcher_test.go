package dispatcher

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "-"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d = New()

			task := TaskInfo{
				Name:     "task_a",
				Desc:     "test task",
				CronExpr: "",
			}

			d.AddTask(task, func() *TaskResult {
				t.Log("Task execute...")
				return nil
			})
			d.Start()
			time.Sleep(time.Hour)
		})
	}
}
