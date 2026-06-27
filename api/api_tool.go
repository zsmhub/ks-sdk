package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"sort"
	"strings"
)

// MakeSign 按快手开放平台规则计算 HMAC_SHA256 签名。
//
// 规则：
//  1. 取参与签名的全部参数（access_token、appkey、method、param、signMethod、
//     timestamp、version），按参数名(key)字典序升序排列；
//  2. 以 key=value 的形式用 & 连接成字符串；
//  3. 在末尾追加 &signSecret=<signSecret>；
//  4. 以 signSecret 作为密钥对上述字符串做 HMAC-SHA256，结果做 Base64 编码即为签名。
//
// 注意：signSecret 既会被拼接进待签名字符串，又会作为 HMAC 的密钥，两处都不可省略。
func MakeSign(params map[string]string, signSecret string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	for i, k := range keys {
		if i > 0 {
			b.WriteByte('&')
		}
		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(params[k])
	}
	b.WriteString("&signSecret=")
	b.WriteString(signSecret)

	mac := hmac.New(sha256.New, []byte(signSecret))
	mac.Write([]byte(b.String()))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// MethodPath 将接口名（open.order.cursor.list）转换为请求路径（/open/order/cursor/list）
func MethodPath(method string) string {
	return "/" + strings.ReplaceAll(method, ".", "/")
}
