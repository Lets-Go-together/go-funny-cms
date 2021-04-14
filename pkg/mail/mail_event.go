package mail

import "fmt"

const (
	DINGTALK = 1
	WECHAT   = 2
)

// Event 邮件发送事件
type Event interface {
	Success(v interface{}) error
	Failed(v interface{}) error
}

// DingTalk 钉钉通知对应操作
type DingTalk struct{}

func (t DingTalk) Success(v interface{}) error {
	fmt.Println("钉钉通知: ", v)
	return nil
}
func (t DingTalk) Failed(v interface{}) error {
	return nil
}

// Wechat 企业微信通知对应操作
type Wechat struct{}

func (t Wechat) Success(v interface{}) error {
	fmt.Println("Wechat通知 Ssccess: ", v)
	return nil
}
func (t Wechat) Failed(v interface{}) error {
	return nil
}
