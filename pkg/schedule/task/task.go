package task

type Result struct {
	Success bool
	Result  interface{}
	Message string
	Logs    string
}

type Type struct {
	Name           string
	Priority       int
	DefaultTrigger *Trigger
}

type Executable interface {
	Execute() *Result
}

type Trigger interface {
	Trig() bool
}

type Task struct {
	trigger Trigger
	Executable
}
