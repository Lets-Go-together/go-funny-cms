package config

func GetAllMailers() map[string]map[string]interface{} {
	mailers := map[string]map[string]interface{}{
		"QQ邮箱配置1": {
			"label":    "2522257384@qq.com",
			"form":     "2522257384@qq.com",
			"host":     "smtp.qq.com",
			"port":     "25",
			"password": "clhbdmoztfujdich",
		},
	}

	return mailers
}

func GetMailerLabels() []string {
	mailers := GetAllMailers()
	labels := []string{}
	for label, _ := range mailers {
		labels = append(labels, label)
	}
	return labels
}
