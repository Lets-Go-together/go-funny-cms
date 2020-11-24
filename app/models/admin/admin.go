package admin

import "gocms/app/models/model"

type Admin struct {
	model.BaseModel
	Account     string
	Password    string
	Description string
	Email       string
	Phone       string
}
