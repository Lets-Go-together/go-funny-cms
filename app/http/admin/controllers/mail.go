package controllers

import (
	"github.com/jordan-wright/email"
	"gocms/app/http/admin/validates"
	"gocms/app/models/admin"
	"gocms/app/service"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/config"
	"gocms/pkg/enum"
	"gocms/pkg/mail"
	"gocms/pkg/response"
	"gocms/wrap"
)

type MailController struct{}

var mailService = new(service.MailService)

// List 查看发送邮件
// 根据权限配置可选，如果当前用户存在查看全部用户邮件的权限即可放开查看所有人的权限
func (m *MailController) List(c *wrap.ContextWrapper) {
	id := admin.AuthUser.Id
	model := &mail.MailerModel{}
	query := config.Db.Model(model)
	list := []mail.MailerModel{}

	params := &validates.MailListQuery{}
	c.ShouldBind(&params)

	if b := rabc.HasPermissionForUser(admin.AuthUser.Account, "showAllEmails", enum.PMETHODANY); b == false {
		query.Where(map[string]interface{}{"submitter_id": id})
	}

	if params.SubmitterId > 0 {
		query.Where(map[string]interface{}{"submitter_id": params.SubmitterId})
	}

	if len(params.Email) > 0 {
		query.Where(map[string]interface{}{"email": params.Email})
	}
	if params.Status > 0 {
		query.Where(map[string]interface{}{"status": params.Status})

	}
	if len(params.StartAt) > 0 {
		query.Where("send_at BETWEEN ? and ?", params.StartAt, params.EndAt)

	}

	query.Scan(&list)

	response.SuccessResponse(list).WriteTo(c)
	return
}

// Send 发送邮件
func (m *MailController) Send(c *wrap.ContextWrapper) {
	params := validates.MailSendValidate{}
	_ = c.ShouldBind(&params)

	express := mail.NewMailerExpress()
	express.Options.Delay = mailService.CalcuateDelayByNow(params.SendAt)
	express.Mailer.Mail = &email.Email{
		To:      params.To,
		Bcc:     params.Bcc,
		Cc:      params.Cc,
		From:    "Jordan Wright <2522257384@qq.com>", // 这个地方有异议 先写死
		Subject: params.Subject,
		HTML:    mailService.GetHtmlForTemplate(params.HTML),
	}
	task := mail.NewTaskExpress()
	if err := task.Dispatch(express); err != nil {
		response.ErrorResponse(500, "邮件通道异常").WriteTo(c)
		return
	}

	response.SuccessResponse().WriteTo(c)
	return
}
