package controllers

import (
	"github.com/spf13/cast"
	role "gocms/app/models/menu"
	"gocms/app/validates/validate"
	"gocms/pkg/config"
	"gocms/pkg/response"
	"gocms/wrap"
)

type MenuController struct{}
type MenuList struct {
	Id          uint64     `json:"id"`
	Name        string     `json:"name"`
	PId         int        `json:"p_id"`
	Weight      int        `json:"weight"`
	Status      int        `json:"status"`
	Icon        string     `json:"icon"`
	Component   string     `json:"component"`
	Description string     `json:"description"`
	CreatedAt   string     `json:"created_at"`
	Children    []MenuList `json:"children" gorm:"-"`
}

var MenuModel role.MenuModel

func (*MenuController) Index(c *wrap.ContextWrapper) {
	var Menus []MenuList
	keyword := c.Query("keyword")
	query := config.Db.Model(MenuModel)

	if len(keyword) > 0 {
		query = query.Where("name like ?", "%"+keyword+"%")
	}
	query = query.Select("id, name, status, icon, weight, p_id, component, created_at, description").Order("weight desc").Scan(&Menus)
	Menus = GetMenuTree(Menus, 1)

	response.SuccessResponse(Menus).WriteTo(c)
	return
}

func (*MenuController) Show(c *wrap.ContextWrapper) {
	var param IdParam
	c.ShouldBind(&param)

	var Menu role.MenuModel
	config.Db.Model(MenuModel).Where("id = ?", cast.ToString(param.Id)).Select("name, weight, status, p_id, icon, created_at, description").Scan(&Menu)

	response.SuccessResponse(Menu).WriteTo(c)
	return
}

func (*MenuController) Store(c *wrap.ContextWrapper) {
	var params role.MenuModel
	c.ShouldBind(&params)

	if !validate.WithResponseMsg(params, c) {
		return
	}
	config.Db.Model(MenuModel).Create(params)
	response.SuccessResponse().WriteTo(c)
	return
}

func (*MenuController) Save(c *wrap.ContextWrapper) {
	var params role.MenuModel
	c.ShouldBind(&params)

	if !validate.WithResponseMsg(params, c) {
		return
	}
	config.Db.Model(MenuModel).Where("id = ?", cast.ToString(params.ID)).Update(params)
	response.SuccessResponse().WriteTo(c)
	return
}

func (*MenuController) Destory(c *wrap.ContextWrapper) {
	var param IdParam
	c.ShouldBind(&param)

	config.Db.Delete(MenuModel, "id = "+cast.ToString(param.Id))
	response.SuccessResponse().WriteTo(c)
	return
}

// 获取权限节点树
func (that *MenuController) Tree(c *wrap.ContextWrapper) {
	var Menus []MenuList
	config.Db.Model(MenuModel).Select("id, name, status, icon, p_id, created_at, description").Order("weight desc").Scan(&Menus)
	tree := GetMenuTree(Menus, 1)

	response.SuccessResponse(tree).WriteTo(c)
	return
}

func GetMenuTree(menus []MenuList, pid int) []MenuList {
	var list []MenuList

	for _, v := range menus {
		if v.PId == pid {
			v.Children = GetMenuTree(menus, cast.ToInt(v.Id))
			list = append(list, v)
		}
	}

	return list
}
