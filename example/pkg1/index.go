package pkg1

import (
	role2 "gocms/app/models/role"
	"gocms/app/service"
)

func init() {

}

func Echo() {
	roleService := new(service.RoleService)
	role := role2.RoleModel{
		Name:        "Surest",
		Description: "Surest",
	}
	roleService.UpdateOrCreateById(role)
}
