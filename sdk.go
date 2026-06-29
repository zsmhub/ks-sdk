package ks_sdk

import (
	"github.com/zsmhub/ks-sdk/api"
)

// 快手开放平台（快手小店）SDK 调用入口。

// NewKsClient 创建业务接口调用客户端，如获取订单列表、订单详情。
// 调用后需通过 SetAccessToken 设置访问令牌。
func NewKsClient(appKey, signSecret string) *api.ApiClient {
	return &api.ApiClient{
		AppKey:      appKey,
		SignSecret:  signSecret,
		AccessToken: "",
	}
}

// NewKsOAuthClient 创建授权接口调用客户端，如获取/刷新 AccessToken。
func NewKsOAuthClient(appKey, appSecret string) *api.OAuthClient {
	return &api.OAuthClient{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}
