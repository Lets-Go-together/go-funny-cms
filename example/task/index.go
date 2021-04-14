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
)

var scheduler = schedule.New()

func Runing() {
	CToSchedule()

	scheduler.Launch()
	time.Sleep(time.Hour)
}

func CToSchedule() {
	scheduler.RegisterTask("dingding", func(context *schedule.Context) error {
		logger.Info("test", "执行成功 ~~~~")
		testDing()
		return nil
	})
	scheduler.RegisterTask("mail", func(context *schedule.Context) error {
		logger.Info("test", "执行成功 ~~~~")
		testMail()
		return nil
	})
	task := schedule.NewTask("dingding", "每个小时执行操作 ....", "0 * * * * ?")
	task = scheduler.AddTask(task)

	t := schedule.NewTask("mail", "每9点执行操作 ....", "0 9 * * * ?")
	t = scheduler.AddTask(t)

	// 初始化的时候立即执行一下
	//scheduler.StartTask(task.Id)
	//scheduler.StartTask(t.Id)
}

func testMail() {
	express := mail.NewMailerExpress()
	express.Mailer.Mail = &email.Email{
		To:      []string{"chenf@surest.cn"},
		From:    "Surest <2522257384@qq.com>",
		Subject: "深圳天气",
		HTML:    []byte(GetNews()),
		Headers: textproto.MIMEHeader{},
	}
	task := mail.NewTaskExpress()
	task.DispatchNow(express)
}

func testDing() {
	var dingToken = []string{"b960b8c2240b7d0f05b1ffbf26b4a7807efa2fb22603127dcbdc618ea48607ea"}
	cli := dingtalk.InitDingTalk(dingToken, "任务")
	cli.SendTextMessage("Test 任务 | TimeAt:" + time.Now().Format("2006-01-02 15:04:05"))
}

// 保持持续性 搞点有点用的通知吧
func GetNews() string {
	body := help.GetUrl("https://way.jd.com/jisuapi/weather?appkey=b9e5400562ec620c8b3224b9899a4b23&city=%E6%B7%B1%E5%9C%B3")
	if len(body) == 0 {
		return ""
	}
	result := gjson.Get(body, "result.result").Map()
	c := gjson.Get(body, "result.result.index").String()
	cresult := gjson.Parse(c).Array()

	title := fmt.Sprintf("<h2>城市: %s | 星期几: %s | 天气: %s</h2>", result["city"], result["week"], fmt.Sprintf("%s%s%s", result["temp"], "-", result["temphigh"]))
	content := "<ul>"
	for _, i := range cresult {
		content += "<li>" + i.Get("iname").String() + " | " + i.Get("ivalue").String() + " | " + i.Get("detail").String() + "</li>"
	}
	content += "</ul>"
	content = title + content

	return content
}
