package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// ApiClient 业务接口调用客户端。
// AppKey / AppSecret / SignSecret 由调用方传入，AccessToken 通过 SetAccessToken 设置。
type ApiClient struct {
	AppKey      string
	AppSecret   string
	SignSecret  string
	AccessToken string
}

// SetAccessToken 设置访问令牌，返回自身以支持链式调用
func (c *ApiClient) SetAccessToken(token string) *ApiClient {
	c.AccessToken = token
	return c
}

// Execute 发起业务接口调用，自动组装公共参数并签名。
// req 为业务请求，resp 为业务响应（指针）。
func (c *ApiClient) Execute(req Request, resp Response) error {
	if c.AccessToken == "" {
		return errors.New("access_token不能为空")
	}

	// 业务参数序列化为 param
	paramBytes, err := json.Marshal(req.Param())
	if err != nil {
		return err
	}

	// 组装参与签名的公共参数
	params := map[string]string{
		"appkey":       c.AppKey,
		"access_token": c.AccessToken,
		"method":       req.Method(),
		"signMethod":   SignMethod,
		"timestamp":    strconv.FormatInt(time.Now().UnixMilli(), 10),
		"version":      Version,
		"param":        string(paramBytes),
	}
	params["sign"] = MakeSign(params, c.SignSecret)

	// 转为 form 表单
	form := make(url.Values, len(params))
	for k, v := range params {
		form.Set(k, v)
	}

	respBody, err := HttpForm(req.HttpMethod(), GatewayUrl+MethodPath(req.Method()), form)
	if err != nil {
		return err
	}

	// debug
	// if writeErr := os.WriteFile("./tmp_file.txt", respBody, 0644); writeErr != nil {
	// 	fmt.Printf("WriteFile error: %v\n", writeErr)
	// }

	if err = json.Unmarshal(respBody, resp); err != nil {
		return err
	}

	if resp.GetResult() == SuccessResult {
		return nil
	}

	return fmt.Errorf("ks: result=%d, error_msg=%s", resp.GetResult(), resp.GetErrorMsg())
}
