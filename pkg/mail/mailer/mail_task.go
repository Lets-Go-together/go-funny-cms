package mailer

type Task struct {
	Next    *Express
	Express *Express
}

// SendNow 立即发送
func (t *Task) SendNow() {

}

// Send 根据配置正常发送模式
func (t *Task) Send() {
	t.Express.Send()
}

// SendDelay 延迟发送模式
func (t *Task) SendDelay() {

}
