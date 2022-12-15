package sms

type Strategy func([]Gateway) []Gateway

func Order(gties []Gateway) []Gateway {
	return gties
}

func Random([]Gateway) []Gateway {
	return nil
}
