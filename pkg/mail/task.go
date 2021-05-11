package mail

import (
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"gocms/app/models/base"
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

type TaskWarp interface {
	UpdateState(id int, state int)
	QueryTaskByState(state int) []Express
	QueryTaskByTag(tag string) []Express
	QueryTaskByEmail(tag string) []Express
	AddTask(express Express) error
}

type TaskExpress struct{}

func ExpressRun() {
	express := &TaskExpress{}
	for true {
		express.load()
		time.Sleep(time.Second * 30)
	}
}

func NewTaskExpress() *TaskExpress {
	return &TaskExpress{}
}

func (that *TaskExpress) Dispatch(express *Express) error {

	model := that.parse(express)
	r := config.Db.Model(&model).Create(&model)

	if r.RowsAffected > 0 {
		return nil
	}

	return nil
}

func (that *TaskExpress) DispatchNow(express *Express) error {
	model := that.parse(express)
	_ = config.Db.Model(&model).Create(&model)
	err := express.Send(express, model.ID)
	return err
}

func (that *TaskExpress) parse(express *Express) *MailerModel {

	model := &MailerModel{
		Email:       express.Mailer.Mail.To[0],
		Subject:     express.Mailer.Mail.Subject,
		Content:     string(express.Mailer.Mail.HTML),
		Attachments: help.ToJson(express.GetAttachments()),
		Status:      TASK_WAIT,
		Mailer:      help.ToJson(express.Mailer),
		SendAt:      base.TimeAt(express.Options.SendAt),
	}

	return model
}

func (that *TaskExpress) load() {
	conditon := map[string]interface{}{
		"status": 1,
	}
	fmt.Println("runing ...")
	list := that.GetQueryTask(conditon)
	wg := sync.WaitGroup{}
	for _, m := range list {
		if that.allowSend(m) {
			express := that.forParse(&m)
			wg.Add(1)
			go func(express *Express, id int) {
				express.Send(express, id)
				wg.Done()
			}(express, m.ID)
		}
	}

	wg.Wait()
}

// 获取发送时间
func (that *TaskExpress) GetSendAt(duration time.Duration) string {
	return time.Now().Add(duration).Format("2006-01-02 15:04:05")
}

func (that *TaskExpress) GetQueryTask(condition interface{}) []MailerModel {

	list := []MailerModel{}
	query := config.Db.Model(&MailerModel{})

	query.Where(condition).Scan(&list)
	return list
}

func (that *TaskExpress) forParse(m *MailerModel) *Express {
	express := NewMailerExpress()
	mailer := &Mailer{}
	var attachments []interface{}
	json.Unmarshal([]byte(m.Mailer), &mailer)
	json.Unmarshal([]byte(m.Attachments), &attachments)
	express.Mailer.Mail = &email.Email{
		To:      []string{m.Email},
		From:    mailer.Username,
		Subject: m.Subject,
		HTML:    []byte(m.Content),
		Headers: textproto.MIMEHeader{},
	}
	express.Attachments = attachments

	return express
}

func (that *TaskExpress) allowSend(m MailerModel) bool {
	sendAt := help.ParseTime(time.Time(m.SendAt).Format(help.TimeLayut))
	now := help.Now()
	return sendAt.Before(now)
}
