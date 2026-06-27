package demo

// 示例凭证，来自快手开放平台为应用分配的证书信息。
// 实际接入时请替换为自己的凭证，并建议从配置/环境变量读取，不要硬编码到代码库。
const (
	AppKey     = "xxx" // 应用 appKey
	AppSecret  = "xxx" // 应用 appSecret
	SignSecret = "xxx" // 签名密钥 signSecret
	MsgSecret  = "xxx" // 消息加密解密 KEY，用于 AES 解密推送消息体

	// 以下为运行示例时需要替换的动态值
	AccessToken = "xxx"
	RedirectUri = "https://xxx.com/oa/ks/auth"
)
