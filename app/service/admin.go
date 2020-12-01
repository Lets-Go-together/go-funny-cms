package service

import (
	"fmt"
	"gocms/app/models/model"
	"gocms/pkg/config"
)

type AdminService struct{}
type listStruct struct {
	ID          uint64       `json:"id"`
	Account     string       `json:"account"`
	Description string       `json:"description"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
	CreatedAt   model.TimeAt `json:"created_at"`
}

func (*AdminService) GetList() *[]listStruct {
	admins := []listStruct{}
	config.Db.Table("admins").Select("id, account, description, email, phone, created_at").Scan(&admins)

	fmt.Println(admins)

	return &admins
}
