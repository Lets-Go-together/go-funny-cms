package notification

import "gocms/app/models/base"

type NotificationUser struct {
	base.BaseModel
	NotificationId int    `json:"notification_id"`
	FollowId       int    `json:"follow_id"`
	Read_at        string `json:"read_at"`
}

func (NotificationUser) notification_user() string {
	return "notification"
}
