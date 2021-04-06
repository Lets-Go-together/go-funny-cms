package mailer

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/jordan-wright/email"
	"gocms/pkg/logger"
	"net/smtp"
	"os"
	"path/filepath"
)

// 主要集中邮件发送类
// 供其他模块api调用
type Mailer struct {
	// https://pkg.go.dev/github.com/jordan-wright/email
	Mailer   *email.Email
	username string
	password string
	host     string
	port     string
}

// InitMailer 初始化邮件服务
// config 中加载项目排序支持如下
// username password host port
// 如未设置，则走默认参数配置
func (that *Mailer) InitMailer(config ...string) {
	that.Mailer = email.NewEmail()
	that.LoadDefaultConfig()
	that.LoadConfig(config...)
}

// Send 发送操作
func (that *Mailer) Send() error {
	addr := fmt.Sprintf("%s:%s", that.host, that.port)
	return that.Mailer.Send(addr, smtp.PlainAuth("", that.username, that.password, that.host))
}

// asyncSend 异步发送
func (that *Mailer) asyncSend() error {
	return errors.New("实现中....")
}

// SendTsxt 发送测试
func (that *Mailer) SendTest(to string) {
	//that.Mailer.From = that.username
	//that.Mailer.To = []string{to}
	//that.Mailer.Subject = "Test Subject"
	//that.Mailer.Text = []byte("Text Body is, of course, supported!")
	//that.Mailer.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	//err := that.Send()
	//logger.PanicError(err, "邮件发送", true)

	//r, e := ioutil.ReadFile("./resources/views/test.html")
	//logger.PanicError(e, "file", true)
	//fmt.Println(string(r))

	var root, _ = os.Getwd()
	var View = jet.NewHTMLSet(filepath.Join(root, "resources/views"))
	templateName := "auth/verify.jet"
	t, err := View.GetTemplate(templateName)
	if err != nil {
		logger.PanicError(err, "template", true)
	}

	var w bytes.Buffer
	vars := make(jet.VarMap)
	if err = t.Execute(&w, vars, nil); err != nil {
		// error when executing template
	}

	fmt.Println(w.String())
}
