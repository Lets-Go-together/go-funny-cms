package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// Logger channel
// --------------

func Info(title string, content interface{}) {
	c, _ := json.Marshal(content)
	handle("Info", title, string(c))
}

func Error(title string, content interface{}) {
	handle("Error", title, content)
}

func Debug(log string) {
	line, _ := callerInfo()
	log = strings.ReplaceAll(log, "\n", "\n\t")
	log = fmt.Sprintf("%s\n\t%s", line, log)
}

func callerInfo() (string, string) {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(3, rpc[:])
	if n < 1 {
		return "-", "-"
	}
	frame, _ := runtime.CallersFrames(rpc).Next()
	filePath := strings.ReplaceAll(frame.File, projectRootPath(), "")
	funcName := strings.Split(frame.Function, ".")[1]
	return fmt.Sprintf("%s:%d %s", filePath, frame.Line, funcName), funcName
}

func projectRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir+"/", "\\", "/", -1)
}
