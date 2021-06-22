package validates

import (
	"gocms/app/validates/validate"
	"gocms/pkg/logger"
	"gocms/wrap"
)

type AddTaskParams struct {
	TaskName string `json:"task_name" validate:"required"`
	Desc     string `json:"desc" validate:"required"`
	CronExpr string `json:"cron_expr" validate:"required"`
}

func (that *AddTaskParams) Validate(ctx *wrap.ContextWrapper) bool {
	err := ctx.ShouldBind(that)
	if err != nil {
		logger.PanicError(err, "添加任务参数验证", false)
		return false
	}
	return validate.WithDefaultResponse(that, ctx)
}

type UpdateTaskParams struct {
	Id       int    `json:"id" validate:"required"`
	TaskName string `json:"task_name"`
	Desc     string `json:"desc"`
	CronExpr string `json:"cron_expr"`
	// 1, 启动
	// 2, 停止
	// 3, 重启
	State int `json:"state"`
}

func (that *UpdateTaskParams) Validate(ctx *wrap.ContextWrapper) bool {
	err := ctx.ShouldBind(that)
	if err != nil {
		logger.PanicError(err, "更新任务参数验证", false)
		return false
	}
	return validate.WithDefaultResponse(that, ctx)
}
