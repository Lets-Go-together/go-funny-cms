package wyyMusic

import (
	"gocms/app/models/wyy_music"
	"gocms/pkg/config"
	"log"
)

func init() {

}

func Run() {

	WyyMusicUserModel := wyy_music.WyyMusicUser{}

	log.Fatal(config.Db.Model(&WyyMusicUserModel).First(&WyyMusicUserModel))

}
