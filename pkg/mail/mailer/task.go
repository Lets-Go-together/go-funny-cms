package mailer

import (
	"encoding/json"
	"github.com/jordan-wright/email"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"net/textproto"
	"sync"
	"time"
)

const (
	TASK_WAIT      = 1
	TASK_RUNING    = 2
	TASK_COMPLETED = 3
	TASK_FAILED    = 4
)

type TaskBroker interface {
	UpdateState(id int, state int)
	QueryTaskByState(state int) []Express
	AddTask(express Express) error
}

type TaskExpress struct{}

func ExpressRun() {
	express := &TaskExpress{}
	for true {
		express.load()
		time.Sleep(time.Second)
	}
}

func NewTaskExpress() *TaskExpress {
	return &TaskExpress{}
}

func (that TaskExpress) Dispatch(express *Express) error {
	if express.Options.Delay == 0 {
		return that.DispatchNow(express)
	}

	model := that.parse(express)
	r := config.Db.Model(&model).Create(&model)

	if r.RowsAffected > 0 {
		return nil
	}

	return nil
}

func (that TaskExpress) DispatchNow(express *Express) error {
	return express.Send()
}

func (that TaskExpress) parse(express *Express) *MailerModel {
	r := Recipient{To: express.Mailer.Mail.To}

	model := &MailerModel{
		Recipient:   help.ToJson(r),
		Subject:     express.Mailer.Mail.Subject,
		Content:     string(express.Mailer.Mail.HTML),
		Attachments: help.ToJson(express.GetAttachments()),
		Status:      TASK_WAIT,
		Mailer:      help.ToJson(express.Mailer),
		SendAt:      that.GetSendAt(express.Options.Delay),
	}

	return model
}

func (that TaskExpress) load() {
	list := that.GetQueryTask(1)
	wg := sync.WaitGroup{}
	for _, m := range list {
		express := that.forParse(&m)
		wg.Add(1)
		go that.Send(express, m.ID, &wg)
	}

	wg.Wait()
}

// 获取发送时间
func (that TaskExpress) GetSendAt(duration time.Duration) string {
	return time.Now().Add(duration).Format("2006-01-02 15:04:05")
}

func (that TaskExpress) GetQueryTask(status ...int) []MailerModel {

	list := []MailerModel{}
	query := config.Db.Model(&MailerModel{})

	if len(status) > 0 {
		query.Where("status in (?)", status).Scan(&list)
	} else {
		query.Scan(&list)
	}

	return list
}

func (that TaskExpress) runing(id int) {
	that.ok(id, TASK_RUNING)
}

func (that TaskExpress) complete(id int) {
	that.ok(id, TASK_COMPLETED)
}
func (that TaskExpress) failed(id int) {
	that.ok(id, TASK_FAILED)
}

func (that TaskExpress) ok(id int, status int) {
	config.Db.Model(&MailerModel{}).Where("id = ?", id).Update(map[string]int{
		"status": status,
	})
}

func (that TaskExpress) forParse(m *MailerModel) *Express {
	express := NewMailerExpress()
	mailer := &Mailer{}
	recipient := &Recipient{}
	json.Unmarshal([]byte(m.Mailer), &mailer)
	json.Unmarshal([]byte(m.Recipient), &recipient)
	express.Mailer.Mail = &email.Email{
		To:      recipient.To,
		From:    mailer.Username,
		Subject: m.Subject,
		HTML:    []byte(m.Content),
		Headers: textproto.MIMEHeader{},
	}

	return express
}

func (that TaskExpress) Send(express *Express, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	that.runing(id)
	if err := express.Send(); err == nil {
		that.complete(id)
	} else {
		that.failed(id)
	}
}
