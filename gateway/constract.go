package gateway

import "kd-saas/common/sms/utils"

type Gateway interface {
	Send(mobile string, msg Msg) (*Response, error)
	GetName() string
}

type Msg interface {
	Content() string
	Args() []any
	ArgsMap() utils.KV
	TplCode() string
}
