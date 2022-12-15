package test

import (
	"fmt"
	"testing"

	"github.com/Curlson/sms"
)

func TestSms(t *testing.T) {
	// 1. 配置
	// 初始化所有 通道（channel/短信服务商）的配置
	conf := sms.NewDefaultConf()
	sender := sms.NewSms(conf)
	// 2. 发送短信
	// 通过自定义模板编号发送信息
	// 自定义模板消息内包含各个渠道的消息信息
	// 根据具体的渠道获取对应的模板
	msg := sms.NewMessage("您的验证码是 3379", sms.KV{"code": "3379"})
	// 可以指定 gateway用以覆盖配置内的 gateway
	// var a = []string{"huyi", "alibaba"}
	//
	// var map[string][]string{
	// 	"huyi": []string{"1001", "验证码%s"},
	// }
	//
	res, _ := sender.Send("18556368687", msg, map[string]sms.KV{
		"mock": {"tpl": "1001"},
		// "huyi": {"content": "您的验证码是：6789。请不要把验证码泄露给其他人。"},
	})
	fmt.Println(res)

	// 1. 模板消息：{code:1, content：您的验证码是 %s，请在 5 分钟内操作完成！}
	// 2. 模板消息
	// 3. 通过配置 conf 获取 sender 对象
}

func TestSmsWithSpecifyGateWay(t *testing.T) {
	// 1. 配置
	// TODO 初始化默认配置和所有 client 配置信息
	conf := sms.NewDefaultConf()
	sender := sms.NewSms(conf)
	// 2. 发送短信
	// 通过自定义模板编号发送信息
	// 自定义模板消息内包含各个渠道的消息信息
	// 根据具体的渠道获取对应的模板
	msg := sms.NewMessage("您的验证码是 3379", sms.KV{"code": "3379"})
	// 可以指定 gateway用以覆盖配置内的 gateway
	// var a = []string{"huyi", "alibaba"}
	//
	sender.Send("18556368687", msg, nil)
}
