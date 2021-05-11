package validates

type EmailValidate struct {
	Id          int           `json:"id"`
	Subject     string        `json:"subject" validate:"required"`
	Emails      []string      `json:"emails" validate:"required"`
	Content     string        `json:"content" validate:"required"`
	Attachments []interface{} `json:"attachments"`
	SendAt      string        `json:"send_at"`
	Mailer      string        `json:"mailer"`
}

type EmailBindValidate struct {
	Email string `validate:"required,email" json:"email"`
}
