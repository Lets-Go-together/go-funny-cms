package controllers

import (
	"errors"
	"fmt"
	"gocms/app/http/admin/validates"
	"gocms/pkg/config"
	"gocms/pkg/response"
	"gocms/pkg/schedule"
	"gocms/wrap"
	"strconv"
)

type SchedulerController struct {
}

func init() {
	config.Scheduler.RegisterTask("test_fail", TestTask)
}

func TestTask(ctx *schedule.Context) error {
	fmt.Println("task execute")
	return errors.New("FAILED")
}

// query param `state` see schedule.TaskState
func (*SchedulerController) List(c *wrap.ContextWrapper) {

	state, err := strconv.Atoi(c.DefaultQuery("state", "0"))
	if err != nil {
		response.ErrorResponse(102, "参数 state 错误").WriteTo(c)
		return
	}

	var tasks []*schedule.Task
	if state != 0 {
		s := schedule.TaskState(uint8(state))
		tasks = config.Scheduler.QueryTaskByState(s)
	} else {
		tasks = config.Scheduler.QueryAllTask()
	}

	if tasks == nil {
		response.SuccessResponse([0]schedule.Task{}).WriteTo(c)
		return
	}
	response.SuccessResponse(tasks).WriteTo(c)
}

func (*SchedulerController) Add(c *wrap.ContextWrapper, param *validates.AddTaskParams) {
	task := schedule.NewTask(param.TaskName, param.Desc, param.CronExpr)
	task.RetryTimes = param.Retry
	t, err := config.Scheduler.AddTask(task)
	if err != nil {
		response.ErrorResponse(100, err.Error()).WriteTo(c)
	} else {
		response.SuccessResponse(t).WriteTo(c)
	}
}

func (*SchedulerController) Delete(c *wrap.ContextWrapper) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.ErrorResponse(102, "参数 id 缺失").WriteTo(c)
		return
	}
	task := config.Scheduler.QueryTaskById(id)
	if task == nil {
		response.ErrorResponse(103, "任务不存在").WriteTo(c)
		return
	} else if task.State == schedule.TaskStateDeleting {
		response.ErrorResponse(104, "正在删除任务").WriteTo(c)
		return
	}

	err = config.Scheduler.DeleteTask(id)
	if err != nil {
		response.ErrorResponse(105, err.Error()).WriteTo(c)
	} else {
		response.SuccessResponse(nil).WriteTo(c)
	}
}

func (*SchedulerController) Update(c *wrap.ContextWrapper, params *validates.UpdateTaskParams) {

	task := config.Scheduler.QueryTaskById(params.Id)
	if task == nil {
		response.ErrorResponse(106, "任务不存在").WriteTo(c)
		return
	}

	rtMap := map[int]map[schedule.TaskState]string{
		1: {schedule.TaskStateStarting: "任务正在启动中"},
		2: {schedule.TaskStateStopping: "任务正在停止中"},
		3: {schedule.TaskStateRebooting: "任务正在重启中"},
	}

	if rtMap[params.State] != nil {
		for newState, errorMsg := range rtMap[params.State] {
			if task.State == newState {
				response.ErrorResponse(107, errorMsg)
				return
			}
			task.State = newState
			break
		}
	}

	if params.TaskName != "" {
		task.Name = params.TaskName
	}
	if params.CronExpr != "" {
		task.CronExpr = params.CronExpr
	}
	if params.Desc != "" {
		task.Desc = params.Desc
	}
	err := config.Scheduler.UpdateTask(task)
	if err != nil {
		response.ErrorResponse(107, err.Error()).WriteTo(c)
	} else {
		response.SuccessResponse(nil).WriteTo(c)
	}
}

func (*SchedulerController) Log(c *wrap.ContextWrapper) {

}
