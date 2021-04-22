package controllers

import (
	"github.com/spf13/cast"
	"gocms/app/models/mail_record"
	"gocms/pkg/config"
	"gocms/pkg/mail"
	"gocms/pkg/response"
	"gocms/wrap"
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

func (m *MailController) Store(c *wrap.ContextWrapper) {

}
