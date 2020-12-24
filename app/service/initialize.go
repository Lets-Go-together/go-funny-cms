package service

import (
	"github.com/jinzhu/gorm"
)

// 批量验证是否可以创建
// == true 为是
func IsAllowOperationModel(where map[string]string, db *gorm.DB) bool {
	var total int
	if _, ok := where["id"]; ok == true {
		delete(where, "id")
		db = db.Where("id", "<>", where["id"])
	}
	db.Where(where).Count(&total)

	return total == 0
}
