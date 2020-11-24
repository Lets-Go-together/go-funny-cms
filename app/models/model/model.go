package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	gorm.Model
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`

	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}
