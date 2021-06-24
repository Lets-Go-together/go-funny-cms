package schedule

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gocms/pkg/schedule/log"
	"math"
	"regexp"
	"time"
)

type TaskState uint8

// 所有的任务状态
// 前五种状态为过渡状态, 表示正在从一个状态往另一个状态迁移, 正在等待 TaskBroker 分发.
const (
	none TaskState = iota
	// 1 新任务加入队列, 等待执行中
	TaskStateInitialize
	// 2 任务启动中
	TaskStateStarting
	// 3 任务停止中
	TaskStateStopping
	// 4 任务重启中
	TaskStateRebooting
	// 5 任务正在删除
	TaskStateDeleting
	// 6 已删除
	TaskStateDeleted

	// 7
	TaskStateRunning
	// 8
	TaskStateStopped
)

func (that TaskState) String() string {
	return taskStateStrMap[that]
}

var taskStateStrMap = map[TaskState]string{
	TaskStateInitialize: "TaskStateInitialize",
	TaskStateStarting:   "TaskStateStarting",
	TaskStateStopping:   "TaskStateStopping",
	TaskStateRebooting:  "TaskStateRebooting",
	TaskStateRunning:    "TaskStateRunning",
	TaskStateStopped:    "TaskStateStopped",
	TaskStateDeleting:   "TaskStateDeleting",
	TaskStateDeleted:    "TaskStateDeleted",
}

// 暂时无用
type TaskType uint8

// 任务处理函数, 在这个函数里执行业务代码, 执行任务时会找到该任务名称对应的执行函数并调用.
type TaskHandleFunc func(ctx *Context) error

// ExecuteInfo 表示任务的执行信息, 每次执行任务时都会更新
type ExecuteInfo struct {
	CreateAt           int64 `json:"create_at"`
	StopAt             int64 `json:"stop_at"`
	LastExecuteAt      int64 `json:"last_executed_at"`
	LastSuccess        int64 `json:"last_success"`
	LastFail           int64 `json:"last_fail"`
	AverageSpanTimeSec int64 `json:"average_span_time_sec"`
	TotalSpanTimeSec   int64 `json:"total_span_time"`
	// 总共执行次数
	TotalExecute uint `json:"total_execute"`
	// 总共失败次数
	TotalSuccess uint `json:"total_success"`
}

// 更新任务的创建时间
func (that *ExecuteInfo) CreateNow() {
	that.CreateAt = that.getNow()
}

// 更新任务停止时间
func (that *ExecuteInfo) StopNow() {
	that.StopAt = that.getNow()
}

// 更新任务执行时间
func (that *ExecuteInfo) ExecuteNow() {
	that.LastExecuteAt = that.getNow()
	that.TotalExecute += 1
}

// 更新任务成功时间
func (that *ExecuteInfo) SuccessNow() {
	that.LastSuccess = that.getNow()
	that.TotalSpanTimeSec += that.getNow() - that.LastExecuteAt
	that.TotalSuccess += 1
}

// 更新任务失败时间
func (that *ExecuteInfo) FailNow() {
	that.LastFail = that.getNow()
	that.TotalSpanTimeSec += that.getNow() - that.LastExecuteAt
	that.TotalSuccess += 1
}

func (that *ExecuteInfo) getNow() int64 {
	return time.Now().UnixNano()
}

func (that *ExecuteInfo) Update() {

}

// Task 表示一个具体的任务
type Task struct {
	// 任务 Id 由调度器决定, 全局唯一
	Id int `json:"id"`
	// TaskId 对应于 Worker 执行的 id, 该 Id 用于从 Worker 停止启动任务
	TaskId int `json:"task_id"`
	// 任务名称, 需要在 Scheduler 中注册过该任务, 否则无法执行
	Name string `json:"name"`
	// 任务状态, 新任务为 TaskStateInitialize
	State TaskState `json:"state"`
	// 任务描述
	Desc string `json:"desc"`
	// 任务执行 cron 表达式
	CronExpr string `json:"cron_expr"`
	// 任务执行超时时长, 超过该时长则直接中断这次任务并断定执行结果为失败.
	Timeout uint16 `json:"cron_timeout"`
	// 任务失败后重试次数
	RetryTimes uint8 `json:"retry_times"`
	// 暂时无用
	Type *TaskType `json:"type"`

	// 任务执行信息
	ExecuteInfo *ExecuteInfo
	// 任务执行时的上下文
	context *Context
	// 创建该任务的中间人 TaskBroker
	broker *TaskBroker
}

// 新建一个任务
func NewTask(name string, desc string, cronExpr string) *Task {
	return &Task{
		Name:       name,
		Desc:       desc,
		State:      TaskStateStopped,
		CronExpr:   cronExpr,
		Timeout:    math.MaxUint16,
		RetryTimes: 0,

		ExecuteInfo: &ExecuteInfo{},
	}
}

func (that *Task) Context() *Context {
	return nil
}

func (that *Task) Log(log string) {

}

// 初始化任务, 反序列化任务时需要调用改函数.
func (that *Task) init() {
	if that.ExecuteInfo == nil {
		that.ExecuteInfo = &ExecuteInfo{}
	}
}

// 改变任务状态, 并通知分发该任务的中间人 TaskBroker 更新到数据源.
func (that *Task) ChangeState(state TaskState) error {
	log.D("task/ChangeState", "task:", that.Name, ",state:", state)
	that.State = state
	b := *that.broker
	return b.UpdateTask(that)
}

func (that *Task) Delete() error {
	log.D("task/Delete", "task:", that.Name, ",id:", that.Id)
	b := *that.broker
	return b.DeleteTask(that.Id)
}

// Deprecated
func (that *Task) Update() error {
	log.D("task/Update", "task:", that.Name, ",id:", that.Id)
	b := *that.broker
	return b.UpdateTask(that)
}

func (that *Task) String() string {
	bs, err := json.Marshal(that)
	if err != nil {
		return "{}"
	}
	return string(bs)
}

// 获取任务是否正在改变状态.
func (that *Task) StateInChange() bool {
	return that.State == TaskStateStarting ||
		that.State == TaskStateStopping ||
		that.State == TaskStateInitialize ||
		that.State == TaskStateRebooting ||
		that.State == TaskStateDeleting
}

func (that *Task) NeedStart() bool {
	return that.State == TaskStateStarting || that.State == TaskStateRunning
}

func (that *Task) NeedStop() bool {
	return that.State == TaskStateStopping || that.State == TaskStateDeleting || that.State == TaskStateRebooting
}

// Context 表示执行一次任务执行上下文信息, 主要在 TaskHandleFunc 中使用
type Context struct {
	Task Task
	// 任务重试次数
	Retry uint8
	// 任务剩余重试次数
	RetryRemaining uint8
}

// TaskResult 表示执行一次任务的执行结果.
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

// TaskHandleFuncMap 表示任务名称对应的处理函数 TaskHandleFunc, 包装了一个 map
type TaskHandleFuncMap struct {
	allPattern map[string]int
	regMap     map[*regexp.Regexp]TaskHandleFunc
}

func newTaskHandleFuncMap() *TaskHandleFuncMap {
	return &TaskHandleFuncMap{
		allPattern: map[string]int{},
		regMap:     map[*regexp.Regexp]TaskHandleFunc{},
	}
}

func (that *TaskHandleFuncMap) SetHandleFunc(pattern string, handleFunc TaskHandleFunc) (err error) {
	if that.allPattern[pattern] != 0 {
		err = errors.New(fmt.Sprintf("task handler for %s already exist", pattern))
	}
	var re *regexp.Regexp
	re, err = regexp.Compile(pattern)
	if err != nil {
		return
	}
	that.regMap[re] = handleFunc
	that.allPattern[pattern] = 1

	//hf := that.funcMap[pattern]
	//that.funcMap[pattern] = append(hf, handleFunc)
	return
}

// taskName 任务名称
// 返回所有与该名称匹配的 TaskHandleFunc
func (that *TaskHandleFuncMap) GetHandleFunc(taskName string) []TaskHandleFunc {

	// TODO 缓存一下 name 对应的 handleFunc 避免每次匹配
	var res []TaskHandleFunc
	for r, f := range that.regMap {
		if r.MatchString(taskName) {
			res = append(res, f)
		}
	}
	return res
}
