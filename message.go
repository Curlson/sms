package sms

// 消息体的结构
type Msg interface {
	Content() string
	Args() []any
	ArgsMap() KV
	TplCode() string
}

type Message struct {
	Msg  string
	Data KV
	Code string
}

func NewMessage(content string, data KV, code ...string) Msg {
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

//
func (m *Message) ArgsMap() KV {
	return m.Data
}

// 获取消息内容
func (m *Message) Content() string {
	return m.Msg
}

func (m *Message) TplCode() string {
	return m.Code
}
