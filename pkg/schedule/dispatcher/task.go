package dispatcher

var (
	TaskTypeDefault = &TaskType{
		Name:           "",
		Priority:       0,
		DefaultTrigger: nil,
	}
)

type TaskResult struct {
	Success bool
	Result  interface{}
	Message string
	Logs    string
}

type TaskType struct {
	Name           string
	Priority       int
	DefaultTrigger *Trigger
}

type Trigger interface {
	Trig() bool
}

type Task interface {
	Execute() *TaskResult
}
