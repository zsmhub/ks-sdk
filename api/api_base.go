package api

const (
	// GatewayUrl 快手开放平台 API 网关地址
	GatewayUrl = "https://openapi.kwaixiaodian.com"
	// Version 接口版本号
	Version = "1"
	// SignMethod 签名算法
	SignMethod = "HMAC_SHA256"
)

// SuccessResult 快手业务返回成功时 result 字段的值（1 表示成功）
const SuccessResult int = 1

// Request 业务请求接口。
// Method 返回接口名（如 open.order.cursor.list），SDK 会据此推导请求路径；
// Param 返回业务参数对象，由 ApiClient 统一序列化为公共参数 param。
type Request interface {
	HttpMethod() string
	Method() string
	Param() any
}

// BaseResponse 快手开放平台公共响应字段
type BaseResponse struct {
	Result    int    `json:"result"`     // 业务返回码，1 表示成功
	Code      string `json:"code"`       // 主返回码
	Msg       string `json:"msg"`        // 主返回信息
	ErrorMsg  string `json:"error_msg"`  // 错误信息
	SubCode   string `json:"sub_code"`   // 子返回码
	SubMsg    string `json:"sub_msg"`    // 子返回信息
	RequestID string `json:"request_id"` // 请求ID
}

// Response 业务响应接口
type Response interface {
	GetResult() int
	GetCode() string
	GetMsg() string
	GetErrorMsg() string
}

func (r BaseResponse) GetResult() int {
	return r.Result
}

func (r BaseResponse) GetCode() string {
	return r.Code
}

func (r BaseResponse) GetMsg() string {
	return r.Msg
}

func (r BaseResponse) GetErrorMsg() string {
	return r.ErrorMsg
}
