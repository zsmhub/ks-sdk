package message

import "encoding/json"

// 文档：https://open.kwaixiaodian.com/zone/new/docs/msg?name=kwaishop_order_addressChange&version=1

type MsgOrderAddressChange struct {
	Oid         int64  `json:"oid"`         // 订单ID
	SellerId    int64  `json:"seller_id"`   // 商家ID
	OpenId      string `json:"openId"`      // 买家openId
	BuyerOpenId string `json:"buyerOpenId"` // 买家openId
	UpdateTime  int64  `json:"updateTime"`  // 更新时间（毫秒）
}

var _ MessageData = MsgOrderAddressChange{}

func (MsgOrderAddressChange) Event() string {
	return "kwaishop_order_addressChange"
}

func (MsgOrderAddressChange) DecodeData(info string) (MsgOrderAddressChange, error) {
	var resp MsgOrderAddressChange
	err := json.Unmarshal([]byte(info), &resp)
	return resp, err
}
