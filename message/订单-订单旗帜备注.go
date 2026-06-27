package message

import "encoding/json"

// 文档：https://open.kwaixiaodian.com/zone/new/docs/msg?name=kwaishop_order_addNote&version=1

type MsgOrderAddNote struct {
	EventType    int    `json:"eventType"`    // 事件类型：1订单备注修改，2订单备注新增，3订单备注删除
	OperatorName string `json:"operatorName"` // 操作人昵称
	OperatorType int    `json:"operatorType"` // 操作人类型
	BizID        int64  `json:"bizId"`        // 订单id
	Note         string `json:"note"`         // 备注
	Type         int    `json:"type"`         // 1通用文本，2收货地址修改，3订单改价
	UpdateTime   int64  `json:"updateTime"`   // 变更时间
	Flag         string `json:"flag"`         // 订单旗帜类型
}

var _ MessageData = MsgOrderAddNote{}

func (MsgOrderAddNote) Event() string {
	return "kwaishop_order_addNote"
}

func (MsgOrderAddNote) DecodeData(info string) (MsgOrderAddNote, error) {
	var resp MsgOrderAddNote
	err := json.Unmarshal([]byte(info), &resp)
	return resp, err
}
