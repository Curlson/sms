package sms

type KV map[string]string

func (kv KV) Get(key string) string {
	return kv[key]
}

func (kv KV) Code() string {
	return kv.Get("code")
}

func (kv KV) Content() string {
	return kv.Get("content")
}
