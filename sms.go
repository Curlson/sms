package sms

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

type Sms struct {
	*Conf
}

func NewSms(conf *Conf) *Sms {
	return &Sms{
		Conf: conf,
	}
}

func (s *Sms) Send(mobile string, msg Msg, gtys map[string]KV) ([]Response, error) {
	var (
		avGties  []string  // 可用的网关: 还未进行 dispatch
		disGties []Gateway // 调度后的策略
		err      error
	)
	// 1. 如果自定义了 gtys 则使用指定的 gtys 内的策略
	//    如果没有指定 gtys 则使用 config 内配置的默认网关内选择
	if avGties, err = s.gateway(gtys); err == nil {
		// 实例化对应的网关
		var senders []Gateway
		for _, v := range avGties {
			if s.Conf.Gateways[v] != nil {
				senders = append(senders, s.Conf.Gateways[v])
			}
		}
		// 获取策略，发送短信
		disGties = s.dispatch(senders...)
		return send(disGties, msg, mobile, gtys)
	}
	return nil, err
}

func send(senders []Gateway, msg Msg, mobile string, gtys map[string]KV) (result []Response, err error) {
	var success bool
	result = make([]Response, 0)
	for _, sender := range senders {
		if _msgMap, ex := gtys[sender.GetName()]; ex && _msgMap != nil {
			// 初始化
			msg = NewMessage(_msgMap.Content(), _msgMap, _msgMap.Code())
		}
		_res, err := sender.Send(mobile, msg)
		res := *_res
		if err == nil && res.Ok {
			success = true
			result = append(result, res)
			break
		} else {
			result = append(result, res)
		}
	}
	if !success {
		err = fmt.Errorf("所有网关发送短信失败")
	}
	return
}

func Result(gty string, res string, status string) map[string]any {
	return map[string]any{
		"status":  status,
		"gateway": gty,
		"result":  res,
	}
}

// 获取所有可用的 gatewaies
func (s *Sms) gateway(specificGty map[string]KV) ([]string, error) {
	// 所有的网关
	available := s.Conf.Gateways
	if len(available) == 0 {
		return nil, fmt.Errorf("没有可用的网关")
	}

	// 0. 获取所有可以使用的 gatewaies
	availableKeys := []string{}
	for k := range available {
		availableKeys = append(availableKeys, k)
	}
	aset := mapset.NewSet(availableKeys...)
	// 没有限定具体网关和配置
	if len(specificGty) == 0 {
		wset := mapset.NewSet(s.Conf.DefGatewaies...)
		// 获取可用的 网关
		result := aset.Intersect(wset)
		// 需要返回有序的 gateway
		return result.ToSlice(), nil
	} else {
		keys := []string{}
		for k := range specificGty {
			keys = append(keys, k)
		}
		sset := mapset.NewSet(keys...)
		return aset.Intersect(sset).ToSlice(), nil
	}
}

func (s *Sms) dispatch(gatewaies ...Gateway) []Gateway {
	fn := s.Conf.DefStrategy
	return fn(gatewaies)
}
