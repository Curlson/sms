package sms

import "time"

type Conf struct {
	// 超时时间
	Timeout time.Duration
	// 使用的网关
	DefGatewaies []string
	// 默认的负载均衡策略
	DefStrategy Strategy
	// 错误日志存放路径
	ErrorLog string
	// 所有网关配置
	Gateways map[string]Gateway
}

// apiid: cf_koodpower
// apikey: 869c0f09b3f6348b787aadef8f2a0538
func NewDefaultConf() *Conf {
	return &Conf{
		Timeout:      5,
		DefGatewaies: []string{"mock", "huyi", "meilian"},
		// 默认顺序发送
		DefStrategy: Order,
		ErrorLog:    "./tmp/log/sms.log",
		Gateways: map[string]Gateway{
			"mock":    NewMock("mock", "mock"),
			"huyi":    NewHuyi("cf_koodpower", "869c0f09b3f6348b787aadef8f2a0538"),
			"meilian": NewMeiLian("kdkjgs", "1ADBB3178591FD5BB0C248518F39BF6D", "ec03f995f3778a69ea861bfafdd5e0f6"),
			// "yunpian": {"api_key": "824f0ff2f71cab52936axxxxxxxxxx"},
			// "aliyun": {
			// 	"access_key_id":     "",
			// 	"access_key_secret": "",
			// 	"sign_name":         "",
			// },
		},
	}
}
