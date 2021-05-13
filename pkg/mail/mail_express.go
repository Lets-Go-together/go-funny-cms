package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"
	"mime"
	"net/smtp"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Express struct {
	// https://pkg.go.dev/github.com/jordan-wright/email
	Mailer      *Mailer
	Title       string
	Event       Event
	Options     Options
	Attachments []interface{}
}

// Options 配置选项
type Options struct {
	LoggerFile string
	NotifeType int
	SendAt     time.Time
	Event      Event
}

// NewMailerExpress 初始化一个邮件发送通道
func NewExpress(mailer *Mailer) *Express {
	express := &Express{
		Mailer:  mailer,
		Options: Options{},
	}

	return express
}

// NewMailerExpress 初始化一个带配置的邮件发送通道
func NewMailerExpress() *Express {
	mailer := NewMailer()
	express := &Express{
		Mailer:  mailer,
		Options: Options{},
	}

	return express
}

// UpdateOptions 设置基础配置
func (that *Express) UpdateOptions(options Options) error {
	if len(options.LoggerFile) > 0 {
		that.Options.LoggerFile = options.LoggerFile
	}

	if options.NotifeType > 0 {
		that.Options.NotifeType = options.NotifeType
	}

	if options.Event != nil {
		that.UpdateEvent(options.Event)
	}

	return nil
}

// SendNow 立即发送
func (that *Express) SendNow() error {
	fmt.Println("send")
	addr := fmt.Sprintf("%s:%s", that.Mailer.Host, that.Mailer.Port)
	return that.Mailer.Mail.Send(addr, smtp.PlainAuth("", that.Mailer.Username, that.Mailer.Password, that.Mailer.Host))
}

// UpdateEvent 支持重置Event
func (that *Express) UpdateEvent(e Event) {
	that.Event = e
}

// GetEvent 重设Event
func (that *Express) GetEvent() Event {
	if that.Event == nil {
		that.defaultEvent()
	}
	return that.Event
}

// SetLoggerFile 设置日志文件
func (that *Express) SetLoggerFile(filename string) error {
	if e := that.isFile(that.Options.LoggerFile); e != nil {
		return e
	}
	return nil
}

// GetLoggerFile 支持重置Event
func (that *Express) GetLoggerFile() string {
	return that.Options.LoggerFile
}

func (that *Express) validateCondition() {
	// 检查是否配置邮件来源
	if len(that.Mailer.Mail.From) == 0 {
		name := config.GetString("MAIL_FORM_NAME", "System")
		that.Mailer.Mail.From = fmt.Sprintf("%s <%s>", name, that.Mailer.Username)
	}
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
	logger.Info("info", result)
	event.Failed("Error")
	return err
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

	that.UpdateEvent(e)
	return nil
}

// isFile 默认日志文件
func (that *Express) isFile(file string) error {
	if _, err := os.Open(file); err != nil {
		return err
	}
	return nil
}

func (that *Express) GetAttachments() []interface{} {
	attachments := that.Attachments
	return attachments
}

func (that *Express) HandleAttachments() {
	attachmentsWrapper := []*email.Attachment{}
	wg := sync.WaitGroup{}
	for _, attachment := range that.Attachments {
		wg.Add(1)
		val := attachment.(map[string]interface{})
		url := val["url"].(string)
		filename := val["filename"].(string)
		go func(url string, group *sync.WaitGroup) {
			mimeType := mime.TypeByExtension(filepath.Ext(filename))
			a, _ := that.Mailer.Mail.Attach(bytes.NewBufferString(help.GetUrl(url)), filename, fmt.Sprintf("%s; charset=utf-8", mimeType))
			attachmentsWrapper = append(attachmentsWrapper, a)
			wg.Done()
		}(url, &wg)
	}
	wg.Wait()

	that.Mailer.Mail.Attachments = attachmentsWrapper
}

func (that *Express) runing(id int) {
	that.ok(id, TASK_RUNING)
}

func (that *Express) complete(id int) {
	that.ok(id, TASK_COMPLETED)
}
func (that *Express) failed(id int, err error) {
	that.ok(id, TASK_FAILED)
	config.Db.Model(&MailerModel{}).Where("id = ?", id).Update(map[string]string{
		"remark": err.Error(),
	})
}

func (that *Express) ok(id int, status int) {
	config.Db.Model(&MailerModel{}).Where("id = ?", id).Update(map[string]int{
		"status": status,
	})
}

func (that *Express) Send(express *Express, id int) error {

	that.HandleAttachments()
	that.validateCondition()

	that.runing(id)
	err := that.SendNow()

	err = that.pipe(err, that.Mailer.Mail)

	if err == nil {
		that.complete(id)
	} else {
		that.failed(id, err)
	}

	return err
}
