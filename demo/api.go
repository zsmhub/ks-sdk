package demo

import (
	"encoding/json"
	"fmt"
	"time"

	ks_sdk "github.com/zsmhub/ks-sdk"
	"github.com/zsmhub/ks-sdk/api"
)

// 调用 快手开放平台 API 示例

// GetOAuthUrl 生成商家授权页地址
func GetOAuthUrl() {
	authUrl := ks_sdk.NewKsOAuthClient(AppKey, AppSecret).
		GetOAuthUrl(RedirectUri, "merchant_shop,merchant_funds,merchant_item,merchant_order,merchant_promotion,merchant_user,merchant_refund,merchant_langbridge,user_info,merchant_video,merchant_material,merchant_comment,merchant_logistics", "state123")
	fmt.Printf("authUrl=%s\n", authUrl)
}

// GetAccessToken 通过授权码换取访问令牌
func GetAccessToken(code string) {
	token, err := ks_sdk.NewKsOAuthClient(AppKey, AppSecret).GetAccessToken(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", token)
}

// RefreshAccessToken 刷新访问令牌
func RefreshAccessToken(refreshToken string) {
	token, err := ks_sdk.NewKsOAuthClient(AppKey, AppSecret).RefreshAccessToken(refreshToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", token)
}

// GetOrderList 获取订单列表（游标方式）
func GetOrderList() {
	now := time.Now()
	req := api.ReqGetOrderList{
		OrderViewStatus: ks_sdk.OrderViewStatusAll,
		PageSize:        10,
		BeginTime:       now.AddDate(0, 0, -1).UnixMilli(),
		EndTime:         now.UnixMilli(),
		Cursor:          "",
		Sort:            ks_sdk.SortAsc,
		QueryType:       ks_sdk.OrderQueryTypeCreate,
	}
	var resp api.RespGetOrderList

	err := ks_sdk.NewKsClient(AppKey, AppSecret, SignSecret).
		SetAccessToken(AccessToken).
		Execute(req, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	ret, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(ret))

	// 翻页：将上一页返回的 cursor 传入下一次请求，直到 cursor 为空或 "nomore"
	fmt.Printf("nextCursor=%s\n", resp.Data.Cursor)
}

// GetOrderDetail 获取订单详情
func GetOrderDetail() {
	req := api.ReqGetOrderDetail{
		Oid: 2617801683312294,
	}
	var resp api.RespGetOrderDetail

	err := ks_sdk.NewKsClient(AppKey, AppSecret, SignSecret).
		SetAccessToken(AccessToken).
		Execute(req, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	ret, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(ret))
}

// GetOrderRefundList 获取订单售后列表（游标方式）
func GetOrderRefundList() {
	beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2026-06-09 00:00:00", time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2026-06-09 23:59:59", time.Local)
	req := api.ReqOrderRefundList{
		BeginTime:   beginTime.UnixMilli(),
		EndTime:     endTime.UnixMilli(),
		Type:        ks_sdk.RefundType9,
		PageSize:    2,
		CurrentPage: 1,
		OrderId:     2616002006698566,
		Sort:        ks_sdk.SortDesc,
		QueryType:   1,
	}
	var resp api.RespOrderRefundList

	err := ks_sdk.NewKsClient(AppKey, AppSecret, SignSecret).
		SetAccessToken(AccessToken).
		Execute(req, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	ret, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(ret))
}
