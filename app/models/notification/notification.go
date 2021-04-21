package notification

import "gocms/app/models/base"

type Notification struct {
	base.BaseModel
	Title        string      `json:"title"`
	Tag          int         `json:"tag"`
	Submitter_id int         `json:"submitter_id"`
	Description  string      `json:"description"`
	Read_at      base.TimeAt `json:"read_at"`
}

func (Notification) TableName() string {
	return "notification"
}
