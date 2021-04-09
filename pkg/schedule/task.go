package schedule

import (
	"encoding/json"
	"errors"
)

type TaskState int

const (
	TaskStateUnknown TaskState = iota
	TaskStateStarting
	TaskSateStopping
	TaskStateRebooting
	TaskStateUpdating

	TaskStateRunning
	TaskStateStopped
	TaskStateDeleting
)

type TaskType int8

type TaskHandleFunc func(ctx *Context) *TaskResult

type Task struct {
	TaskId     int       `json:"task_id"`
	Name       string    `json:"name"`
	State      TaskState `json:"state"`
	Desc       string    `json:"desc"`
	CronExpr   string    `json:"cron_expr"`
	Timeout    int       `json:"cron_timeout"`
	RetryTimes int8      `json:"retry_times"`
	Type       *TaskType `json:"type"`

	context *Context
	broker  *TaskBroker
}

func (that *Task) Execute(context *Context) *TaskResult {
	return nil
}

func (that *Task) Context() *Context {
	return nil
}

func (that *Task) Log(tag string, log string) {

}

func (that *Task) ChangeState(state TaskState) error {
	that.State = state
	b := *that.broker
	b.UpdateTask(that)
	return nil
}

func (that *Task) String() string {
	bytes, err := json.Marshal(that)
	if err != nil {
		return "{}"
	}
	return string(bytes)
}

func (that *Task) StateInChange() bool {
	return that.State == TaskStateStarting ||
		that.State == TaskSateStopping ||
		that.State == TaskStateUnknown ||
		that.State == TaskStateRebooting
}

type Context struct {
	retryTimes int
}

type TaskResult struct {
	Success bool
	Message string
}

type TaskHandleFuncMap struct {
	funcMap map[string]TaskHandleFunc
}

func newTaskHandleFuncMap() *TaskHandleFuncMap {
	return &TaskHandleFuncMap{
		funcMap: map[string]TaskHandleFunc{},
	}
}

func (that *TaskHandleFuncMap) SetHandleFunc(task string, handleFunc TaskHandleFunc) (err error) {
	if that.funcMap[task] != nil {
		err = errors.New("type already exist")
	}
	that.funcMap[task] = handleFunc
	return
}

func (that *TaskHandleFuncMap) GetHandleFunc(task string) TaskHandleFunc {
	return that.funcMap[task]
}

type ExecuteInfo struct {
	CreateAt           uint64 `json:"create_at"`
	StopAt             uint64 `json:"stop_at"`
	LastExecutedAt     uint64 `json:"last_executed_at"`
	LastSuccess        uint64 `json:"last_success"`
	AverageSpanTimeSec uint64 `json:"average_s"`
	TotalSpanTimeSec   uint64 `json:"total_span_time"`
	TotalExecute       int    `json:"total_execute"`
	TotalSuccess       int    `json:"total_success"`
}
