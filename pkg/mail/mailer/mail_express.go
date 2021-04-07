package mailer

import (
	"fmt"
	"github.com/jordan-wright/email"
	"gocms/pkg/logger"
	"io/ioutil"
	"net/smtp"
)

type Express struct {
	// https://pkg.go.dev/github.com/jordan-wright/email
	Mail    *email.Email
	Mailer  *Mailer
	Event   Event
	Options Options
}

type Options struct {
	LoggerFile string
	NotifeType int
}

// Init 初始化express
func (that *Express) GetExpress(options Options) *Express {
	if len(options.LoggerFile) == 0 {
		that.Options.LoggerFile = that.defaultLoggerFile()
	}

	if options.NotifeType == 0 {

	}
}

// Send 发送操作
func (that *Express) Send() error {
	addr := fmt.Sprintf("%s:%s", that.Mailer.Host, that.Mailer.Port)
	return that.Mail.Send(addr, smtp.PlainAuth("", that.Mailer.Username, that.Mailer.Password, that.Mailer.Host))
}

// defaultLoggerFile 默认日志文件
func (that *Express) defaultLoggerFile() string {
	return "./storage/mail/system.log"
}

// checkIsFile 默认日志文件
func (that *Express) checkIsFile() error {
	return nil
}

// setEvent 默认日志文件
func (that *Express) setEvent() error {
	if that.Options.NotifeType == DINGTALK {
		e := DingTalk{}
		that.Event = e
	}
	switch expr {

	}
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
