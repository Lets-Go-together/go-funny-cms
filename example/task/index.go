package task

import (
	"fmt"
	"gocms/pkg/mail/mailer"
	"gocms/pkg/schedule"
)

func main() {

}

func SchedlueRun() {
	fmt.Println("启动成功 ! \n")
	var scheduler = schedule.New()
	scheduler.Launch()
}

func ExpressRun() {
	fmt.Println("启动成功 ! \n")
	mailer.ExpressRun()
}
