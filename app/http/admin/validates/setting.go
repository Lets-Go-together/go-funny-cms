package validates

type EmailValidate struct {
	Email string `validate:"required,email" json:"email"`
}
