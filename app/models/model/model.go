package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	gorm.Model
	ID uint64 `json:"id"`

	CreatedAt time.Time `gorm:"column:created_at;index"json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"json:"updated_at"`
}
