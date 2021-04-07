package mailer

import (
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"gocms/pkg/logger"
	"io/ioutil"
	"net/smtp"
	"time"
)

type Express struct {
	// https://pkg.go.dev/github.com/jordan-wright/email
	Mail    *email.Email
	Mailer  *Mailer
	Event   Event
	Options Options
}

// Options 配置选项
type Options struct {
	LoggerFile string
	NotifeType int
	Delay      time.Duration
}

// Init 初始化express
func (that *Express) GetExpress(options *Options) *Express {
	_ = that.defaultLoggerFile(options.LoggerFile)
	_ = that.defaultEvent()
	return that
}

// Send 发送操作
func (that *Express) Send() error {
	addr := fmt.Sprintf("%s:%s", that.Mailer.Host, that.Mailer.Port)
	return that.Mail.Send(addr, smtp.PlainAuth("", that.Mailer.Username, that.Mailer.Password, that.Mailer.Host))
}

// defaultLoggerFile 默认日志文件
func (that *Express) defaultLoggerFile(file string) error {
	that.Options.LoggerFile = "./storage/mail/system.log"
	return nil
}

// defaultEvent 默认event通知
func (that *Express) defaultEvent() error {
	if that.Options.NotifeType == DINGTALK {
		e := DingTalk{}
		that.Event = e
	}
	var e Event
	switch that.Options.NotifeType {
	case DINGTALK:
		e = &DingTalk{}
	case WECHAT:
		e = &Wechat{}
	default:
		return errors.New("Event 不存在")
	}

	that.SetEvent(e)
	return nil
}

// checkIsFile 默认日志文件
func (that *Express) checkIsFile() error {
	return nil
}

// setEvent 支持重置Event
func (that *Express) SetEvent(e Event) {
	that.Event = e
}

// SendTsxt 发送测试
func (that *Express) SendTest(to string) {
	that.Mail.From = that.Mailer.Username
	that.Mail.To = []string{to}
	that.Mail.Subject = "Test Subject"
	that.Mail.Text = []byte("Text Body is, of course, supported!")
	that.Mail.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	err := that.Send()
	logger.PanicError(err, "邮件发送", true)

	r, e := ioutil.ReadFile("./resources/views/test.html")
	logger.PanicError(e, "file", true)
	fmt.Println(string(r))

	// https://github.com/CloudyKit/jet/wiki/V1-Documentation
}
