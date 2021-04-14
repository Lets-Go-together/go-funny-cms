package mail

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"gocms/pkg/config"
)

// 主要集中邮件发送类
// 供其他模块api调用
type Mailer struct {
	Mail     *email.Email
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

// 加载配置
func (that *Mailer) LoadConfig(mailer ...string) *Mailer {
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

	return that
}

// 加载默认配置
func NewMailer() *Mailer {
	username := config.GetString("MAIL_USERNAME")
	password := config.GetString("MAIL_PASSWORD")
	host := config.GetString("MAIL_HOST")
	port := config.GetString("MAIL_PORT")
	Mail := email.NewEmail()

	return &Mailer{
		Mail:     Mail,
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

// 设置用户名
func (that *Mailer) ForUsername(username string) {
	// 可以在这个地方自定义验证参数
	that.Username = username
}

// 设置用户密码
func (that *Mailer) ForPassword(password string) {
	that.Password = password
}

// 设置smtp域名
func (that *Mailer) ForHost(host string) {
	that.Host = host
}

// 设置端口
func (that *Mailer) ForPort(port string) {
	that.Port = port
}

func (that Mailer) Value() (driver.Value, error) {
	return json.Marshal(that)
}

func (v *Mailer) Scan(data interface{}) error {
	fmt.Println("Scan")
	bytes, _ := data.([]byte)
	r := Mailer{}
	err := json.Unmarshal(bytes, &r)
	*v = Mailer(r)

	return err
}
