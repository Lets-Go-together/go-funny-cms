package validates

type MailListQuery struct {
	SubmitterId int    `json:"submitter_id"`
	Email       string `json:"email"`
	StartAt     string `json:"start_at"`
	EndAt       string `json:"end_at"`
	Status      int    `json:"status"`
}

type MailSendValidate struct {
	To          []string `json:"to" validate:"required"`
	Bcc         []string `json:"bcc"`
	Cc          []string `json:"cc"`
	Subject     string   `json:"subject" validate:"required"`
	HTML        string   `json:"html" validate:"required"`
	Attachments string   `json:"attachments" validate:"required"`
	ReplyTo     string   `json:"replyto" validate:"required"`
	SendAt      string   `json:"sendat" validate:"required"`
}
