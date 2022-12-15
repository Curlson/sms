package sms

import "fmt"

var (
	ErrInvalidMobile = fmt.Errorf("手机号码 格式错误")
	ErrMsgTemplate   = fmt.Errorf("短信模版 获取失败")
	ErrSendFailed    = fmt.Errorf("短信发送 失败")
)
