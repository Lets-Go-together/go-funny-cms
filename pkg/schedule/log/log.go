package log

import "fmt"

type Logger interface {
	D(typ int, log string)
	E(typ int, log string)
	Err(typ int, err error)
	I(typ int, log string)
}

func D(tag string, format string, p ...interface{}) {
	log("D", tag, format, p)
}
func E(tag string, format string, p ...interface{}) {
	log("E", tag, format, p)
}
func Err(tag string, err error) {
	log("E", tag, err.Error())
}
func I(tag string, format string, p ...interface{}) {
	log("I", tag, format, p)
}

func log(level string, tag string, log string, p ...interface{}) {
	lo := fmt.Sprintf(log, p)
	l := fmt.Sprintf("%s/%s: %s", level, tag, lo)
	fmt.Println(l)
}
