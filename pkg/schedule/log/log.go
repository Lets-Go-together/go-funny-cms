package log

import "fmt"

type Logger interface {
	D(typ int, log string)
	E(typ int, log string)
	Err(typ int, err error)
	I(typ int, log string)
}

func D(tag string, l string) {
	log("D", tag, l)
}
func E(tag string, l string) {
	log("E", tag, l)
}
func Err(tag string, err error) {
	log("E", tag, err.Error())
}
func I(tag string, l string) {
	log("I", tag, l)
}

func log(level string, tag string, log string) {
	l := fmt.Sprintf("%s/%s: %s", level, tag, log)
	fmt.Println(l)
}
