package log

import (
	"bytes"
	"fmt"
)

type Logger interface {
	D(typ int, log string)
	E(typ int, log string)
	Err(typ int, err error)
	I(typ int, log string)
}

func D(tag string, p ...interface{}) {
	log("D", tag, p)
}
func E(tag string, p ...interface{}) {
	log("E", tag, p)
}
func Err(tag string, err error) {
	log("E", tag, err.Error())
}
func I(tag string, p ...interface{}) {
	log("I", tag, p)
}

func log(level string, tag string, p ...interface{}) {
	b := bytes.Buffer{}
	for _, i := range p {
		b.WriteString(fmt.Sprintf("%v", i))
	}
	l := fmt.Sprintf("%s/%s: %s", level, tag, b.String())
	fmt.Println(l)
}
