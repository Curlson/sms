package sms

import (
	"kd-saas/common/sms/gateway"
	"kd-saas/common/sms/utils"
)

// 消息体的结构

type Message struct {
	Msg  string
	Data utils.KV
	Code string
}

func NewMessage(content string, data utils.KV, code ...string) gateway.Msg {
	var tplCode string
	if len(code) > 0 {
		tplCode = code[0]
	}
	return &Message{
		Msg:  content,
		Data: data,
		Code: tplCode,
	}
}

// 所有参数
func (m *Message) Args() []any {
	values := []any{}
	for _, v := range m.Data {
		values = append(values, v)
	}
	return values
}

func (m *Message) ArgsMap() utils.KV {
	return m.Data
}

// 获取消息内容
func (m *Message) Content() string {
	return m.Msg
}

func (m *Message) TplCode() string {
	return m.Code
}
