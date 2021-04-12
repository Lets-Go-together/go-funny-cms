package mailer

import (
	"fmt"
	"github.com/jordan-wright/email"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"sync"
	"time"
)

const (
	TASK_WAIT      = 1
	TASK_RUNING    = 2
	TASK_COMPLETED = 3
	TASK_FAILED    = 4
)

type TaskExpress struct{}

func Run() {
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

	fmt.Println(r.Error.Error())
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
		Attachments: express.GetAttachments(),
		Status:      TASK_WAIT,
		Mailer:      help.ToJson(express.Mailer),
	}

	return model
}

func (that TaskExpress) load() {
	list := that.GetQueryTask()
	wg := sync.WaitGroup{}
	for _, m := range list {
		express := that.forParse(&m)
		wg.Add(1)
		go that.Send(express, m.ID, &wg)
	}

	wg.Wait()
}

func (that TaskExpress) GetQueryTask(status ...int) []MailerModel {
	statues := []int{}
	if len(status) > 0 {
		statues = status
	}

	list := []MailerModel{}
	config.Db.Model(&MailerModel{}).Where("status in (?)", statues).Scan(&list)

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
	express.Mailer.Mail = &email.Email{
		//To: m.Recipient.To,
		//From: m.Mailer.Username,
		//Subject: m.Subject,
		//HTML: []byte(m.Content),
		//Headers: textproto.MIMEHeader{},
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
