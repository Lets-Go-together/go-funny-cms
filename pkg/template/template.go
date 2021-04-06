package template

import (
	"bytes"
	"github.com/CloudyKit/jet"
)

type Template struct {
	W    *bytes.Buffer
	View *jet.Template
	Path string
}

// Init 参数初始化
func (t *Template) Init() {
	t.Path = ""
}

// Html 获取渲染结果
func (t *Template) Html() string {
	return t.W.String()
}
