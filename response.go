package sms

type Response struct {
	Ok   bool   `json:"ok"`
	Msg  string `json:"msg"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewReponse(ok bool, msg, gateway string, id ...string) *Response {
	r := &Response{
		Ok:   ok,
		Msg:  msg,
		Name: gateway,
	}
	if len(id) > 0 {
		r.Id = id[0]
	}
	return r
}
