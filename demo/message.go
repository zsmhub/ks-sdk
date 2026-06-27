package demo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	ks_sdk "github.com/zsmhub/ks-sdk"
	"github.com/zsmhub/ks-sdk/message"
)

// 接收快手开放平台消息推送示例（使用标准库 net/http）

// CallbackMain 启动一个简单的 HTTP 服务接收推送
func CallbackMain() {
	http.HandleFunc("/message/auth", handleAuth) // 获取商家授权回调
	http.HandleFunc("/message/push", handlePush) // 获取消息推送

	fmt.Println("listening on :1323 ...")
	_ = http.ListenAndServe(":1323", nil)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code == "" {
		http.Error(w, "code is empty", http.StatusBadRequest)
		return
	}
	token, err := ks_sdk.NewKsOAuthClient(AppKey, AppSecret).GetAccessToken(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("token=%+v\n", token)
	writeSuccess(w)
}

func handlePush(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "read body error", http.StatusBadRequest)
		return
	}

	// 用 msgSecret 解密得到明文JSON
	event, err := message.DecryptMessage(body, MsgSecret)
	if err != nil {
		fmt.Println("解密失败:", err)
		writeSuccess(w)
		return
	}

	fmt.Printf("event=%v\n", event)

	// 按消息标识分发处理
	switch event.Event {
	case message.MsgOrderPaySuccess{}.Event():
		fmt.Printf("订单已支付: %+v\n", event)
		pay, err := (message.MsgOrderPaySuccess{}).DecodeData(event.Info)
		if err != nil {
			fmt.Printf("订单已支付error: %+v\n", err)
		}
		fmt.Printf("订单已支付: %+v\n", pay)
	default:
		fmt.Println("未识别的消息类型:", event.Event)
	}

	// 返回成功，否则快手会重试
	writeSuccess(w)
}

func writeSuccess(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(message.SuccessResp())
}
