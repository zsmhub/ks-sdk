package api

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HttpClient SDK 使用的 http 客户端，超时 1 分钟
var HttpClient = &http.Client{Timeout: time.Minute}

// HttpForm 以 application/x-www-form-urlencoded 方式发起 POST 请求
func HttpForm(httpMethod string, rawUrl string, form url.Values) ([]byte, error) {
	req, err := http.NewRequest(httpMethod, rawUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	return do(req)
}

func PostJson(rawUrl string, bodyJson string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, rawUrl, strings.NewReader(bodyJson))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return do(req)
}

// Get 发起 GET 请求，query 作为查询参数
func Get(rawUrl string, query url.Values) ([]byte, error) {
	if len(query) > 0 {
		if strings.Contains(rawUrl, "?") {
			rawUrl += "&" + query.Encode()
		} else {
			rawUrl += "?" + query.Encode()
		}
	}

	req, err := http.NewRequest(http.MethodGet, rawUrl, nil)
	if err != nil {
		return nil, err
	}

	return do(req)
}

func do(req *http.Request) ([]byte, error) {
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
