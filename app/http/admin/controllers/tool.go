package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"gocms/pkg/config"
	"gocms/pkg/enum"
	"gocms/pkg/logger"
	"gocms/pkg/response"
	"gocms/wrap"
	"os"
	"time"
)

type ToolController struct{}

func (*ToolController) Pwd(c *wrap.ContextWrapper) {
	pwd, _ := os.Getwd()

	response.SuccessResponse(map[string]string{
		"current": pwd,
	}).WriteTo(c)
}

// 获取七牛云上传密钥信息
func (*ToolController) Qiniu(c *wrap.ContextWrapper) {
	uploadInfo := make(map[string]string)
	//info, _ := config.Redis.Get(enum.CACHE_QINIU).Result()
	//if len(info) > 0 {
	//	r := gjson.Parse(info).Map()
	//	for i, v := range r {
	//		uploadInfo[i] = v.String()
	//	}
	//
	//	response.SuccessResponse(uploadInfo).WriteTo(c)
	//	return
	//}

	bucket := config.GetString("QINIU_BUKET_PATH")
	accessKey := config.GetString("QINIU_AK")
	secretKey := config.GetString("QINIU_SK")
	if len(bucket) == 0 || len(accessKey) == 0 || len(secretKey) == 0 {
		logger.PanicError(errors.New("缺失七牛配置参数"), "获取七牛参数", false)
		response.SuccessResponse(uploadInfo).WriteTo(c)
		return
	}

	putPolicy := storage.PutPolicy{
		Scope: bucket,
		//SaveKey: fmt.Sprintf("%s/%s$(ext)", uploadDir, help.Md5V(time.Now().String())),
	}
	mac := qbox.NewMac(accessKey, secretKey)
	token := putPolicy.UploadToken(mac)
	uploadInfo = map[string]string{
		"token":      token,
		"bucket":     bucket,
		"url":        config.GetString("QINIU_UPLOAD_URL"),
		"upload_dir": config.GetString("QINIU_UPLOAD_DIR"),
		"host":       config.GetString("QINIU_HOST"),
	}

	fmt.Println(uploadInfo)

	json, _ := json.Marshal(uploadInfo)
	config.Redis.Set(enum.CACHE_QINIU, string(json), time.Minute*50)

	response.SuccessResponse(uploadInfo).WriteTo(c)
}
