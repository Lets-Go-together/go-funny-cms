package service

import (
	"time"
)

type MailService struct{}

// GetHtmlForTemplate 获取使用模板后的html
func (m MailService) GetHtmlForTemplate(text string) []byte {
	return []byte(text)
}

// CalcuateDelay 计算延迟时间
func (m MailService) CalcuateDelayByNow(timeAt string) time.Duration {
	feature, _ := time.Parse("2006-01-02 15:04:05", timeAt)
	now := time.Now()
	if !now.Before(feature) {
		return 0
	}

	r := feature.Sub(now).Microseconds()
	return time.Duration(r)
}
