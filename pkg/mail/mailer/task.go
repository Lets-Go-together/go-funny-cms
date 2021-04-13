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

type TaskExpress struct{}

// 运行一个实例
func ExpressRun() {
	express := &TaskExpress{}
	for true {
		express.load()
		time.Sleep(time.Second * 10)
	}
}

// 初始化一个邮件实例
func NewTaskExpress() *TaskExpress {
	return &TaskExpress{}
}

// 异步执行发送
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

// 发送操作
func (that TaskExpress) Send(express *Express, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	that.runing(id)
	if err := express.Send(); err == nil {
		that.complete(id)
	} else {
		that.failed(id)
	}
}

// 同步执行发送
func (that TaskExpress) DispatchNow(express *Express) error {
	return express.Send()
}

// 解析入库
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

// 加载且读取
func (that TaskExpress) load() {
	statues := []int{1}
	condition := that.conditionTimeAt()
	list := that.GetQueryTask(statues, condition)
	wg := sync.WaitGroup{}
	for _, m := range list {
		express := that.forParse(&m)
		wg.Add(1)
		go that.Send(express, m.ID, &wg)
	}

	wg.Wait()
}

// 范围执行时间
func (that TaskExpress) conditionTimeAt() map[string]string {
	startAt := time.Now().Format("2006-01-02 15:04:05")
	endAt := time.Now().Add(time.Minute).Format("2006-01-02 15:04:05")

	return map[string]string{
		"start_at": startAt,
		"end_at":   endAt,
	}
}

// 获取发送时间
func (that TaskExpress) GetSendAt(duration time.Duration) string {
	return time.Now().Add(duration).Format("2006-01-02 15:04:05")
}

// 获取执行任务
func (that TaskExpress) GetQueryTask(status []int, condition map[string]string) []MailerModel {

	list := []MailerModel{}
	query := config.Db.Model(&MailerModel{})

	if len(status) > 0 {
		query.Where("status in (?)", status).Scan(&list)
	} else {
		query.Scan(&list)
	}

	return list
}

// 发送中
func (that TaskExpress) runing(id int) {
	that.ok(id, TASK_RUNING)
}

// 发送完成
func (that TaskExpress) complete(id int) {
	that.ok(id, TASK_COMPLETED)
}

// 发送失败
func (that TaskExpress) failed(id int) {
	that.ok(id, TASK_FAILED)
}

// 状态更新
func (that TaskExpress) ok(id int, status int) {
	config.Db.Model(&MailerModel{}).Where("id = ?", id).Update(map[string]int{
		"status": status,
	})
}

// 解析值实例中
func (that TaskExpress) forParse(m *MailerModel) *Express {
	express := NewMailerExpress()
	mailer := &Mailer{}
	recipient := &Recipient{}
	_ = json.Unmarshal([]byte(m.Mailer), &mailer)
	_ = json.Unmarshal([]byte(m.Recipient), &recipient)
	express.Mailer.Mail = &email.Email{
		To:      recipient.To,
		From:    mailer.Username,
		Subject: m.Subject,
		HTML:    []byte(m.Content),
		Headers: textproto.MIMEHeader{},
	}

	return express
}
