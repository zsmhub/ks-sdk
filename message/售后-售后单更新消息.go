package message

import "encoding/json"

// 文档：https://open.kwaixiaodian.com/zone/new/docs/msg?name=kwaishop_aftersales_updateRefund&version=1

type MsgAfterSalesUpdateRefund struct {
	OrderId           int64 `json:"orderId"`           // 订单ID
	RefundId          int64 `json:"refundId"`          // 售后单ID
	HandlingWay       int   `json:"handlingWay"`       // 退款方式，枚举： [1, "退货退款"] [10, "仅退款"] [3, "换货"][4, "补寄"][5, "维修"]
	SpecialRefundType int   `json:"specialRefundType"` // 特殊退款类型[0, "非特殊退款"] [1, "价保"]
	Status            int   `json:"status"`            // 订单退款状态
	EventType         int   `json:"eventType"`         // 售后事件类型
	UpdateTime        int64 `json:"updateTime"`        // 更新时间（毫秒）
}

var _ MessageData = MsgAfterSalesUpdateRefund{}

func (MsgAfterSalesUpdateRefund) Event() string {
	return "kwaishop_aftersales_updateRefund"
}

func (MsgAfterSalesUpdateRefund) DecodeData(info string) (MsgAfterSalesUpdateRefund, error) {
	var resp MsgAfterSalesUpdateRefund
	err := json.Unmarshal([]byte(info), &resp)
	return resp, err
}
