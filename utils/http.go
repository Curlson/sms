package utils

import (
	"io"
	"net/http"
	urlStd "net/url"
	"time"
)

// 请求超时时间: 单位 s
var timeout time.Duration = 5 * time.Second

func Post(url string, body io.Reader, header ...http.Header) (data []byte, err error) {

	// 超时时间设置
	client := &http.Client{Timeout: timeout}

	// 初始化 request
	request, _err := http.NewRequest(
		http.MethodPost,
		url,
		body,
	)
	if _err != nil {
		return nil, _err
	}
	// 设置指定请求头
	if len(header) > 0 {
		request.Header = header[0]
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	// 关闭资源
	defer resp.Body.Close()
	data = make([]byte, 0)
	data, err = io.ReadAll(resp.Body)
	return
}

func Get(url *urlStd.URL, query urlStd.Values, header ...http.Header) (data []byte, err error) {

	// 超时设置
	client := &http.Client{Timeout: timeout}

	// query 参数设置
	url.RawQuery = query.Encode()
	// 配置请求
	request, _err := http.NewRequest(
		http.MethodGet,
		url.String(),
		nil,
	)
	if _err != nil {
		return nil, _err
	}
	// 设置指定请求头
	if len(header) > 0 {
		request.Header = header[0]
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	// 关闭资源
	defer resp.Body.Close()
	data = make([]byte, 0)
	data, err = io.ReadAll(resp.Body)
	return
}
