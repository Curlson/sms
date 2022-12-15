package sms

type Gateway interface {
	Send(mobile string, msg Msg) (*Response, error)
	GetName() string
}
