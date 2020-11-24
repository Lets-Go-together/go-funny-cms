package admin

import "gocms/app/models/model"

type Admin struct {
	model.BaseModel
	Account     string `json:"account"`
	Password    string `json:"password"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}
