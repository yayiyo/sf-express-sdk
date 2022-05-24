package sf

var sfTimeTmp = "2006-01-02 15:04:05"

//
const (
	// 统一接入平台校验成功，调用后端服务成功；
	// 注意：不代表后端业务处理成功，实际业务处理结果，
	// 需要查看响应属性apiResultData中的详细结果
	serviceApiResultCodeA1000 = "A1000"
	// 必传参数不可为空
	serviceApiResultCodeA1001 = "A1001"
	// 请求时效已过期
	serviceApiResultCodeA1002 = "A1002"
	// IP无效
	serviceApiResultCodeA1003 = "A1003"
	//	无对应服务权限
	serviceApiResultCodeA1004 = "A1004"
	// 流量受控
	serviceApiResultCodeA1005 = "A1005"
	// 数字签名无效
	serviceApiResultCodeA1006 = "A1006"
	// 重复请求
	serviceApiResultCodeA1007 = "A1007"
	// 数据解密失败
	serviceApiResultCodeA1008 = "A1008"
	// 目标服务异常或不可达
	serviceApiResultCodeA1009 = "A1009"
	// 系统异常
	serviceApiResultCodeA1099 = "A1099"
)

var resultCodeText = map[string]string{
	serviceApiResultCodeA1000: "Success",
	serviceApiResultCodeA1001: "Required parameters cannot be empty",            // 必传参数不可为空
	serviceApiResultCodeA1002: "The request time limit has expired",             // 请求时效已过期
	serviceApiResultCodeA1003: "Invalid IP",                                     // IP无效
	serviceApiResultCodeA1004: "There are no corresponding service permissions", // 无对应服务权限
	serviceApiResultCodeA1005: "Traffic is controlled",                          // 流量受控
	serviceApiResultCodeA1006: "The digital signature is invalid",               // 数字签名无效
	serviceApiResultCodeA1007: "Repeat request",                                 // 重复请求
	serviceApiResultCodeA1008: "Data decryption failed",                         // 数据解密失败
	serviceApiResultCodeA1009: "The target service is abnormal or unreachable",  // 目标服务异常或不可达
	serviceApiResultCodeA1099: "System exception",                               // 系统异常
}

// ServiceCode
const (
	// 通用寄件类
	serviceCodeCreateOrder         = "EXP_RECE_CREATE_ORDER"          // 下订单
	serviceCodeSearchRoutes        = "EXP_RECE_SEARCH_ROUTES"         // 路由查询
	serviceCodeUpdateOrder         = "EXP_RECE_UPDATE_ORDER"          // 订单确认/取消
	serviceCodeQuerySFWayBill      = "EXP_RECE_QUERY_SFWAYBILL"       // 清单运费查询
	serviceCodeGetSubMailNo        = "EXP_RECE_GET_SUB_MAILNO"        // 子单号申请
	serviceCodeSearchOrderResp     = "EXP_RECE_SEARCH_ORDER_RESP"     // 订单结果查询
	serviceCodeDeliveryNotice      = "EXP_RECE_DELIVERY_NOTICE"       // 派送通知
	serviceCodeCreateExchangeOrder = "EXP_RECE_CREATE_EXCHANGE_ORDER" // 换货下单
	serviceCodeWantedIntercept     = "EXP_RECE_WANTED_INTERCEPT"      // 截单转寄退回
	serviceCodeCreateReverseOrder  = "EXP_RECE_CREATE_REVERSE_ORDER"  // 仓配退货下单
	serviceCodeCancelReverseOrder  = "EXP_RECE_CANCEL_REVERSE_ORDER"  // 仓配退货消单
	serviceCodePreOrder            = "EXP_RECE_PRE_ORDER"             // 预下单

	// 服务查询类
	serviceCodeFilterOrderBSP     = "EXP_RECE_FILTER_ORDER_BSP"     // 订单筛选
	serviceCodeValidateWayBillNo  = "EXP_RECE_VALIDATE_WAYBILLNO"   // 运单号合法性校验
	serviceCodeQueryGISDepartment = "EXP_RECE_QUERY_GIS_DEPARTMENT" // 服务网点查询
)

// Error Code https://open.sf-express.com/developSupport/976720?activeIndex=146623
const (
	// 寄件人信息
	errorCodeShipperAddress     = 1010 // 寄件地址不能为空
	errorCodeShipperContactName = 1010 // 寄件联系人不能为空
	errorCodeShipperContactTel  = 1012 // 寄件电话不能为空

	// 收件人信息
	errorCodeReceiverAddress     = 1014 // 到件地址不能为空
	errorCodeReceiverContactName = 1015 // 到件联系人不能为空
	errorCodeReceiverContactTel  = 1016 // 到件电话不能为空

	// 出口件信息
	errorCodeInternationalShipmentsPostCode          = 1020 // 出口件邮编不能为空
	errorCodeInternationalShipmentsCommodityName     = 1023 // 拖寄物品名不能为空
	errorCodeInternationalShipmentsCommodityQuantity = 1028 // 出口件时，拖寄物数量不能为空
	errorCodeInternationalShipmentsDeclaredValue     = 1038 // 出口件声明价值不能为空

	// 订单信息
	errorCodeOrderMonthlyCardNumber     = 6126 // 月结卡号不合法
	errorCodeOrderAVSName               = 6127 // 增值服务名不能为空
	errorCodeOrderAVSNameInvalid        = 6128 // 增值服务名不合法
	errorCodeOrderVolumeInvalid         = 6130 // 体积参数不合法
	errorCodeOrderCODAmountInvalid      = 6138 // 代收货款金额传入错误
	errorCodeOrderCODAmountLessThanZero = 6139 // 代收货款金额小于零错误
	errorCodeOrderNotFound              = 6150 // 找不到该订单

	// 国际件
	errorCodeInternationalShipmentsShipperPostCode  = 6200 // 国际件寄方邮编不能为空
	errorCodeInternationalShipmentsReceiverPostCode = 6201 // 国际件到方邮编不能为空
	errorCodeInternationalShipmentsCargoQuantity    = 6202 // 国际件货物数量不能为空
	errorCodeInternationalShipmentsCargoUnit        = 6203 // 国际件货物单位不能为空
	errorCodeInternationalShipmentsCargoUnitWeight  = 6204 // 国际件货物单位重量不能为空
	errorCodeInternationalShipmentsCargoUnitValue   = 6205 // 国际件货物单价不能为空
	errorCodeInternationalShipmentsCargoCurrency    = 6206 // 国际件货物币种不能为空
	errorCodeInternationalShipmentsOriginCode       = 6207 // 国际件原产地不能为空

	// 业务
	errorCodeAWBNumberExceedsLimit               = 8003  // 查询单号超过最大限制
	errorCodeAWBNumber                           = 8013  // 未传入查询单号
	errorCodeDuplicatedOrderID                   = 8016  // 重复下单
	errorCodeOrderAndAWBNumberNotMatch           = 8017  // 订单号与运单号不匹配
	errorCodeOrderInformation                    = 8018  // 未获取到订单信息
	errorCodeOrderConfirmed                      = 8252  // 订单已确认
	errorCodeOrderCancelled                      = 8037  // 订单已取消
	errorCodeBusinessTemplateNotFound            = 8027  // 不存在的业务模板
	errorCodeShipmentTypeInvalid                 = 8057  // 快件类型为空或未配置
	errorCodeExceedMaxNumberOfSubWaybills        = 8067  // 超过最大能申请子单号数量
	errorCodeAppointmentTimeInvalid              = 8096  // 您的预约超出今日营业时间，无法上门收件。
	errorCodeMonthlyCardCreateOrderForbidden     = 8114  // 传入了不可发货的月结卡号
	errorCodeExceedMaxNumberOfOrders             = 8117  // 下单包裹不能大于307个
	errorCodeMonthlyCardInvalid                  = 8119  // 月结卡号不存在或已失效
	errorCodeApplySubWaybillsForbidden           = 8168  // 订单已生成路由不能申请子单
	errorCodeAWBNumberNotRight                   = 8191  // 运单号格式不正确
	errorCodeConsValueCurrencyCodeRequired       = 8194  // 跨境件必须包含申明价值和币别
	errorCodeInformationException                = 8196  // 信息异常
	errorCodeAWBNumberInvalid                    = 8247  // 运单号不合法
	errorCodePaymentMethodsNotSupported          = 8256  // 部分快件产品不支持到付和寄付现结，请调整付款方式后下单
	errorCodeScheduledDeliveryNotSupported       = 8051  // 定时派送不在时效范围内，下单失败
	errorCodeSourceScheduledDeliveryNotSupported = 8052  // 原寄地不在定时派送服务范围内
	errorCodeDestScheduledDeliveryNotSupported   = 8053  // 目的地不在定时派送服务范围内
	errorCodeBookingsNumberOverloaded            = 8177  // 类似 （正值运力高峰期，普通会员（非会员）的寄件通道预约已满，敬请谅解） 提示语组成 BPS：策略编号
	errorCodeAgreementNotFound                   = 8179  // 卡号下未查到关联相应协议
	errorCodeScheduledDeliveryTimeNotMatch       = 20011 // 产品与定时派送服务时间段不匹配
	errorCodeScheduledDeliveryOverweight         = 20012 // 定时派送服务不支持重量超过300KG的快件
	errorCodeProhibitedItemsForbidden            = 20035 // 托寄物违禁品不可收寄  Prohibited items cannot be received or sent
	errorCodeProductForbidden                    = 20036 // 适用产品不满足
)

var errorSuggestion = map[int]string{}

func suggestionMsg(code int) string {
	if msg, ok := errorSuggestion[code]; ok {
		return msg
	}
	return ""
}

const (
	errorCodeS0000 = "S0000" // 成功
	errorCodeS0001 = "S0001" // 非法的JSON格式
	errorCodeS0002 = "S0002" // 必填参数%s为空
	errorCodeS0003 = "S0003" // 系统发生数据错误或运行时异常
	errorCodeS0004 = "S0004" // 参数%s超过最大长度%d
	errorCodeS0005 = "S0005" // 参数超过最大值
	errorCodeS0006 = "S0006" // 参数%s不能小于%d
	errorCodeS0007 = "S0007" // 参数%s数据类型错误
)
