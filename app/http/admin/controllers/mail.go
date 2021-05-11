package controllers

import (
	"github.com/jordan-wright/email"
	"github.com/spf13/cast"
	"gocms/app/http/admin/validates"
	"gocms/app/models/mail_record"
	"gocms/app/validates/validate"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/mail"
	"gocms/pkg/response"
	"gocms/wrap"
	"net/textproto"
)

type MailController struct{}

func (m *MailController) List(c *wrap.ContextWrapper) {
	page := c.DefaultQuery("page", 1)
	pageSize := c.DefaultQuery("pageSize", 10)

	model := &mail_record.MailRecord{}
	models := []mail_record.MailRecord{}
	query := config.Db.Model(&model)
	query = query.Limit(pageSize).Offset(page).Scan(&models)

	mailIds := []int{}
	for _, model := range models {
		mailIds = append(mailIds, cast.ToInt(model.ID))
	}

	condition := map[string]interface{}{
		"id": mailIds,
	}
	mailsModel := mail.NewTaskExpress().GetQueryTask(condition)
	response.SuccessResponse(mailsModel).WriteTo(c)
	return
}

// Store 处理邮件发送
func (m *MailController) Store(c *wrap.ContextWrapper) {
	var params validates.EmailValidate
	_ = c.ShouldBind(&params)

	if !validate.WithResponseMsg(params, c) {
		return
	}

	for _, to := range params.Emails {
		express := mail.NewMailerExpress()
		express.Mailer.Mail = &email.Email{
			To:      []string{to},
			Subject: params.Subject,
			HTML:    []byte(params.Content),
			Headers: textproto.MIMEHeader{},
		}
		express.Attachments = params.Attachments
		SendAt := help.ParseTime(params.SendAt)
		express.Options.SendAt = SendAt

		task := mail.NewTaskExpress()
		err := task.Dispatch(express)
		if err != nil {
			response.ErrorResponse(501, err.Error()).WriteTo(c)
			return
		}
	}

	response.SuccessResponse().WriteTo(c)
	return
}

func (m *MailController) Mailer(c *wrap.ContextWrapper) {
	mailers := config.GetMailerLabels()

	response.SuccessResponse(mailers).WriteTo(c)
	return
}
