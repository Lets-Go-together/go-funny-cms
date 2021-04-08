package mailer

type Task struct {
	Next    *Express
	Express *Express
}

type TaskQueue interface {
	// 添加到尾部, 如果队列满则抛出错误
	Add(task *Express) error
	// 添加到尾部, 如果队列满则阻塞
	Put(task *Express)
	// 返回队列头部元素, 空则返回 nil
	Peek() *Express
	// 返回并移除队列头部元素, 空则返回 nil
	Poll() *Express
	// 返回并移除队列头部元素, 队列为空则阻塞
	Take() *Express
}

// SendNow 立即发送
func (t *Task) SendNow() {

}

// Send 根据配置正常发送模式
func (t *Task) Send() {

}

// SendDelay 延迟发送模式
func (t *Task) SendDelay() {

}
