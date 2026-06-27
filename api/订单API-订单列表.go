package api

import "net/http"

// 文档：https://open.kwaixiaodian.com/zone/new/docs/api?name=open.order.cursor.list&version=1

type ReqGetOrderList struct {
	// OrderViewStatus 订单状态，不传查询全部。常见：1=待付款 2=待发货 ... 详见常量
	OrderViewStatus int `json:"orderViewStatus,omitempty"`
	// PageSize 每页数量，最大 100
	PageSize int `json:"pageSize,omitempty"`
	// BeginTime 查询开始时间，毫秒级时间戳，不能小于90天前，且需要小于结束时间
	BeginTime int64 `json:"beginTime,omitempty"`
	// EndTime 查询结束时间，毫秒级时间戳，与开始时间的时间范围不大于7天
	EndTime int64 `json:"endTime,omitempty"`
	// Cursor 游标，第一页传空字符串或 "0"，后续传上一页返回的 cursor
	Cursor string `json:"cursor"`
	// Sort 排序，1=升序 2=降序
	Sort int `json:"sort,omitempty"`
	// QueryType 时间类型，1=按创建时间 2=按更新时间
	QueryType int `json:"queryType,omitempty"`
	// CpsType 分销类型 0-全部 1-普通订单 2-分销订单
	CpsType int `json:"cpsType"`
}

var _ Request = ReqGetOrderList{}

func (r ReqGetOrderList) HttpMethod() string {
	return http.MethodGet
}

func (r ReqGetOrderList) Method() string {
	return "open.order.cursor.list"
}

func (r ReqGetOrderList) Param() any {
	return r
}

type (
	RespGetOrderList struct {
		BaseResponse
		Data struct {
			Cursor    string        `json:"cursor"`
			OrderList []OrderDetail `json:"orderList"`
			PageSize  int           `json:"pageSize"`
			BeginTime int64         `json:"beginTime"`
			EndTime   int64         `json:"endTime"`
		} `json:"data"`
	}
	OrderSellerRoleInfo struct {
		RoleID   int64  `json:"roleId"`
		RoleName string `json:"roleName"`
		RoleType int    `json:"roleType"`
	}
	OrderBaseInfo struct {
		DiscountFee                   int                 `json:"discountFee"`
		BuyerNick                     string              `json:"buyerNick"`
		PayTime                       int64               `json:"payTime"`
		OrderLabels                   []any               `json:"orderLabels"`
		Channel                       string              `json:"channel"`
		Remark                        string              `json:"remark"`
		RemindShipmentSign            int                 `json:"remindShipmentSign"`
		Oid                           int64               `json:"oid"`
		SellerOpenID                  string              `json:"sellerOpenId"`
		ExpressFee                    int                 `json:"expressFee" comment:"运费"`
		OrderSellerRoleInfo           OrderSellerRoleInfo `json:"orderSellerRoleInfo"`
		BuyerImage                    string              `json:"buyerImage"`
		PayType                       int                 `json:"payType"`
		MultiplePiecesNo              int                 `json:"multiplePiecesNo"`
		OrderDomainCode               string              `json:"orderDomainCode"`
		ExpressCode                   string              `json:"expressCode"`
		EnableSplitDeliveryOrder      bool                `json:"enableSplitDeliveryOrder"`
		ValidPromiseShipmentTimeStamp int64               `json:"validPromiseShipmentTimeStamp"`
		GovernmentDiscount            int                 `json:"governmentDiscount"`
		SellerNick                    string              `json:"sellerNick"`
		DisableDeliveryReasonCode     []any               `json:"disableDeliveryReasonCode"`
		RecvTime                      int                 `json:"recvTime"`
		BuyerOpenID                   string              `json:"buyerOpenId"`
		CpsType                       int                 `json:"cpsType"`
		PromiseTimeStampOfDelivery    int64               `json:"promiseTimeStampOfDelivery"`
		RefundTime                    int                 `json:"refundTime"`
		RiskCode                      int                 `json:"riskCode"`
		UpdateTime                    int64               `json:"updateTime"`
		OrderRiskEventInfo            []any               `json:"orderRiskEventInfo"`
		TheDayOfDeliverGoodsTime      int                 `json:"theDayOfDeliverGoodsTime"`
		CommentStatus                 int                 `json:"commentStatus"`
		SendTime                      int64               `json:"sendTime"`
		TradeInPayAfterPromoAmount    int                 `json:"tradeInPayAfterPromoAmount"`
		PreSale                       int                 `json:"preSale"`
		CoType                        int                 `json:"coType"`
		CreateTime                    int64               `json:"createTime"`
		TotalFee                      int                 `json:"totalFee"`
		AllActivityType               []int               `json:"allActivityType"`
		SellerDelayPromiseTimeStamp   int64               `json:"sellerDelayPromiseTimeStamp"`
		PayChannel                    string              `json:"payChannel"`
		RemindShipmentTime            int                 `json:"remindShipmentTime"`
		ActivityType                  int                 `json:"activityType"`
		AllowanceExpressFee           int                 `json:"allowanceExpressFee"`
		PriorityDelivery              bool                `json:"priorityDelivery"`
		PayChannelDiscount            int                 `json:"payChannelDiscount"`
		Status                        int                 `json:"status"`
	}
	OrderAddress struct {
		DistrictCode         int    `json:"districtCode"`
		Town                 string `json:"town"`
		City                 string `json:"city"`
		TownCode             int    `json:"townCode"`
		CityCode             int    `json:"cityCode"`
		ProvinceCode         int    `json:"provinceCode"`
		EncryptedMobile      string `json:"encryptedMobile"`
		EncryptedConsignee   string `json:"encryptedConsignee"`
		DesensitiseConsignee string `json:"desensitiseConsignee"`
		EncryptedAddress     string `json:"encryptedAddress"`
		Province             string `json:"province"`
		District             string `json:"district"`
		DesensitiseMobile    string `json:"desensitiseMobile"`
		DesensitiseAddress   string `json:"desensitiseAddress"`
	}
	OrderLogisticsInfo struct {
		LogisticsID int64  `json:"logisticsId"`
		ExpressNo   string `json:"expressNo"`
		ExpressCode int    `json:"expressCode"`
	}
	ItemExtra struct {
		BrandName    string       `json:"brandName"`
		EnergyLevel  string       `json:"energyLevel"`
		CategoryInfo CategoryInfo `json:"categoryInfo"`
		ProductNo    string       `json:"productNo"`
	}
	CategoryInfo struct {
		GovCategory     string `json:"govCategory"`
		ItemCid         int    `json:"itemCid"`
		GovCategoryCode string `json:"govCategoryCode"`
		CategoryName    string `json:"categoryName"`
	}
	OrderItemInfo struct {
		ItemPicURL    string    `json:"itemPicUrl"`
		ItemType      int       `json:"itemType"`
		DiscountFee   int       `json:"discountFee"`
		OriginalPrice int       `json:"originalPrice"`
		ItemTitle     string    `json:"itemTitle"`
		OrderItemID   int64     `json:"orderItemId"`
		Num           int       `json:"num"`
		ItemExtra     ItemExtra `json:"itemExtra"`
		WarehouseCode string    `json:"warehouseCode"`
		ItemID        int64     `json:"itemId"`
		RelItemID     int64     `json:"relItemId"`
		RelSkuID      int       `json:"relSkuId"`
		Price         int       `json:"price"`
		ItemLinkURL   string    `json:"itemLinkUrl"`
		SkuNick       string    `json:"skuNick" comment:"商家商品编码"`
		SkuDesc       string    `json:"skuDesc"`
		GoodsCode     string    `json:"goodsCode"`
		SkuID         int64     `json:"skuId"`
	}
	OrderDeliveryInfo struct {
		SplitDeliveryOrder  bool   `json:"splitDeliveryOrder"`
		MergeDeliveryType   int    `json:"mergeDeliveryType"`
		DeliveryNum         int    `json:"deliveryNum"`
		EnableAppendPackage bool   `json:"enableAppendPackage"`
		TotalPackageNum     int    `json:"totalPackageNum"`
		OpenAddressID       string `json:"openAddressId"`
		MaxPackageNum       int    `json:"maxPackageNum"`
		DeliveryStatus      int    `json:"deliveryStatus"`
		PackageNum          int    `json:"packageNum"`
	}
	OrderCpsInfo struct {
		KwaiMoneyUserID       int    `json:"kwaiMoneyUserId"`
		DistributorName       string `json:"distributorName"`
		ActivityUserID        int    `json:"activityUserId"`
		DistributorID         int    `json:"distributorId"`
		ActivityUserNickName  string `json:"activityUserNickName"`
		KwaiMoneyUserNickName string `json:"kwaiMoneyUserNickName"`
	}
	OrderRefund struct {
		RefundID          int64 `json:"refundId"`
		RefundStatus      int   `json:"refundStatus"`
		HandlingWay       int   `json:"handlingWay"`
		SpecialRefundType int   `json:"specialRefundType"`
	}
	OrderDetail struct {
		// 订单基础信息
		OrderBaseInfo OrderBaseInfo `json:"orderBaseInfo"`
		// 订单商品信息
		OrderItemInfo OrderItemInfo `json:"orderItemInfo"`
		// 订单退款信息，返回最新的一个，若要获取当前订单关联所有退款单，可通过退款单列表API获取
		OrderRefundList []OrderRefund `json:"orderRefundList"`
		// 订单物流信息，发货前也会展示物流轨迹进，详情见新增订单发货前物流轨迹
		OrderLogisticsInfo []OrderLogisticsInfo `json:"orderLogisticsInfo"`
		// 订单当前生效的收货地址信息
		OrderAddress OrderAddress `json:"orderAddress"`
		// 多包裹发货信息
		OrderDeliveryInfo *OrderDeliveryInfo `json:"orderDeliveryInfo"`
		// 分销信息
		OrderCpsInfo *OrderCpsInfo `json:"orderCpsInfo,omitempty"`
	}
)

var _ Response = RespGetOrderList{}
