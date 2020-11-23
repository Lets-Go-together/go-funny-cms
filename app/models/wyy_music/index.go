package wyy_music

import (
	"github.com/jinzhu/gorm"
	"time"
)

type WyyMusicUserModel struct {
	gorm.Model
	WyyId         uint   `gorm:"AUTO_INCREMENT;primary_key"`
	WyyAccount    string `gorm:"type:varchar(30);not null"`
	WyyPwd        string `gorm:"type:char(32);not null"`
	WyyCreateTime time.Time
	WyyStatus     int `gorm:"type:tinyint(32);not null;default:1"`
}
