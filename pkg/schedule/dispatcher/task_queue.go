package dispatcher

type BlockingQueue interface {
	// 添加到尾部, 如果队列满则抛出错误
	Add(task *Task) error
	// 添加到尾部, 如果队列满则阻塞
	Put(task *Task)
	// 返回队列头部元素, 空则返回 nil
	Peek() *Task
	// 返回并移除队列头部元素, 空则返回 nil
	Poll() *Task
	// 返回并移除队列头部元素, 队列为空则阻塞
	Take() *Task
}
