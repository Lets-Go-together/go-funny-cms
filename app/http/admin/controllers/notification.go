package controllers

import (
	"encoding/json"
	"github.com/spf13/cast"
	"gocms/app/http/admin/validates"
	"gocms/app/models/admin"
	"gocms/app/models/notification"
	"gocms/app/service"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/response"
	"gocms/wrap"
)

type NotificationController struct{}

func (n *NotificationController) List(c *wrap.ContextWrapper) {
	tag := c.DefaultQuery("tag", 0)
	page := c.DefaultQuery("page", 1)
	pageSize := c.DefaultQuery("pageSize", 10)

	model := &notification.Notification{}
	models := []notification.Notification{}
	query := config.Db.Model(&model)
	if len(tag) > 0 {
		query = query.Where("tag = ?", tag)
	}
	query = query.Limit(pageSize).Offset(page).Scan(&models)

	response.SuccessResponse(models).WriteTo(c)
	return
}

func (n *NotificationController) Store(c *wrap.ContextWrapper) {
	params := &validates.NotificationSaveValidate{}
	c.ShouldBind(&params)

	description, _ := json.Marshal(map[string]string{"content": params.Description})
	model := &notification.Notification{
		Title:        params.Title,
		Tag:          params.Tag,
		Submitter_id: cast.ToInt(admin.AuthUser.Id),
		Description:  string(description),
		Read_at:      nil,
	}

	follow_ids := []int{}
	if len(params.FollowIds) > 0 {
		follow_ids = params.FollowIds
	} else {
		list := new(service.AdminService).GetAllAdmins(map[string]interface{}{}, "id")
		for _, i := range list {
			follow_ids = append(follow_ids, cast.ToInt(i.ID))
		}
	}

	config.Db.Model(&notification.Notification{}).Create(&model)

	nUsers := []notification.NotificationUser{}
	for _, id := range follow_ids {
		nUsers = append(nUsers, notification.NotificationUser{
			NotificationId: cast.ToInt(model.ID),
			FollowId:       id,
		})
	}

	config.Db.Model(&notification.NotificationUser{}).Create(nUsers)

	response.SuccessResponse().WriteTo(c)
	return
}

func (n *NotificationController) Readed(c *wrap.ContextWrapper) {
	id := c.DefaultQuery("id", 0)

	nUser := &notification.NotificationUser{
		Read_at: help.GetCurrentTimestamp(),
	}
	config.Db.Model(&nUser).Where("follow_id = ? and notification_id = ?", admin.AuthUser.Id, id).Update(nUser)

	response.SuccessResponse().WriteTo(c)
	return
}
