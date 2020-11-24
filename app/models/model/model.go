package model

import (
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
