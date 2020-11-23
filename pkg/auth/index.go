package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"gocms/pkg/logger"
	"io"
	"time"
)

func Encrypto() string {
	key := []byte("0123456789ABCDEF")
	plaintext := []byte(time.Now().String())

	block, err := aes.NewCipher(key)
	logger.PanicError(err, "获取密钥实例", true)

	nonce := make([]byte, 12)
	_, err = io.ReadFull(rand.Reader, nonce)
	logger.PanicError(err, "ReadFull", true)

	aesgcm, err := cipher.NewGCM(block)
	logger.PanicError(err, "NewGCM", true)

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return string(ciphertext)
}

func decrypto() {
	key := []byte("0123456789ABCDEF")

	ciphertext, _ := hex.DecodeString("08f24c28f0fc9aef5812a35ce66235bc2488d6c29b") //加密生成的结果

	nonce, _ := hex.DecodeString("000000000000000000000000") //加密用的nonce

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(plaintext))
}
