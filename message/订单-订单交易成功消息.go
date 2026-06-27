package message

import "encoding/json"

// 文档：https://open.kwaixiaodian.com/zone/new/docs/msg?name=kwaishop_order_orderSuccess&version=1

type MsgOrderSuccess struct {
	Oid        int64 `json:"oid"`        // 订单ID
	SellerId   int64 `json:"sellerId"`   // 商家ID
	Status     int   `json:"status"`     // 订单状态
	UpdateTime int64 `json:"updateTime"` // 更新时间（毫秒）
}

var _ MessageData = MsgOrderSuccess{}

func (MsgOrderSuccess) Event() string {
	return "kwaishop_order_orderSuccess"
}

func (MsgOrderSuccess) DecodeData(info string) (MsgOrderSuccess, error) {
	var resp MsgOrderSuccess
	err := json.Unmarshal([]byte(info), &resp)
	return resp, err
}
