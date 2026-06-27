package message

import "encoding/json"

// 文档：https://open.kwaixiaodian.com/zone/new/docs/msg?name=kwaishop_order_addOrder&version=1

type MsgOrderAddOrder struct {
	Oid         int64  `json:"oid"`         // 订单id
	SellerID    int64  `json:"sellerId"`    // 商家id
	OpenID      string `json:"openId"`      // 买家openId
	CreateTime  int64  `json:"createTime"`  // 订单创建时间（毫秒）
	BuyerOpenID string `json:"buyerOpenId"` // 买家openId
}

var _ MessageData = MsgOrderAddOrder{}

func (MsgOrderAddOrder) Event() string {
	return "kwaishop_order_addOrder"
}

func (MsgOrderAddOrder) DecodeData(info string) (MsgOrderAddOrder, error) {
	var resp MsgOrderAddOrder
	err := json.Unmarshal([]byte(info), &resp)
	return resp, err
}
