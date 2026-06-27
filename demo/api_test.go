package demo

import "testing"

func TestGetOAuthUrl(t *testing.T) {
	GetOAuthUrl()
}

func TestGetAccessToken(t *testing.T) {
	code := "4a8df474a93aa4434335c4d3e9a1773dc2b319c0c65b965a720ff33c8b8569882b73c559"
	GetAccessToken(code)
}

func TestGetOrderList(t *testing.T) {
	GetOrderList()
}

func TestGetOrderDetail(t *testing.T) {
	GetOrderDetail()
}

func TestGetOrderRefundList(t *testing.T) {
	GetOrderRefundList()
}
