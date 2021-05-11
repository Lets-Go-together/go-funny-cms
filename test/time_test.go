package test

import (
	"fmt"
	"gocms/pkg/help"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	sendAt := help.ParseTime("2021-05-10 18:32:57")
	now := time.Now()
	fmt.Println("历史时间", sendAt)
	fmt.Println("当前时间", now)
	fmt.Println("当前时间是否在历史时间之前", now.Before(sendAt))
}
