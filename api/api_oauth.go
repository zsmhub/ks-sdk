package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// 文档：https://open.kwaixiaodian.com/zone/new/docs/dev?pageSign=e1d9e229332f4f233a04b44833a5dfe71614263940720

const (
	// OAuthAuthorizeUrl 商家授权页地址
	OAuthAuthorizeUrl = "https://open.kwaixiaodian.com/oauth/authorize"
	// OAuthAccessTokenUrl 获取访问令牌地址
	OAuthAccessTokenUrl = GatewayUrl + "/oauth2/access_token"
	// OAuthRefreshTokenUrl 刷新访问令牌地址
	OAuthRefreshTokenUrl = GatewayUrl + "/oauth2/refresh_token"
)

// OAuthClient 授权接口调用客户端
type OAuthClient struct {
	AppKey    string
	AppSecret string
}

// OAuthTokenResult 令牌结果，获取与刷新共用
type OAuthTokenResult struct {
	// Result 为结果标识，1 表示成功。
	Result int `json:"result"`
	// Error 错误标识，成功时为空。
	Error string `json:"error"`
	// ErrorMsg 错误描述。
	ErrorMsg string `json:"error_msg"`
	// AccessToken 访问令牌，用于调用业务接口。
	AccessToken string `json:"access_token"`
	// RefreshToken 刷新令牌，用于在访问令牌过期前刷新。
	RefreshToken string `json:"refresh_token"`
	// ExpiresIn 访问令牌有效期，单位秒，默认为172800，即48小时。
	ExpiresIn int64 `json:"expires_in"`
	// RefreshTokenExpiresIn 刷新令牌有效期（秒）。默认为180天。
	RefreshTokenExpiresIn int64 `json:"refresh_token_expires_in"`
	// OpenID 授权商家的 openId。
	OpenID string `json:"open_id"`
	// Scopes 实际授予的权限列表。
	Scopes []string `json:"scopes"`
}

// GetOAuthUrl 生成商家授权页地址。
// redirectUri 为授权回调地址，scope 为申请的权限范围（多个用逗号分隔，可空），state 为透传参数。
func (c *OAuthClient) GetOAuthUrl(redirectUri, scope, state string) string {
	q := url.Values{}
	q.Set("app_id", c.AppKey)
	q.Set("redirect_uri", redirectUri)
	q.Set("response_type", "code")
	if scope != "" {
		q.Set("scope", scope)
	}
	if state != "" {
		q.Set("state", state)
	}
	return OAuthAuthorizeUrl + "?" + q.Encode()
}

// GetAccessToken 通过授权码 code 换取访问令牌
func (c *OAuthClient) GetAccessToken(code string) (OAuthTokenResult, error) {
	q := url.Values{}
	q.Set("app_id", c.AppKey)
	q.Set("app_secret", c.AppSecret)
	q.Set("grant_type", "code")
	q.Set("code", code)

	return c.requestToken(OAuthAccessTokenUrl, q)
}

// RefreshAccessToken 通过 refresh_token 刷新访问令牌
func (c *OAuthClient) RefreshAccessToken(refreshToken string) (OAuthTokenResult, error) {
	q := url.Values{}
	q.Set("app_id", c.AppKey)
	q.Set("app_secret", c.AppSecret)
	q.Set("grant_type", "refresh_token")
	q.Set("refresh_token", refreshToken)

	return c.requestToken(OAuthRefreshTokenUrl, q)
}

func (c *OAuthClient) requestToken(rawUrl string, q url.Values) (OAuthTokenResult, error) {
	var resp OAuthTokenResult

	body, err := Get(rawUrl, q)
	if err != nil {
		return resp, err
	}

	// 避免 json.Unmarshal 出现 unexpected end of JSON input 错误
	if len(body) == 0 {
		return resp, errors.New("http resp body is nil")
	}

	if err = json.Unmarshal(body, &resp); err != nil {
		return resp, err
	}

	if resp.Result == SuccessResult {
		return resp, nil
	}

	return resp, fmt.Errorf("ks: result=%d, error=%s, error_msg=%s", resp.Result, resp.Error, resp.ErrorMsg)
}
