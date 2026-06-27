package ks_sdk

// 列表排序方式（sort）
const (
	SortAsc  int = 1 // 升序
	SortDesc int = 2 // 降序
)

// 订单查询状态（orderViewStatus / status）
const (
	OrderViewStatusAll        int = iota + 1 // 全部
	OrderViewStatusWaitPay                   // 待付款
	OrderViewStatusWaitSend                  // 待发货
	OrderViewStatusDelivering                // 待收货（已发货）
	OrderViewStatusSign                      // 已收货
	OrderViewStatusSuccess                   // 交易成功订单
	OrderViewStatusClose                     // 已关闭订单
)

// 订单状态
const (
	OrderStatusAll        int = 0  // 未知状态
	OrderStatusWaitPay    int = 10 // 待付款
	OrderStatusPay        int = 30 // 已付款
	OrderStatusDelivering int = 40 // 已发货
	OrderStatusSign       int = 50 // 已签收
	OrderStatusSuccess    int = 70 // 订单成功
	OrderStatusFailed     int = 80 // 订单失败，订单取消后会转为“订单失败”状态
)

// 订单列表查询时间类型（queryType）
const (
	OrderQueryTypeCreate int = 1 // 按创建时间查询
	OrderQueryTypeUpdate int = 2 // 按更新时间查询
)

// 分销类型（cpsType）
const (
	CpsTypeNone    int = 0 // 非分销
	CpsTypeDistrib int = 1 // 分销订单
)

// 退款方式
const (
	AfterSalesHandlingWay1  int = 1  // 退货退款
	AfterSalesHandlingWay10 int = 10 // 仅退款
	AfterSalesHandlingWay3  int = 3  // 换货
	AfterSalesHandlingWay4  int = 4  // 补寄
	AfterSalesHandlingWay5  int = 5  // 维修
)

// 订单退款状态
const (
	AfterSalesStatus10 int = 10 // 买家已经申请退款，等待卖家同意
	AfterSalesStatus12 int = 12 // 卖家已拒绝，等待买家处理
	AfterSalesStatus20 int = 20 // 协商纠纷，等待平台处理
	AfterSalesStatus30 int = 30 // 卖家已经同意退款，等待买家退货
	AfterSalesStatus40 int = 40 // 买家已经退货，等待卖家确认收货
	AfterSalesStatus45 int = 45 // 卖家已经发货，等待买家确认收货
	AfterSalesStatus50 int = 50 // 卖家已经同意退款，等待系统执行退款
	AfterSalesStatus60 int = 60 // 退款成功
	AfterSalesStatus70 int = 70 // 退款关闭
)

// 售后事件类型
const (
	AfterSalesEventType0  int = 0  // 未知
	AfterSalesEventType2  int = 2  // 卖家拒绝退款
	AfterSalesEventType3  int = 3  // 卖家同意退款
	AfterSalesEventType4  int = 4  // 平台或买家修改退款记录
	AfterSalesEventType5  int = 5  // 卖家编辑后买家同意修改
	AfterSalesEventType6  int = 6  // 提交退货信息
	AfterSalesEventType7  int = 7  // 卖家确认收货
	AfterSalesEventType50 int = 50 // 退款打款中
	AfterSalesEventType60 int = 60 // 退款成功
	AfterSalesEventType70 int = 70 // 退款失败
)

// 退款单请求类型
const (
	RefundType8 int = 8 // 等待退款
	RefundType9 int = 9 // 全部退款订单
)
