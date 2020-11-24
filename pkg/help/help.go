package help

import (
	//"bytes"
	//"math/big"
	//"crypto/rand"
	"crypto/md5"
	"encoding/hex"
	"os"
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
