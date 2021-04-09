package schedule

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

			// 初始化
			var scheduler = New()
			scheduler.Launch()

			// 注册任务
			scheduler.RegisterTask("email", func(context *Context) *TaskResult {
				t.Log("Task execute...")
				// do something
				return &TaskResult{Success: true, Message: ""}
			})

			// 新建任务
			task := Task{
				Name:     "email", // 固定的, 必须注册过
				Desc:     "test task",
				CronExpr: "* * * * * ?",
			}
			// 查询任务
			scheduler.AddTask(&task)
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
