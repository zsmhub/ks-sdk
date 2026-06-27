package api

import (
	"net/http"
)

// 文档：https://open.kwaixiaodian.com/zone/new/docs/api?name=open.seller.order.refund.detail&version=1

type ReqOrderRefundDetail struct {
	RefundID int64 `json:"refundId"`
}

var _ Request = ReqOrderRefundDetail{}

func (r ReqOrderRefundDetail) HttpMethod() string {
	return http.MethodPost
}

func (r ReqOrderRefundDetail) Method() string {
	return "open.seller.order.refund.detail"
}

func (r ReqOrderRefundDetail) Param() any {
	return r
}

type (
	RespOrderRefundDetail struct {
		BaseResponse
		Data RefundDetail `json:"data"`
	}
	RefundDetail struct {
		RefundID                           int64                `json:"refundId"`
		Oid                                int                  `json:"oid"`
		SkuID                              int                  `json:"skuId"`
		RefundFee                          int                  `json:"refundFee"`
		Status                             int                  `json:"status"`
		RefundReason                       int                  `json:"refundReason"`
		BuyerID                            int                  `json:"buyerId"`
		SellerID                           int                  `json:"sellerId"`
		RefundType                         int                  `json:"refundType"`
		HandlingWay                        int                  `json:"handlingWay"`
		RefundDesc                         string               `json:"refundDesc"`
		LogisticsID                        int                  `json:"logisticsId"`
		Pictures                           string               `json:"pictures"`
		NegotiateStatus                    int                  `json:"negotiateStatus"`
		ValidNegotiateBuyerModifyTimeStamp int                  `json:"validNegotiateBuyerModifyTimeStamp"`
		UpdateTime                         int64                `json:"updateTime"`
		RelSkuID                           int                  `json:"relSkuId"`
		RelItemID                          int                  `json:"relItemId"`
		SellerDisagreeImages               string               `json:"sellerDisagreeImages"`
		LogisticsInfo                      LogisticsInfo        `json:"logisticsInfo"`
		SellerDisagreeDesc                 string               `json:"sellerDisagreeDesc"`
		SellerDisagreeReason               int                  `json:"sellerDisagreeReason"`
		TimeLimitNegotiateChange           int                  `json:"timeLimitNegotiateChange"`
		NegotiateReason                    string               `json:"negotiateReason"`
		CreateTime                         int64                `json:"createTime"`
		NegotiateUpdateTime                int64                `json:"negotiateUpdateTime"`
		SubmitTime                         int64                `json:"submitTime"`
		ExpireTime                         int64                `json:"expireTime"`
		EndTime                            int64                `json:"endTime"`
		ItemID                             int                  `json:"itemId"`
		ReceiptStatus                      int                  `json:"receiptStatus"`
		ProductNum                         int                  `json:"productNum"`
		RefundReasonDesc                   string               `json:"refundReasonDesc"`
		BuyerOpenID                        string               `json:"buyerOpenId"`
		Address                            string               `json:"address"`
		BuyerExchangeAddress               BuyerExchangeAddress `json:"buyerExchangeAddress"`
		ExchangeInfo                       ExchangeInfo         `json:"exchangeInfo"`
		SpecialRefundType                  int                  `json:"specialRefundType"`
		LogisticsRiskInfos                 []any                `json:"logisticsRiskInfos"`
		ReturnFreightInfo                  ReturnFreightInfo    `json:"returnFreightInfo"`
	}
	LogisticsInfo struct {
		ExpressNo   string `json:"expressNo"`
		ExpressCode int    `json:"expressCode"`
	}
	BuyerExchangeAddress struct {
		EncryptedConsignee   string `json:"encryptedConsignee"`
		DesensitiseConsignee string `json:"desensitiseConsignee"`
		EncryptedMobile      string `json:"encryptedMobile"`
		DesensitiseMobile    string `json:"desensitiseMobile"`
		Province             string `json:"province"`
		ProvinceCode         int    `json:"provinceCode"`
		City                 string `json:"city"`
		CityCode             int    `json:"cityCode"`
		District             string `json:"district"`
		DistrictCode         int    `json:"districtCode"`
		EncryptedAddress     string `json:"encryptedAddress"`
		DesensitiseAddress   string `json:"desensitiseAddress"`
	}
	ExchangeInfo struct {
		SkuID    int    `json:"skuId"`
		SkuDesc  string `json:"skuDesc"`
		Quantity int    `json:"quantity"`
	}
	ReturnFreightInfo struct {
		ReturnFreightAmount      int      `json:"returnFreightAmount"`
		ReturnFreightStatus      int      `json:"returnFreightStatus"`
		ReturnFreightApplyDesc   string   `json:"returnFreightApplyDesc"`
		ReturnFreightApplyImages []string `json:"returnFreightApplyImages"`
	}
)

var _ Response = RespOrderRefundDetail{}
