package schedule

import (
	"strconv"
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

			// 初始化
			var scheduler = New()
			// 注册任务
			scheduler.RegisterTask("email", func(context *Context) error {
				t.Log("Task execute..." + strconv.Itoa(context.Task.Id))
				// do something
				return nil
			})
			scheduler.Launch()

			// 新建任务
			//task := NewTask("email", "test task", "* * * * * ?")
			// 查询任务
			//task = scheduler.AddTask(task)
			//t.Log("Task added:" + strconv.Itoa(task.Id))
			//scheduler.QueryAllTask()
			//scheduler.QueryTaskById(1)
			//scheduler.StopTask(1)
			//scheduler.StartTask(1)

			time.Sleep(time.Hour)
		})
	}
}

func TestRedisTaskSource_Initialize(t *testing.T) {

}
