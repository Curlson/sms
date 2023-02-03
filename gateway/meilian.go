package gateway

import (
	"net/http"
	urlStd "net/url"
	"strings"

	"kd-saas/common/sms/utils"
)

type MeiLian struct {
	Username    string
	PasswordMd5 string
	Apikey      string
}

func NewMeiLian(username, password, apikey string) Gateway {
	return &MeiLian{
		Username:    username,
		PasswordMd5: password,
		Apikey:      apikey,
	}
}

func (ml *MeiLian) GetName() string {
	return "meilian"
}

/*
success:msgid  提交成功，发送状态请见 4.1
error:msgid  提交失败
error:Missing username  用户名为空
error:Missing password  密码为空
error:Missing apikey  APIKEY 为空
error:Missing recipient  收件人手机号码为空
error:Missing message content  短信内容为空或编码不正确
error:Account is blocked  帐号被禁用
error:Unrecognized encoding  编码未能识别
error:APIKEY or password error  APIKEY 或密码错误
error:Unauthorized IP address  未授权 IP 地址
error:Account balance is insufficient  余额不足
error:Throughput Rate Exceeded  发送频率受限
error:Invalid md5 password length  MD5 密码长度非 32 位
*/
func (ml *MeiLian) Send(mobile string, msg Msg) (*Response, error) {
	b, err := ml.send(mobile, msg.Content())
	if err != nil {
		return nil, err
	}
	ok, info, msgid := ml.parseResult(b)
	return NewReponse(ok, info, ml.GetName(), msgid), nil
}

func (ml *MeiLian) parseResult(res []byte) (ok bool, msg string, msgid string) {
	result := strings.SplitN(string(res), ":", 2)
	ok = result[0] == "success"

	if ok {
		msgid = result[1]
		msg = "提交成功"
	} else {
		msg = result[1] // msgid 也可能在这里
	}
	return
}

func (ml *MeiLian) header() (header http.Header) {
	header = make(http.Header)
	header.Set("Content-Type", "application/x-www-from-urlencoded")
	return header
}

func (ml *MeiLian) QueryValues(mobile, content string) urlStd.Values {
	v := urlStd.Values{}
	v.Set("username", ml.Username)
	v.Set("password_md5", ml.PasswordMd5)
	v.Set("apikey", ml.Apikey)
	v.Set("mobile", mobile)
	v.Set("content", urlStd.QueryEscape(content))
	v.Set("encode", "UTF-8")
	return v
}
func (ml *MeiLian) send(mobile, content string) ([]byte, error) {
	var (
		url, _      = urlStd.Parse("http://m.5c.com.cn/api/send/index.php")
		queryValues = ml.QueryValues(mobile, content)
		header      = ml.header()
	)
	return utils.Get(url, queryValues, header)
}
