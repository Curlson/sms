package test

import (
	"fmt"
	"testing"

	"github.com/Curlson/sms"
)

func TestDispatch(t *testing.T) {

}

func TestMeilian(t *testing.T) {
	// 1. 配置
	// 初始化所有 通道（channel/短信服务商）的配置
	conf := sms.NewDefaultConf()
	sender := sms.NewSms(conf)
	// 2. 发送短信
	// 通过自定义模板编号发送信息
	// 自定义模板消息内包含各个渠道的消息信息
	// 根据具体的渠道获取对应的模板
	msg := sms.NewMessage("您的验证码是 3379", sms.KV{"code": "3379"})
	res, _ := sender.Send("18556368687", msg, map[string]sms.KV{"meilian": {"content": "您好，您的验证码是：12345【美联】"}})
	fmt.Println(res)
}

func TestHuyi(t *testing.T) {
	// 1. 配置
	// 初始化所有 通道（channel/短信服务商）的配置
	conf := sms.NewDefaultConf()
	sender := sms.NewSms(conf)
	// 2. 发送短信
	// 通过自定义模板编号发送信息
	// 自定义模板消息内包含各个渠道的消息信息
	// 根据具体的渠道获取对应的模板
	msg := sms.NewMessage("您的验证码是 3379", sms.KV{"code": "3379"})
	res, _ := sender.Send("18556368687", msg, map[string]sms.KV{
		"huyi": {"content": "您的验证码是：6789。请不要把验证码泄露给其他人。"},
	})
	fmt.Println(res)
}

func TestMock(t *testing.T) {
	// 1. 配置
	// 初始化所有 通道（channel/短信服务商）的配置
	conf := sms.NewDefaultConf()
	sender := sms.NewSms(conf)
	// 2. 发送短信
	// 通过自定义模板编号发送信息
	// 自定义模板消息内包含各个渠道的消息信息
	// 根据具体的渠道获取对应的模板
	msg := sms.NewMessage("您的验证码是 3379", sms.KV{"code": "3379"})
	res, _ := sender.Send("18556368687", msg, map[string]sms.KV{"mock": nil})
	fmt.Println(res)
}
