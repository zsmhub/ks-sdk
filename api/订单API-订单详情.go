package api

import (
	"net/http"
)

// 文档：https://open.kwaixiaodian.com/zone/new/docs/api?name=open.order.detail&version=1

type ReqGetOrderDetail struct {
	// Oid 订单号
	Oid int64 `json:"oid"`
}

var _ Request = ReqGetOrderDetail{}

func (r ReqGetOrderDetail) HttpMethod() string {
	return http.MethodGet
}

func (r ReqGetOrderDetail) Method() string {
	return "open.order.detail"
}

func (r ReqGetOrderDetail) Param() any {
	return r
}

type (
	RespGetOrderDetail struct {
		BaseResponse
		Data struct {
			// 订单基础信息
			OrderBaseInfo OrderDetailBaseInfo `json:"orderBaseInfo"`
			// 订单商品信息
			OrderItemInfo OrderItemInfo `json:"orderItemInfo"`
			// 订单退款信息，返回最新的一个，若要获取当前订单关联所有退款单，可通过退款单列表API获取
			OrderDetailRefundList []OrderRefund `json:"orderDetailRefundList"`
			// 订单物流信息，发货前也会展示物流轨迹进，详情见新增订单发货前物流轨迹
			OrderLogisticsInfo []OrderLogisticsInfo `json:"orderLogisticsInfo"`
			// 订单当前生效的收货地址信息
			OrderAddress OrderAddress `json:"orderAddress"`
			// 多包裹发货信息
			OrderDeliveryInfo *OrderDeliveryInfo `json:"orderDeliveryInfo"`
			// 分销信息
			OrderCpsInfo *OrderCpsInfo `json:"orderCpsInfo,omitempty"`
		} `json:"data"`
	}
	OrderDetailBaseInfo struct {
		DiscountFee                   int                 `json:"discountFee"`
		BuyerNick                     string              `json:"buyerNick"`
		PayTime                       int64               `json:"payTime"`
		Channel                       string              `json:"channel"`
		OutSubBiz                     string              `json:"outSubBiz"`
		RemindShipmentSign            int                 `json:"remindShipmentSign"`
		SellerOpenID                  string              `json:"sellerOpenId"`
		ExpressFee                    int                 `json:"expressFee"`
		OrderSellerRoleInfo           OrderSellerRoleInfo `json:"orderlSellerRoleInfo"`
		BuyerImage                    string              `json:"buyerImage"`
		PayType                       int                 `json:"payType"`
		MultiplePiecesNo              int                 `json:"multiplePiecesNo"`
		OrderDomainCode               string              `json:"orderDomainCode"`
		ExpressCode                   string              `json:"expressCode"`
		EnableSplitDeliveryOrder      bool                `json:"enableSplitDeliveryOrder"`
		ValidPromiseShipmentTimeStamp int64               `json:"validPromiseShipmentTimeStamp"`
		GovernmentDiscount            int                 `json:"governmentDiscount"`
		DisableDeliveryReasonCode     []any               `json:"disableDeliveryReasonCode"`
		CpsType                       int                 `json:"cpsType"`
		PromiseTimeStampOfDelivery    int64               `json:"promiseTimeStampOfDelivery"`
		RiskCode                      int                 `json:"riskCode"`
		OutBizTradeNo                 string              `json:"outBizTradeNo"`
		TheDayOfDeliverGoodsTime      int                 `json:"theDayOfDeliverGoodsTime"`
		CommentStatus                 int                 `json:"commentStatus"`
		SendTime                      int64               `json:"sendTime"`
		TradeInPayAfterPromoAmount    int                 `json:"tradeInPayAfterPromoAmount"`
		RemindShipmentTime            int                 `json:"remindShipmentTime"`
		AllowanceExpressFee           int                 `json:"allowanceExpressFee"`
		Status                        int                 `json:"status"`
		OrderLabels                   []any               `json:"orderLabels"`
		Remark                        string              `json:"remark"`
		Oid                           int64               `json:"oid"`
		ShopNewBuyer                  bool                `json:"shopNewBuyer"`
		PrivacyInfoAuthTime           int                 `json:"privacyInfoAuthTime"`
		SellerNick                    string              `json:"sellerNick"`
		RecvTime                      int64               `json:"recvTime"`
		BuyerOpenID                   string              `json:"buyerOpenId"`
		RefundTime                    int                 `json:"refundTime"`
		CarrierType                   int                 `json:"carrierType"`
		OrderTaxInfo                  OrderDetailTaxInfo  `json:"orderTaxInfo"`
		PlatformNewBuyer              bool                `json:"platformNewBuyer"`
		UpdateTime                    int64               `json:"updateTime"`
		OrderRiskEventInfo            []any               `json:"orderRiskEventInfo"`
		PreSale                       int                 `json:"preSale"`
		PrivacyInfoAuthStatus         int                 `json:"privacyInfoAuthStatus"`
		CoType                        int                 `json:"coType"`
		CreateTime                    int64               `json:"createTime"`
		TotalFee                      int                 `json:"totalFee" comment:"订单商品总价，不含运费"`
		AllActivityType               []int               `json:"allActivityType"`
		SellerDelayPromiseTimeStamp   int64               `json:"sellerDelayPromiseTimeStamp"`
		PayChannel                    string              `json:"payChannel"`
		ActivityType                  int                 `json:"activityType"`
		CarrierID                     int                 `json:"carrierId"`
		PriorityDelivery              bool                `json:"priorityDelivery"`
		PayChannelDiscount            int                 `json:"payChannelDiscount"`
	}
	OrderDetailTaxInfo struct {
		SellerBearAmount int `json:"sellerBearAmount"`
		TaxAmount        int `json:"taxAmount"`
	}
)

var _ Response = RespGetOrderDetail{}
