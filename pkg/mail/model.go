package mail

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type MailerModel struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Recipient   string    `json:"recipient"`
	Subject     string    `json:"subject"`
	Content     string    `json:"content"`
	Attachments string    `json:"attachments"`
	Status      int       `json:"status"`
	Mailer      string    `json:"mailer"`
	SendAt      string    `json:"send_at"`
	CreatedAt   time.Time `json:"created_at" gorm:"-,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"-,omitempty"`
}

func (MailerModel) TableName() string {
	return "email_tasks"
}

type Recipient struct {
	To []string `json:"to"`
}

func (that Recipient) Value() (driver.Value, error) {
	return "", nil
	//return json.Marshal(that)
}

func (v *Recipient) Scan(data interface{}) error {
	fmt.Println("Scan")
	bytes, _ := data.([]byte)
	r := Recipient{}
	err := json.Unmarshal(bytes, &r)
	*v = Recipient(r)

	return err
}
