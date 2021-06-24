package logger

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	currentTime = time.Now().Format("2006-01-02")
	logFileName = flag.String("log", "./storage/log/"+currentTime+".log", "Log file name")
)

// 处理 Error 的情况
// isExit == true 时，将停止运行并输出
func PanicError(err error, source string, isExit bool) {
	if err != nil {
		dispatchNotice(err.Error(), source)
		if isExit == true {
			fmt.Println("错误消息:"+err.Error(), "\n"+"来源: "+source+"\n-----------")
			os.Exit(-1)
		}
	}
}

// 发送通知
func dispatchNotice(msg string, source string) {
	// 发送日志
	Error("error : "+source, msg)

	// 钉钉通知
}

func handle(level string, title string, content interface{}) {
	// TODO 修复单元测试无法获取 flag 问题
	//if true {
	//	return
	//}
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "Server Failed")
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("[%s] %s : %s \n", level, title, content)
}
