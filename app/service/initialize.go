package service

import (
	"fmt"
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

	for k, v := range where {
		db = db.Where(k+" = ?", v)
	}
	db.Count(&total)

	fmt.Println(total)
	return total == 0
}
