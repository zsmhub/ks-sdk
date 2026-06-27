package message

// 文档：https://open.kwaixiaodian.com/zone/new/docs/dev?pageSign=120a69028f0e9ea69145644317ae8bcd1614263915925

// PushMessage 快手开放平台消息推送的外层信封。
// 真正的业务数据在 EncryptedMsg 中，需用 msgSecret 解密后得到明文 JSON。
type Event struct {
	// EventID 消息唯一ID，可用于幂等去重
	EventID string `json:"eventId"`
	// MsgID 业务消息内容唯一id
	MsgID string `json:"msgId"`
	// BizID 业务id如订单id、退款单id、商品id
	BizID int64 `json:"bizId"`
	// UserID 商家ID
	UserID int64 `json:"userId"`
	// OpenID 商家openId
	OpenID string `json:"openId"`
	// AppKey 应用Id
	AppKey string `json:"appKey"`
	// Event 消息标识
	Event string `json:"event"`
	// Info 为各消息内容，业务内容Json串
	Info string `json:"info"`
	// CreateTime 消息创建时间(毫秒时间戳)
	CreateTime int64 `json:"createTime"`
	// UpdateTime 更新时间(毫秒时间戳)
	UpdateTime int64 `json:"updateTime"`
	// Test 是否为心跳测试事件。为 true 时通常无需处理业务
	Test bool `json:"test"`
}

// MessageData 业务消息数据接口
type MessageData interface {
	Event() string // 返回消息标识
}

// MessageResp 收到消息后需要返回给快手的应答体。
// 返回成功内容，否则快手会认为推送失败并重试（最多 3 次）。
type MessageResp struct {
	Result int64 `json:"result"` // 1 表示成功
}

// SuccessResp 标准成功应答
func SuccessResp() MessageResp {
	return MessageResp{Result: 1}
}
