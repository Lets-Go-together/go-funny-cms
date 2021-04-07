package dispatcher

import "errors"

type TaskState int

const (
	TaskStateUnknown TaskState = iota
	TaskStateStarting
	TaskStateRunning
	TaskSateStopping
	TaskStateStopped
	TaskStateRebooting
	TaskStateDeleting
)

type TaskHandleFunc func() *TaskResult

type TaskParam struct {
}

type TaskResult struct {
	Success bool
	Result  interface{}
	Message string
	Logs    string
}

type TaskType int

type TaskManager struct {
	funcMap map[TaskInfo]*TaskHandleFunc
	nameMap map[string]*TaskInfo
}

func (that *TaskManager) SetHandleFunc(task TaskInfo, handleFunc *TaskHandleFunc) (err error) {
	if that.funcMap[task] != nil {
		err = errors.New("type already exist")
	}
	that.funcMap[task] = handleFunc
	that.nameMap[task.Name] = &task
	return
}

func (that *TaskManager) GetTaskHandleFunc(task TaskInfo) *TaskHandleFunc {
	return that.funcMap[task]
}

func (that *TaskManager) GetInfoByName(taskName string) *TaskInfo {
	return that.nameMap[taskName]
}

type Trigger interface {
	Trigger() bool
}

type Task interface {
	Execute() *TaskResult
	Entity() *TaskInfo
}

type TaskInfo struct {
	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	TaskId   int       `json:"task_id"`
	State    TaskState `json:"state"`
	CronExpr string    `json:"cron_expr"`
	Type     *TaskType `json:"type"`
}

func (that *TaskInfo) StateInChange() bool {
	return that.State == TaskStateRunning || that.State == TaskSateStopping || that.State == TaskStateRebooting
}
