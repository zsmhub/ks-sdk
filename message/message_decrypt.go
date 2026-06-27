package message

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

// DecryptMessage 解密推送消息体。
// 返回解密后的 Event 结构体
func DecryptMessage(encryptedMsg []byte, msgSecret string) (*Event, error) {
	plain, err := DecryptRaw(encryptedMsg, msgSecret)
	if err != nil {
		return nil, err
	}
	var ev Event
	if err := json.Unmarshal(plain, &ev); err != nil {
		return nil, fmt.Errorf("消息解析失败: %v, 消息明文: %s", err, string(plain))
	}
	return &ev, nil
}

// DecryptRaw 解密推送消息体。
// 返回解密后的明文 JSON 字节。
func DecryptRaw(encryptedMsg []byte, msgSecret string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(msgSecret)
	if err != nil {
		return nil, errors.New("msgSecret base64 解码失败: " + err.Error())
	}
	if len(key) < aes.BlockSize {
		return nil, errors.New("msgSecret 长度不足，无法取出 IV")
	}

	ciphertext, err := base64.StdEncoding.DecodeString(string(bytes.TrimSpace(encryptedMsg)))
	if err != nil {
		return nil, errors.New("encryptedMsg base64 解码失败: " + err.Error())
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("kssdk: 创建 AES 解密器失败: %w", err)
	}
	if len(ciphertext) == 0 || len(ciphertext)%block.BlockSize() != 0 {
		return nil, errors.New("kssdk: 密文长度非法")
	}

	iv := make([]byte, block.BlockSize()) // 全 0 的 IV
	mode := cipher.NewCBCDecrypter(block, iv)
	plain := make([]byte, len(ciphertext))
	mode.CryptBlocks(plain, ciphertext)

	plain, err = pkcs7Unpad(plain, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return plain, nil
}

// pkcs7Unpad 去除 PKCS7 填充。
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("kssdk: 解密结果为空")
	}
	pad := int(data[length-1])
	if pad <= 0 || pad > blockSize || pad > length {
		return nil, errors.New("kssdk: PKCS7 填充非法")
	}
	for _, b := range data[length-pad:] {
		if int(b) != pad {
			return nil, errors.New("kssdk: PKCS7 填充内容非法")
		}
	}
	return data[:length-pad], nil
}
