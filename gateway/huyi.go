package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"kd-saas/common/sms/utils"
)

type Huyi struct {
	ApiId  string
	ApiKey string
}

func NewHuyi(id, key string) Gateway {
	return &Huyi{
		ApiId:  id,
		ApiKey: key,
	}
}

type ResponseHuyi struct {
	Code  int    `json:"code"`
	Smsid string `json:"smsid"`
	Msg   string `json:"msg"`
}

func (r *ResponseHuyi) success() (success bool) {
	if r.Code == 2 {
		success = true
	}
	return
}

func (h *Huyi) GetName() string {
	return "huyi"
}

func (h *Huyi) Send(mobile string, msg Msg) (*Response, error) {
	b, err := h.send(mobile, msg.Content())
	// 请求初始化问题
	if err != nil {
		return nil, err
	}
	// 解析请求结果
	var response ResponseHuyi
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, fmt.Errorf("响应结果解析失败: " + string(b))
	}
	// 包装结果
	return NewReponse(response.success(), response.Msg, h.GetName(), response.Smsid), nil

}

func (h *Huyi) send(mobile string, content string) ([]byte, error) {
	var (
		url                = "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json"
		header http.Header = h.header()
		body   io.Reader   = h.body(mobile, content)
	)
	return utils.Post(url, body, header)
}

func (h *Huyi) header() (header http.Header) {
	header = make(http.Header)
	header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return header
}

func (h *Huyi) body(_mobile, _content string) (body *strings.Reader) {
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	v.Set("account", h.ApiId)
	v.Set("password", utils.GetMd5String([]byte(h.ApiId+h.ApiKey+_mobile+_content+_now)))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	body = strings.NewReader(v.Encode())
	return
}
