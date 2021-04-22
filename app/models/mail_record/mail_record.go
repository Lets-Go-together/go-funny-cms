package mail_record

import "gocms/app/models/base"

type MailRecord struct {
	base.BaseModel
	SubmitterId int `json:"submitter_id"`
	EmailId     int `json:"email_id"`
}

func (MailRecord) TableName() string {
	return "mail_record"
}
