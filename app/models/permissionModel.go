package models

type Permission struct {
	BaseModel
	Name   string `json:"account"`
	Icon   string `json:"password,omitempty"`
	Url    string `json:"description"`
	Status int    `json:"email,omitempty"`
	Method string `json:"phone"`
	Pid    int    `json:"avatar,omitempty"`
}
