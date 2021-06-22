package schedule

type TaskProcessor func([]*Task)

// TaskBroker 表示任务存储源与任务执行队列的中间人, 任务执行者从这订阅任务动态和更新查询任务.
//
// TaskBroker 负责以下内容:
// - 1, 如何读取 Task (读取任务的频率, 方法), 从何处读取 Task(Redis, MySQL ..)
// - 2, 如何通知 TaskProcessor 更新 Task (频率, 发布哪些任务, 数量等)
// - 3, 如何查询, 更新 Task
// - 4, 恢复任务 (恢复哪些任务, 如何恢复任务, 恢复后如何发布)
type TaskBroker interface {
	// 运行 TaskBroker, 做一些初始化工作
	Launch()
	// 从数据源读取所有任务, 筛选出需要执行的任务并分发给 TaskProcessor
	RestoreTask()
	// 开始订阅任务清单更新
	StartConsuming(tasks TaskProcessor)
	// 更新任务
	UpdateTask(info *Task) error
	// 停止指定 id 的任务
	StopTask(id int) error
	// 启动指定 id 的任务
	StartTask(id int) error
	// 添加任务到任务清单
	AddTask(info *Task) (*Task, error)

	DeleteTask(id int) error
	// 查询所有任务
	QueryAllTask() []*Task
	// 查询所有指定状态的任务
	QueryTaskByState(state TaskState) []*Task
	// 查询指定名称的任务
	QueryTaskByName(name string) []*Task
	// 查询指定 id 的任务
	QueryTaskById(taskId int) *Task
}
