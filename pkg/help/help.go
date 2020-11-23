package help

import (
	//"bytes"
	//"math/big"
	//"crypto/rand"
	"os"
)

// 获取env
func GetEnv(key string) string {
	return os.Getenv(key)
}

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
}
