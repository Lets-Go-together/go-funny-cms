package schedule

import (
	"fmt"
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
				t.Log("Task execute...", context.Task.Name, strconv.Itoa(context.Task.Id))
				// do something
				return nil
			})
			scheduler.Launch()

			// 新建任务
			//task := NewTask("email", "test task", "* * * * * ?")
			// 添加任务并立即执行
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
func TestRedisTaskBroker_QueryTaskByState(t *testing.T) {
	scheduler := New()
	scheduler.RegisterTask(".", func(ctx *Context) error {
		t.Log("Execute: " + ctx.Task.Name)
		return nil
	})
	scheduler.Launch()
	time.Sleep(time.Second * 3)
	s := map[string][]*Task{}
	s["running"] = scheduler.broker.QueryTaskByState(TaskStateRunning)
	s["stopped"] = scheduler.broker.QueryTaskByState(TaskStateStopped)
	s["starting"] = scheduler.broker.QueryTaskByState(TaskStateStarting)
	s["stopping"] = scheduler.broker.QueryTaskByState(TaskStateStopping)
	for s2, tasks := range s {
		t.Log("state:", s2, ",count:", len(tasks))
		for _, task := range tasks {
			t.Log("\t", task.Name)
		}
	}
}

func TestScheduler_StopTask(t *testing.T) {
	scheduler := New()
	scheduler.RegisterTask(".", func(ctx *Context) error {
		t.Log("Execute: " + ctx.Task.Name)
		return nil
	})
	scheduler.Launch()
	tasks := scheduler.QueryAllTask()
	for _, task := range tasks {
		if task.State == TaskStateStopped {
			continue
		}
		t.Log("Stopping task: ", task)
		scheduler.StopTask(task.Id)
		break
	}
	time.Sleep(time.Hour)
}

func TestScheduler_StartTask(t *testing.T) {
	scheduler := New()
	scheduler.RegisterTask(".", func(ctx *Context) error {
		t.Log("Execute: " + ctx.Task.Name)
		return nil
	})
	scheduler.Launch()
	time.Sleep(time.Second * 3)
	tasks := scheduler.QueryAllTask()
	for _, task := range tasks {
		if task.State == TaskStateRunning {
			continue
		}
		t.Log("Starting task: ", task)
		scheduler.StartTask(task.Id)
		break
	}
	time.Sleep(time.Hour)
}

func TestScheduler_AddTask(t *testing.T) {

	n := "test_task_" + time.Now().Format("01-02_15_04_05")
	scheduler := New()
	scheduler.RegisterTask(".", func(ctx *Context) error {
		t.Log("execute: " + ctx.Task.Name)
		return nil
	})
	scheduler.RegisterTask(n, func(ctx *Context) error {
		t.Log(fmt.Sprintf("Task %s(id=%d) executing.", ctx.Task.Name, ctx.Task.Id))
		return nil
	})
	scheduler.Launch()
	task, err := scheduler.AddTask(NewTask(n, "test task added by unit test.", "* * * * * ?"))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(fmt.Sprintf("Task %s added, id=%d.", task.Name, task.Id))
	}
	time.Sleep(time.Hour)
}

func TestScheduler_QueryAllTask(t *testing.T) {

	scheduler := New()
	scheduler.RegisterTask("email", func(ctx *Context) error {
		return nil
	})
	scheduler.Launch()
	tasks := scheduler.QueryAllTask()
	t.Log("All Task:")
	for i := range tasks {
		t.Log(tasks[i])
	}
}
