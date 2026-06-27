package demo

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/zsmhub/ks-sdk/message"
)

func TestMessage(t *testing.T) {
	event, err := message.DecryptMessage([]byte("6HPxQ4wqbh4yf2bF+19ONg=="), MsgSecret)
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}
	plain, _ := json.Marshal(event)
	fmt.Println(string(plain))
}
