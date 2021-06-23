package schedule

// Worker 表示一个任务执行者, 负责执行分派的任务
//
// Worker 负责以下事务的处理:
// 1, Task 的执行管理(任务管理, 执行表达式解析及执行)
// 2, Task 停止与启动管理
// 3, Task 的重试, 超时等执行过程管理
// 4, Task 执行日志输出
type Worker interface {
	// 处理任务, 该任务可能已在执行, 状态更变则根据具体情况停止启动任务
	Process(task *Task) error
	// 开始执行任务
	Start()
	// 停止执行所有任务
	Stop()
	// 进行初始化工作
	Launch(handleFunMap *TaskHandleFuncMap)
}
