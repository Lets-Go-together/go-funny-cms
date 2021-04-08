package mailer

import (
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"gocms/pkg/logger"
	"net/smtp"
	"os"
	"time"
)

type Express struct {
	// https://pkg.go.dev/github.com/jordan-wright/email
	Mailer  *Mailer
	Event   Event
	Options Options
}

// Options 配置选项
type Options struct {
	LoggerFile string
	NotifeType int
	Delay      time.Duration
	Event      Event
}

func NewExpress(mailer *Mailer) *Express {
	express := &Express{
		Mailer:  mailer,
		Options: Options{},
	}

	return express
}

func NewMailerExpress() *Express {
	mailer := NewMailer()
	express := &Express{
		Mailer:  mailer,
		Options: Options{},
	}

	return express
}

// SetOptions 设置基础配置
func (that *Express) SetOptions(options Options) error {
	if len(options.LoggerFile) > 0 {
		that.Options.LoggerFile = options.LoggerFile
	}

	if options.NotifeType > 0 {
		that.Options.NotifeType = options.NotifeType
	}

	if options.Event != nil {
		that.SetEvent(options.Event)
	}

	return nil
}

// Send 发送操作
func (that *Express) SendNow() error {
	addr := fmt.Sprintf("%s:%s", that.Mailer.Host, that.Mailer.Port)
	return that.Mailer.Mail.Send(addr, smtp.PlainAuth("", that.Mailer.Username, that.Mailer.Password, that.Mailer.Host))
}

// defaultLoggerFile 默认日志文件
func (that *Express) defaultLoggerFile(file string) error {
	that.Options.LoggerFile = "./storage/mail/system.log"
	if e := that.isFile(that.Options.LoggerFile); e != nil {
		return e
	}
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
		e = &DingTalk{}
	}

	that.SetEvent(e)
	return nil
}

// isFile 默认日志文件
func (that *Express) isFile(file string) error {
	if _, err := os.Open(file); err != nil {
		return err
	}
	return nil
}

// setEvent 支持重置Event
func (that *Express) SetEvent(e Event) {
	that.Event = e
}

// setEvent 支持重置Event
func (that *Express) SetLoggerFile(filename string) error {
	if e := that.isFile(that.Options.LoggerFile); e != nil {
		return e
	}
	return nil
}

// GetEvent 支持重置Event
func (that *Express) GetEvent() Event {
	if that.Event == nil {
		that.defaultEvent()
	}
	return that.Event
}

// GetLoggerFile 支持重置Event
func (that *Express) GetLoggerFile() string {
	return that.Options.LoggerFile
}

// sendTo 支持重置Event
func (that *Express) Send() error {
	err := that.SendNow()
	that.pipe(err, that.Mailer.Mail)

	return err
}

// pipe 处理各种问题
func (that *Express) pipe(err error, email *email.Email) error {
	event := that.GetEvent()
	mailJson, _ := json.Marshal(email)
	result := map[string]string{
		"to":     string(mailJson),
		"result": "",
	}

	if err == nil {
		event.Success("Success")
		return err
	}

	result["result"] = "发送失败, err: " + err.Error()
	return err
}

// SendTsxt 发送测试
func (that *Express) SendTest(to string) {
	that.Mailer.Mail.From = that.Mailer.Username
	that.Mailer.Mail.To = []string{to}
	that.Mailer.Mail.Subject = "Test Subject"
	that.Mailer.Mail.Text = []byte("Text Body is, of course, supported!")
	that.Mailer.Mail.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	err := that.Send()
	logger.PanicError(err, "邮件发送", true)

	// https://github.com/CloudyKit/jet/wiki/V1-Documentation
}
