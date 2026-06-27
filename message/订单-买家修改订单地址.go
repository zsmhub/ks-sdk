package message

import "encoding/json"

// 文档：https://open.kwaixiaodian.com/zone/new/docs/msg?name=kwaishop_order_addressUpdateAudit&version=1

type MsgOrderAddressUpdateAudit struct {
	Oid         int64  `json:"oid"`         // 订单ID
	OpenId      string `json:"openId"`      // 买家openId
	BuyerOpenId string `json:"buyerOpenId"` // 买家openId
	CreateTime  int64  `json:"createTime"`  // 创建时间（毫秒）
	AuditStatus int    `json:"auditStatus"` // 审核状态，10审批中、20成功、30拒绝
	AuditTime   int64  `json:"auditTime"`   // 审核时间（毫秒）
}

var _ MessageData = MsgOrderAddressUpdateAudit{}

func (MsgOrderAddressUpdateAudit) Event() string {
	return "kwaishop_order_addressUpdateAudit"
}

func (MsgOrderAddressUpdateAudit) DecodeData(info string) (MsgOrderAddressUpdateAudit, error) {
	var resp MsgOrderAddressUpdateAudit
	err := json.Unmarshal([]byte(info), &resp)
	return resp, err
}
