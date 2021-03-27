package help

import (
	//"bytes"
	//"math/big"
	//"crypto/rand"
	"crypto/md5"
	"encoding/hex"
	"os"
	"reflect"
	"time"
)

// 获取env
func GetEnv(key string) string {
	return os.Getenv(key)
}

// 获取随机字符串
func createRandomString(len int) string {
	//var container string
	//var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	//b := bytes.NewBufferString(str)
	//
	//length := b.Len()
	//bigInt := big.NewInt(int64(length))
	//
	//for i := 0; i < len; i++ {
	//	randomInt, _ := rand.Int(rand.Reader, bigInt)
	//}
	return ""
}

// 生成一个md5
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

// 获取偏移量
func GetOffset(page int, pageSize int) int {
	return (page - 1) * pageSize
}

// 获取当前时间
func getCurrentTimestamp() string {
	return time.Now().Format("2006-01-02")
}

func StructMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StructsMap(obj interface{}) []map[string]interface{} {
	var data []map[string]interface{}
	n := obj.([]interface{})

	for _, item := range n {
		data = append(data, StructMap(item))
	}
	return data
}

// 获取默认参数值
func GetDefaultParam(params ...interface{}) interface{} {
	if len(params) > 0 {
		return params[0]
	}
	return nil
}
