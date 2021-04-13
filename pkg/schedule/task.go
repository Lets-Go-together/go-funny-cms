package schedule

import (
	"bytes"
	"encoding/json"
	"errors"
	"math"
	"time"
)

type TaskState uint8

const (
	TaskStateInitialize TaskState = iota
	TaskStateStarting
	TaskSateStopping
	TaskStateRebooting

	TaskStateRunning
	TaskStateStopped
	TaskStateDeleting
)

var taskStateStrMap = map[TaskState]string{
	TaskStateInitialize: "TaskStateInitialize",
	TaskStateStarting:   "TaskStateStarting",
	TaskSateStopping:    "TaskSateStopping",
	TaskStateRebooting:  "askStateRebooting",
	TaskStateRunning:    "TaskStateRunning",
	TaskStateStopped:    "TaskStateStopped",
	TaskStateDeleting:   "TaskStateDeleting",
}

type TaskType uint8

type TaskHandleFunc func(ctx *Context) error

type ExecuteInfo struct {
	CreateAt           int64 `json:"create_at"`
	StopAt             int64 `json:"stop_at"`
	LastExecuteAt      int64 `json:"last_executed_at"`
	LastSuccess        int64 `json:"last_success"`
	LastFail           int64 `json:"last_fail"`
	AverageSpanTimeSec int64 `json:"average_span_time_sec"`
	TotalSpanTimeSec   int64 `json:"total_span_time"`
	TotalExecute       uint  `json:"total_execute"`
	TotalSuccess       uint  `json:"total_success"`
}

func (that *ExecuteInfo) CreateNow() {
	that.CreateAt = that.getNow()
}

func (that *ExecuteInfo) StopNow() {
	that.StopAt = that.getNow()
}
func (that *ExecuteInfo) ExecuteNow() {
	that.LastExecuteAt = that.getNow()
	that.TotalExecute += 1
}

func (that *ExecuteInfo) SuccessNow() {
	that.LastSuccess = that.getNow()
	that.TotalSpanTimeSec += that.getNow() - that.LastExecuteAt
	that.TotalSuccess += 1
}

func (that *ExecuteInfo) FailNow() {
	that.LastFail = that.getNow()
	that.TotalSpanTimeSec += that.getNow() - that.LastExecuteAt
	that.TotalSuccess += 1
}

func (that *ExecuteInfo) getNow() int64 {
	return time.Now().UnixNano()
}

type Task struct {
	Id         int       `json:"id"`
	TaskId     int       `json:"task_id"`
	Name       string    `json:"name"`
	State      TaskState `json:"state"`
	Desc       string    `json:"desc"`
	CronExpr   string    `json:"cron_expr"`
	Timeout    uint16    `json:"cron_timeout"`
	RetryTimes uint8     `json:"retry_times"`
	Type       *TaskType `json:"type"`

	executeInfo *ExecuteInfo
	context     *Context
	broker      *TaskBroker
}

func NewTask(name string, desc string, cronExpr string) *Task {
	return &Task{
		Name:       name,
		Desc:       desc,
		State:      TaskStateInitialize,
		CronExpr:   cronExpr,
		Timeout:    math.MaxUint16,
		RetryTimes: 0,

		executeInfo: &ExecuteInfo{},
	}
}

func (that *Task) StateName() string {
	return taskStateStrMap[that.State]
}

func (that *Task) Context() *Context {
	return nil
}

func (that *Task) Log(log string) {

}

func (that *Task) init() {
	if that.executeInfo == nil {
		that.executeInfo = &ExecuteInfo{}
	}
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
		that.State == TaskStateInitialize ||
		that.State == TaskStateRebooting
}

type Context struct {
	Task           Task
	Retry          uint8
	RetryRemaining uint8
}

type TaskResult struct {
	Success        bool
	Message        string
	RetryRemaining uint8

	logs bytes.Buffer
}

func (that *TaskResult) Log(log string) {
	that.logs.WriteString(log)
	that.logs.WriteString("\n")
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
