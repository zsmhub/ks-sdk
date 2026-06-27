package api

import (
	"net/http"
)

// 文档：https://open.kwaixiaodian.com/zone/new/docs/api?name=open.seller.order.refund.pcursor.list&version=1

type ReqOrderRefundList struct {
	// 订单生成的开始时间（单位毫秒），在当前时间的近90天内
	BeginTime int64 `json:"beginTime,omitempty"`
	// 订单生成的截止时间（单位毫秒）在当前时间的近90天内，且与开始时间的时间范围不大于1天
	EndTime int64 `json:"endTime,omitempty"`
	// 退款单请求类型，8 等待退款 9 全部退款订单
	Type int `json:"type,omitempty"`
	// 每次请求数量，最多一页100条
	PageSize int `json:"pageSize,omitempty"`
	// 当前页码，如 1
	CurrentPage int `json:"currentPage,omitempty"`
	// 订单id，可根据订单id查询关联的所有售后单列表
	OrderId int64 `json:"orderId,omitempty"`
	// 排序方式，1时间降序 2时间升序 ，默认降序
	Sort int `json:"sort,omitempty"`
	// 查找方式，1按创建时间查找 2按更新时间查找 ，默认创建时间
	QueryType int `json:"queryType,omitempty"`
	// 游标内容，第一次传空串，之后传上一次的pcursor返回值，若返回“nomore”则标识到底
	Pcursor string `json:"pcursor,omitempty"`
	// 退款状态，枚举：[10, "买家已经申请退款，等待卖家同意"] [12, "卖家已拒绝，等待买家处理"] [20, "协商纠纷，等待平台处理"] [30, "卖家已经同意退款，等待买家退货"] [40, "买家已经退货，等待卖家确认收货"] [45, "卖家已经发货，等待买家确认收货"] [50, "卖家已经同意退款，等待系统执行退款"] [60, "退款成功"] [70, "退款关闭"]
	Status int `json:"status,omitempty"`
}

var _ Request = ReqOrderRefundList{}

func (r ReqOrderRefundList) HttpMethod() string {
	return http.MethodPost
}

func (r ReqOrderRefundList) Method() string {
	return "open.seller.order.refund.pcursor.list"
}

func (r ReqOrderRefundList) Param() any {
	return r
}

type (
	RespOrderRefundList struct {
		BaseResponse
		Data struct {
			CurrentPage         int               `json:"currentPage"`
			PageSize            int               `json:"pageSize"`
			TotalPage           int               `json:"totalPage"`
			TotalSize           int               `json:"totalSize"`
			BeginTime           int64             `json:"beginTime"`
			EndTime             int64             `json:"endTime"`
			Pcursor             string            `json:"pcursor"`
			DataList            []RefundOrderInfo `json:"dataList"`
			RefundOrderInfoList []RefundOrderInfo `json:"refundOrderInfoList"`
		} `json:"data"`
	}
	RefundOrderInfo struct {
		RefundFee            int    `json:"refundFee"`
		LogisticsID          int    `json:"logisticsId"`
		RefundReasonDesc     string `json:"refundReasonDesc"`
		Oid                  int64  `json:"oid"`
		RefundType           int    `json:"refundType"`
		RefundDesc           string `json:"refundDesc"`
		RelItemID            int64  `json:"relItemId"`
		UrgeTimes            int    `json:"urgeTimes"`
		SellerID             int64  `json:"sellerId"`
		RelSkuID             int    `json:"relSkuId"`
		NegotiateStatus      int    `json:"negotiateStatus"`
		NegotiateUpdateTime  int64  `json:"negotiateUpdateTime"`
		RequireReturnFreight bool   `json:"requireReturnFreight"`
		SkuID                int64  `json:"skuId"`
		RefundReason         int    `json:"refundReason"`
		UpdateTime           int64  `json:"updateTime"`
		HandlingWay          int    `json:"handlingWay"`
		ItemID               int64  `json:"itemId"`
		ExpireTime           int    `json:"expireTime"`
		SubmitTime           int64  `json:"submitTime"`
		CreateTime           int64  `json:"createTime"`
		SkuNick              string `json:"skuNick"`
		EndTime              int64  `json:"endTime"`
		RefundID             int64  `json:"refundId"`
		ReceiptStatus        int    `json:"receiptStatus"`
		Status               int    `json:"status"`
	}
)

var _ Response = RespOrderRefundList{}
