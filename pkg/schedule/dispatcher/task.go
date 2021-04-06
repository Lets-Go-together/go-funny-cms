package dispatcher

import "errors"

//goland:noinspection GoUnusedGlobalVariable
var (
	TaskTypeDefault = &TaskType{
		Name:     "default-type",
		Priority: -1,
	}

	TaskTypeCron = &TaskType{
		Name:     "cron-type",
		Priority: -1,
	}
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

type TaskType struct {
	Name      string `json:"name"`
	Priority  int    `json:"priority"`
	TriggerId int16  `json:"trigger_id"`
}

type TaskTypeExecuteFuncMap struct {
	funcMap map[string]TaskHandleFunc
}

func (that *TaskTypeExecuteFuncMap) PutUnique(typeName string, handleFunc TaskHandleFunc) (err error) {
	if that.funcMap[typeName] != nil {
		err = errors.New("type already exist")
	}
	that.funcMap[typeName] = handleFunc
	return
}

func (that TaskTypeExecuteFuncMap) Get(typeName string) TaskHandleFunc {
	return that.funcMap[typeName]
}

type Trigger interface {
	Trigger() bool
}

type Task interface {
	Execute() *TaskResult
	Entity() *TaskEntity
}

type TaskEntity struct {
	Name        string    `json:"name"`
	TaskId      uint64    `json:"task_id"`
	State       int8      `json:"state"`
	Type        *TaskType `json:"type"`
	ExecuteFunc TaskHandleFunc
}

func (that *TaskEntity) Execute() *TaskResult {
	return nil
}
