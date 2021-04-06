package mailer

import "gocms/pkg/config"

// 加载配置
func (that *Mailer) LoadConfig(mailer ...string) {
	for i, v := range mailer {
		switch i {
		case 0:
			that.ForUsername(v)
		case 1:
			that.ForPassword(v)
		case 2:
			that.ForHost(v)
		case 3:
			that.ForPort(v)
		}
	}
}

// 加载默认配置
func (that *Mailer) LoadDefaultConfig() {
	username := config.GetString("MAIL_USERNAME")
	password := config.GetString("MAIL_PASSWORD")
	host := config.GetString("MAIL_HOST")
	port := config.GetString("MAIL_PORT")
	that.LoadConfig(username, password, host, port)
}

// 设置用户名
func (that *Mailer) ForUsername(username string) {
	// 可以在这个地方自定义验证参数
	that.username = username
}

// 设置用户密码
func (that *Mailer) ForPassword(password string) {
	that.password = password
}

// 设置smtp域名
func (that *Mailer) ForHost(host string) {
	that.host = host
}

// 设置端口
func (that *Mailer) ForPort(port string) {
	that.port = port
}
