package api

import (
	"net/http"
)

// 文档：https://open.kwaixiaodian.com/zone/new/docs/api?name=open.seller.order.goods.deliver&version=1

type ReqOrderGoodsDeliver struct {
	// 快递单号
	ExpressNo string `json:"expressNo"`
	// 快递公司code：https://open.kwaixiaodian.com/zone/new/docs/dev?pageSign=3d4aa40abaab83c4a5f5839f6848dcb71614264033704
	ExpressCode int `json:"expressCode"`
	// 订单编号
	OrderId int64 `json:"orderId"`
}

var _ Request = ReqOrderGoodsDeliver{}

func (r ReqOrderGoodsDeliver) HttpMethod() string {
	return http.MethodPost
}

func (r ReqOrderGoodsDeliver) Method() string {
	return "open.seller.order.goods.deliver"
}

func (r ReqOrderGoodsDeliver) Param() any {
	return r
}

type (
	RespOrderGoodsDeliver struct {
		BaseResponse
	}
)

var _ Response = RespOrderGoodsDeliver{}
