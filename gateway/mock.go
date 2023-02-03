package gateway

type Mock struct {
	AppId  string
	AppKey string
}

func NewMock(id, key string) Gateway {
	return &Mock{
		AppId:  "mock",
		AppKey: "mock",
	}
}
func (m *Mock) GetName() string {
	return "mock"
}

func (m *Mock) Send(mobile string, msg Msg) (*Response, error) {
	return NewReponse(true, "Mock 短信发送成功", m.GetName()), nil
}
