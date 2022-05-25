package sf

// 必填 条件：C 是：T 否：F
// 《增值服务产品表》 https://open.sf-express.com/developSupport/734349?activeIndex=313317
// 《快件产品类别表》 https://open.sf-express.com/developSupport/734349?activeIndex=324604
// 《城市代码表》 https://open.sf-express.com/developSupport/734349?activeIndex=571489
// 《国际件币别表》 https://open.sf-express.com/developSupport/734349?activeIndex=482730

type CreateOrderReq struct {
	Language                   string           `json:"language"`                   // T 响应报文的语言， 缺省值为zh-CN， 目前支持以下值zh-CN 表示中文简体， zh-TW或zh-HK或 zh-MO表示中文繁体， en表示英文
	OrderId                    string           `json:"orderId"`                    // T 客户订单号，不能重复
	WaybillNoInfoList          []*WaybillNoInfo `json:"waybillNoInfoList"`          // F 顺丰运单号
	CustomsInfo                *CustomsInfo     `json:"customsInfo"`                // F 报关信息
	CargoDetails               []*CargoDetail   `json:"cargoDetails"`               // T 托寄物信息
	CargoDesc                  string           `json:"cargoDesc"`                  // F 拖寄物类型描述,如： 文件，电子产品，衣服等
	ExtraInfoList              []*ExtraInfo     `json:"extraInfoList"`              // F 扩展属性
	ServiceList                []*Service       `json:"serviceList"`                // F 增值服务信息，支持附录： 《增值服务产品表》
	ContactInfoList            []*ContactInfo   `json:"contactInfoList"`            // T 收寄双方信息
	MonthlyCard                string           `json:"monthlyCard"`                // C 顺丰月结卡号 月结支付时传值，现结不需传值；沙箱联调可使用测试月结卡号7551234567（非正式，无须绑定，仅支持联调使用）
	PayMethod                  int              `json:"payMethod"`                  // F 付款方式，支持以下值： 1:寄方付 2:收方付 3:第三方付
	ExpressTypeId              int              `json:"expressTypeId"`              // F 快件产品类别， 支持附录 《快件产品类别表》 的产品编码值，仅可使用与顺丰销售约定的快件产品
	ParcelQty                  int              `json:"parcelQty"`                  // F 包裹数，一个包裹对应一个运单号；若包裹数大于1，则返回一个母运单号和N-1个子运单号
	TotalLength                float64          `json:"totalLength"`                // F 客户订单货物总长，单位厘米， 精确到小数点后3位， 包含子母件
	TotalWidth                 float64          `json:"totalWidth"`                 // F 客户订单货物总宽，单位厘米， 精确到小数点后3位， 包含子母件
	TotalHeight                float64          `json:"totalHeight"`                // F 客户订单货物总高，单位厘米， 精确到小数点后3位， 包含子母件
	Volume                     float64          `json:"volume"`                     // F 订单货物总体积，单位立方厘米, 精确到小数点后3位，会用于计抛 (是否计抛具体商务沟通中 双方约定)
	TotalWeight                float64          `json:"totalWeight"`                // C 订单货物总重量（郑州空港海关必填）， 若为子母件必填， 单位千克， 精确到小数点后3位，如果提供此值， 必须>0 (子母件需>6)
	TotalNetWeight             float64          `json:"totalNetWeight"`             // F 商品总净重
	SendStartTm                string           `json:"sendStartTm"`                // F 要求上门取件开始时间， 格式： YYYY-MM-DD HH24:MM:SS， 示例： 2012-7-30 09:30:00
	IsDocall                   int              `json:"isDocall"`                   // F 是否通过手持终端 通知顺丰收派 员上门收件，支持以下值： 1：要求 0：不要求
	IsSignBack                 int              `json:"isSignBack"`                 // F 是否返回签回单 （签单返还）的运单号， 支持以下值： 1：要求 0：不要求
	CustReferenceNo            string           `json:"custReferenceNo"`            // F 客户参考编码：如客户原始订单号
	TemperatureRange           int              `json:"temperatureRange"`           // C 温度范围类型，当 express_type为12 医药温控件 时必填，支持以下值： 1:冷藏 3：冷冻
	OrderSource                string           `json:"orderSource"`                // F 订单平台类型 （对于平台类客户， 如果需要在订单中 区分订单来源， 则可使用此字段） 天猫:tmall， 拼多多：pinduoduo， 京东 : jd 等平台类型编码
	Remark                     string           `json:"remark"`                     // F 备注
	IsOneselfPickup            int              `json:"isOneselfPickup"`            // F 快件自取，支持以下值： 1：客户同意快件自取 0：客户不同意快件自取
	FilterField                string           `json:"filterField"`                // F 筛单特殊字段用来人工筛单
	IsReturnQRCode             int              `json:"isReturnQRCode"`             // F 是否返回用来退货业务的 二维码URL， 支持以下值： 1：返回二维码 0：不返回二维码
	SpecialDeliveryTypeCode    string           `json:"specialDeliveryTypeCode"`    // F 特殊派送类型代码 1:身份验证
	SpecialDeliveryValue       string           `json:"specialDeliveryValue"`       // F 特殊派件具体表述 证件类型: 证件后8位如： 1:09296231（1 表示身份证， 暂不支持其他证件）
	MerchantPayOrderNo         string           `json:"merchantPayOrderNo"`         // F 商户支付订单号
	IsReturnSignBackRouteLabel int              `json:"isReturnSignBackRouteLabel"` // F 是否返回签回单路由标签： 默认0， 1：返回路由标签， 0：不返回
	IsReturnRouteLabel         int              `json:"isReturnRouteLabel"`         // T 是否返回路由标签： 默认0， 1：返回路由标签， 0：不返回；除部分特殊用户外，其余用户都默认返回
	IsUnifiedWaybillNo         int              `json:"isUnifiedWaybillNo"`         // F 是否使用国家统一面单号 1：是， 0：否（默认）
	PodModelAddress            string           `json:"podModelAddress"`            // F 签单返还范本地址
	InProcessWaybillNo         string           `json:"inProcessWaybillNo"`         // F 头程运单号（郑州空港海关必填）
	IsGenWaybillNo             int              `json:"isGenWaybillNo"`             // F 是否需求分配运单号1：分配 0：不分配（若带单号下单，请传值0）
}

type T struct {
	WaybillNoInfoList []*WaybillNoInfo  `json:"waybillNoInfoList"`
	RouteLabelInfo    []*RouteLabelInfo `json:"routeLabelInfo"`
	ContactInfoList   interface{}       `json:"contactInfoList"`
	SendStartTm       interface{}       `json:"sendStartTm"`
	CustomerRights    interface{}       `json:"customerRights"`
	ExpressTypeId     interface{}       `json:"expressTypeId"`
}

type Order struct {
	OrderId                   string            `json:"orderId"`                   // T 客户订单号
	OriginCode                string            `json:"originCode"`                // F 原寄地区域代码，可用于顺丰 电子运单标签打印
	DestCode                  string            `json:"destCode"`                  // F 目的地区域代码，可用于顺丰 电子运单标签打印
	FilterResult              int               `json:"filterResult"`              // F 筛单结果： 1：人工确认 2：可收派 3：不可以收派
	Remark                    string            `json:"remark"`                    // C 如果filter_result=3时为必填， 不可以收派的原因代码： 1：收方超范围 2：派方超范围 3：其它原因 高峰管控提示信息 【数字】：【高峰管控提示信息】 (如 4：温馨提示 ，1：春运延时)
	URL                       int64             `json:"url"`                       // F 二维码URL （用于CX退货操作的URL）
	PaymentLink               string            `json:"paymentLink"`               // F 用于第三方支付运费的URL
	IsUpstairs                string            `json:"isUpstairs"`                // F 是否送货上楼 1:是
	IsSpecialWarehouseService string            `json:"isSpecialWarehouseService"` // F   true 包含特殊仓库增值服务
	ServiceList               []*Service        `json:"serviceList"`               // F 下单补充的增值服务信息
	ReturnExtraInfoList       []*ExtraInfo      `json:"returnExtraInfoList"`       // F 返回信息扩展属性
	WaybillNoInfoList         []*WaybillNoInfo  `json:"waybillNoInfoList"`         // F 顺丰运单号
	RouteLabelInfo            []*RouteLabelInfo `json:"routeLabelInfo"`            // T 路由标签，除少量特殊场景用户外，其余均会返回
	ContactInfoList           interface{}       `json:"contactInfoList"`
	SendStartTm               interface{}       `json:"sendStartTm"`
	CustomerRights            interface{}       `json:"customerRights"`
	ExpressTypeId             interface{}       `json:"expressTypeId"`
	MappingMark               interface{}       `json:"mappingMark"`
	AgentMailno               interface{}       `json:"agentMailno"`
}

type RouteLabelInfo struct {
	Code           string          `json:"code"`           // T	返回调用结果，1000：调用成功； 其他调用失败
	RouteLabelData *RouteLabelData `json:"routeLabelData"` // T	路由标签数据详细数据，除少量特殊场景用户外，其余均会返回
	Message        string          `json:"message"`        // F   失败异常描述
}

type RouteLabelData struct {
	WaybillNo              string      `json:"waybillNo"`           // F 运单号
	SourceTransferCode     string      `json:"sourceTransferCode"`  // F 原寄地中转场
	SourceCityCode         string      `json:"sourceCityCode"`      // F 原寄地城市代码
	SourceDeptCode         string      `json:"sourceDeptCode"`      // F 原寄地网点代码
	SourceTeamCode         string      `json:"sourceTeamCode"`      // F 原寄地单元区域
	DestCityCode           string      `json:"destCityCode"`        // F 目的地城市代码, eg:755
	DestDeptCode           string      `json:"destDeptCode"`        // F 目的地网点代码, eg:755AQ
	DestDeptCodeMapping    string      `json:"destDeptCodeMapping"` // F 目的地网点代码映射码
	DestTeamCode           string      `json:"destTeamCode"`        // F 目的地单元区域, eg:001
	DestTeamCodeMapping    string      `json:"destTeamCodeMapping"` // F 目的地单元区域映射码
	DestTransferCode       string      `json:"destTransferCode"`    // F 目的地中转场
	DestRouteLabel         string      `json:"destRouteLabel"`      // T 若返回路由标签，则此项必会返回。如果手打是一段码，检查是否地址异常。打单时的路由标签信息如果是大网的路由标签, 这里的值是目的地网点代码, 如果 是同城配的路由标签, 这里的值是根据同城配的设置映射出来的值, 不同的配置结果会不一样, 不能根据-符号切分(如:上海同城配, 可能是:集散点-目的地网点-接驳点, 也有可能是目的地网点代码-集散点-接驳点)
	ProName                string      `json:"proName"`             // F 产品名称 对应RLS:pro_name
	CargoTypeCode          string      `json:"cargoTypeCode"`       // F 快件内容: 如:C816、SP601
	LimitTypeCode          string      `json:"limitTypeCode"`       // F 时效代码, 如:T4
	ExpressTypeCode        string      `json:"expressTypeCode"`     // F 产品类型, 如:B1
	CodingMapping          string      `json:"codingMapping"`       // T 入港映射码 eg:S10 地址详细必会返回
	CodingMappingOut       string      `json:"codingMappingOut"`    // F 出港映射码
	XBFlag                 string      `json:"xbFlag"`              // F XB标志 0:不需要打印XB 1:需要打印XB
	PrintFlag              string      `json:"printFlag"`           // F 打印标志 返回值总共有9位, 每位只 有0和1两种, 0表示按丰密 面单默认的规则, 1是显示, 顺序如下, 如111110000 表示打印寄方姓名、寄方 电话、寄方公司名、寄方 地址和重量, 收方姓名、收 方电话、收方公司和收方 地址按丰密面单默认规则 1:寄方姓名 2:寄方电话 3:寄方公司名 4:寄方地址 5:重量 6:收方姓名 7:收方电话 8:收方公司名 9:收方地址
	TwoDimensionCode       string      `json:"twoDimensionCode"`    // F 二维码 根据规则生成字符串信息, 格式为MMM = { ‘k1’:’(目的 地中转场代码)’, ‘k2’:’(目的 地原始网点代码)’, ‘k3’:’(目 的地单元区域)’, ‘k4’:’(附件 通过三维码(express_type_code、 limit_type_code、 cargo_type_code)映射时效类型)’, ‘k5’:’(运单 号)’, ‘k6’:’(AB标识)’, ‘k7’:’( 校验码)’}
	ProCode                string      `json:"proCode"`             // F 时效类型: 值为二维码中的K4
	PrintIcon              string      `json:"printIcon"`           // F 打印图标, 根据托寄物判断需 要打印的图标(重货, 蟹类, 生鲜, 易碎，Z标) 返回值有8位，每一位只有0和1两种， 0表示按运单默认的规则， 1表示显示。后面两位默认0备用。 顺序如下：重货, 蟹类, 生鲜, 易碎, 医药类, Z标, 酒标, 0 如：00000000表示不需要打印重货，蟹类，生鲜，易碎, 医药, Z标, 酒标, 备用
	ABFlag                 string      `json:"abFlag"`              // F AB标
	WaybillIconList        interface{} `json:"waybillIconList"`     // F 面单图标
	ErrMsg                 string      `json:"errMsg"`              // F 查询出现异常时返回信息。 返回代码: 0 系统异常 1 未找到面单
	DestPortCode           string      `json:"destPortCode"`        // F 目的地口岸代码
	DestCountry            string      `json:"destCountry"`         // F 目的国别(国别代码如:JP)
	DestPostCode           string      `json:"destPostCode"`        // F 目的地邮编
	GoodsValueTotal        float64     `json:"goodsValueTotal"`     // F 总价值(保留两位小数, 数字类型, 可补位)
	CurrencySymbol         string      `json:"currencySymbol"`      // F 币种
	GoodsNumber            int         `json:"goodsNumber"`         // F 件数
	DestAddrKeyWord        string      `json:"destAddrKeyWord"`     // F 目的地地址关键词
	XbFlag                 string      `json:"xbFlag"`
	CusBatch               string      `json:"cusBatch"`
	CheckCode              string      `json:"checkCode"`
	ProIcon                string      `json:"proIcon"`
	FileIcon               string      `json:"fileIcon"`
	FbaIcon                string      `json:"fbaIcon"`
	IcsmIcon               string      `json:"icsmIcon"`
	DestGisDeptCode        string      `json:"destGisDeptCode"`
	NewIcon                interface{} `json:"newIcon"`
	SendAreaCode           interface{} `json:"sendAreaCode"`
	DestinationStationCode interface{} `json:"destinationStationCode"`
	SxLabelDestCode        interface{} `json:"sxLabelDestCode"`
	SxDestTransferCode     interface{} `json:"sxDestTransferCode"`
	SxCompany              interface{} `json:"sxCompany"`
	NewAbFlag              interface{} `json:"newAbFlag"`
	RongType               interface{} `json:"rongType"`
}

type WaybillNoInfo struct {
	WaybillType int     `json:"waybillType"` // F 运单号类型1：母单 2 :子单 3 : 签回单
	WaybillNo   string  `json:"waybillNo"`   // F 运单号
	BoxNo       string  `json:"boxNo"`       // F 箱号
	Length      float64 `json:"length"`      // F 长(cm)
	Width       float64 `json:"width"`       // F 宽(cm)
	Height      float64 `json:"height"`      // F 高(cm)
	Weight      float64 `json:"weight"`      // F 重量(kg)
	Volume      float64 `json:"volume"`      // F 体积（立方厘米）
}

type ContactInfo struct {
	ContactType int    `json:"contactType"`       // T 地址类型： 1，寄件方信息 2，到件方信息
	Company     string `json:"company,omitempty"` // C 公司名称
	Contact     string `json:"contact"`           // C 联系人
	Tel         string `json:"tel"`               // C 联系电话
	Mobile      string `json:"mobile"`            // F 手机
	ZoneCode    string `json:"zoneCode"`          // C 城市代码或国家代码,参考附录 《城市代码表》，如果是跨境件，则此字段为必填
	Country     string `json:"country"`           // T 国家或地区2位代码 参照附录《城市代码表》
	Province    string `json:"province"`          // F 所在省级行政区名称，必须是标准的省级行政区名称如：北 京、广东省、广西壮族自治区等；此字段影响原寄地代码识 别，建议尽可能传该字段的值
	City        string `json:"city"`              // F 所在地级行政区名称，必须是标准的城市称谓 如：北京市、 深圳市、大理白族自治州等； 此字段影响原寄地代码识别， 建议尽可能传该字段的值
	County      string `json:"county"`            // F 所在县/区级行政区名称，必须 是标准的县/区称谓，如：福田区，南涧彝族自治县、准格尔旗等
	Address     string `json:"address"`           // T 详细地址，若有四级行政区划，如镇/街道等信息可拼接至此字段，格式样例：镇/街道+详细地址。若province/city 字段的值不传，此字段必须包含省市信息，避免影响原寄地代码识别，如：广东省深圳市福田区新洲十一街万基商务大厦10楼；此字段地址必须详细，否则会影响目的地中转识别；
	PostCode    string `json:"postCode"`          // C 邮编，跨境件必填（中国内地， 港澳台互寄除外）
	Email       string `json:"email"`             // F 邮箱地址
	TaxNo       string `json:"taxNo"`             // F 税号
}

type CargoDetail struct {
	Name                          string  `json:"name"`                          // T 货物名称，如果需要生成电子 运单，则为必填
	Count                         float64 `json:"count"`                         // C 货物数量 跨境件报关需要填写
	Unit                          string  `json:"unit"`                          // C 货物单位，如：个、台、本， 跨境件报关需要填写
	Weight                        float64 `json:"weight"`                        // C 订单货物单位重量，包含子母件， 单位千克，精确到小数点后3位 跨境件报关需要填写
	Amount                        float64 `json:"amount"`                        // C 货物单价，精确到小数点后3位， 跨境件报关需要填写
	Currency                      string  `json:"currency"`                      // C 货物单价的币别： 参照附录《国际件币别表》
	GoodPrepardNo                 string  `json:"goodPrepardNo"`                 // F 商品海关备案号
	HsCode                        string  `json:"hsCode"`                        // F 海关编码
	ProductRecordNo               string  `json:"productRecordNo"`               // F 货物产品国检备案编号
	SourceArea                    string  `json:"sourceArea"`                    // C 原产地国别， 跨境件报关需要填写
	TaxNo                         string  `json:"taxNo"`                         // F 商品行邮税号
	GoodsCode                     string  `json:"goodsCode"`                     // F 商品编号
	CargoSku                      string  `json:"cargoSku"`                      // F 商品编码
	VisualInspection              string  `json:"visualInspection"`              // F 外观
	Brand                         string  `json:"brand"`                         // F 货物品牌
	Specifications                string  `json:"specifications"`                // F 货物规格型号
	Manufacturer                  string  `json:"manufacturer"`                  // F 生产厂家
	ShipmentWeight                float64 `json:"shipmentWeight"`                // F 托寄物毛重
	Length                        float64 `json:"length"`                        // F 托寄物长
	Width                         float64 `json:"width"`                         // F 托寄物宽
	Height                        float64 `json:"height"`                        // F 托寄物高
	Volume                        float64 `json:"volume"`                        // F 托寄物体积
	CargoDeclaredValue            float64 `json:"cargoDeclaredValue"`            // F 托寄物声明价值（杭州口岸必填）
	DeclaredValueDeclaredCurrency string  `json:"declaredValueDeclaredCurrency"` // F 托寄物声明价值币别（杭州口岸必填）
	CargoID                       string  `json:"cargoId"`                       // F 货物id（逆向物流）
	IntelligentInspection         int     `json:"intelligentInspection"`         // F 智能验货标识(1-是,0-否) （逆向物流）
	SNCode                        string  `json:"snCode"`                        // F 货物标识码（逆向物流）
	CheckRemark                   string  `json:"checkRemark"`                   // F 收货备注，如果商品需要收派员验货则要填写这一项
	StateBarCode                  string  `json:"stateBarCode"`                  // F 国条码
}

type Service struct {
	Name   string `json:"name"`   // T 增值服务名，如COD等，支持附录： 《增值服务产品表》
	Value  string `json:"value"`  // C 增值服务扩展属性，参考增值 服务传值说明
	Value1 string `json:"value1"` // C 增值服务扩展属性1
	Value2 string `json:"value2"` // C 增值服务扩展属性2
	Value3 string `json:"value3"` // C 增值服务扩展属性3
	Value4 string `json:"value4"` // C 增值服务扩展属性4
}

type CustomsInfo struct {
	DeclaredValue         float64 `json:"declaredValue"`         // C 客户订单货物总声明价值， 包含子母件，精确到小数点 后3位。如果是跨境件，则必填
	DeclaredValueCurrency string  `json:"declaredValueCurrency"` // F 中国内地 默认CNY， 否则 默认USD	货物声明价值币别，跨境 件报关需要填写 参照附录币别代码附件
	CustomsBatchs         string  `json:"customsBatchs"`         // F 报关批次
	TaxPayMethod          int     `json:"taxPayMethod"`          // F 税金付款方式，支持以下值： 1:寄付 2：到付
	TaxSettleAccounts     string  `json:"taxSettleAccounts"`     // F 税金结算账号
	PaymentTool           string  `json:"paymentTool"`           // F 支付工具
	PaymentNumber         string  `json:"paymentNumber"`         // F 支付号码
	OrderName             string  `json:"orderName"`             // F 客户订单下单人姓名
	OrderCertType         string  `json:"orderCertType"`         // F 客户订单下单人证件类型
	OrderCertNo           string  `json:"orderCertNo"`           // F 客户订单下单人证件号
	Tax                   string  `json:"tax"`                   // F 税款
}

// CreateOrder
// attrName	attrVal
// attr001	物品件数（杭州口岸必填）
// attr002	物品净重，郑州空港口岸, 郑州综保税必须字段，最多保留小数点后四位，单位Kg
// attr003	进出口时间（郑州空港口岸必须字段格式：yyyy-MM-dd hh:mm:ss）
// attr005	商品运输费用,无则填0（郑州综保,杭州海关,重庆口岸 要求字段）
// attr006	报关批次（杭州口岸必填，yyyy-dd-mm）
// attr007	商品保险费用，无则填0（郑州综保,杭州海关,重庆口岸 要求字段）
// attr009	杭州海关版本代码：03
// attr010	电商企业十位编码，郑州综保字段

// UpdateOrder
// attr001
// attr002
// userId	商家或合作店铺id 丰网业务
// branchCode	丰网合作网点（丰网必填) 丰网业务
// branchAddressId	丰网合作地址id 丰网业务
// channelCode	渠道编码 丰网业务

type ExtraInfo struct {
	AttrName string `json:"attrName"` // F 扩展字段说明：attrName为字段定义， 具体如下表，value存在attrVal
	AttrVal  string `json:"attrVal"`  // F 扩展字段值
}

type SearchRoutesReq struct {
	Language        string   `json:"language"`        // 返回描述语语言 0：中文 1：英文 2：繁体
	TrackingType    string   `json:"trackingType"`    // 查询号类别: 1:根据顺丰运单号查询,trackingNumber将被当作顺丰运单号处理 2:根据客户订单号查询,trackingNumber将被当作客户订单号处理
	TrackingNumber  []string `json:"trackingNumber"`  // 查询号: trackingType=1,则此值为顺丰运单号 如果trackingType=2,则此值为客户订单号
	MethodType      string   `json:"methodType"`      // 路由查询类别: 1:标准路由查询 2:定制路由查询
	ReferenceNumber string   `json:"referenceNumber"` // 参考编码(目前针对亚马逊客户,由客户传)
	CheckPhoneNo    string   `json:"checkPhoneNo"`    // 电话号码验证
}

type SearchRoutesResp struct {
	RouteResps []*RoutesResp `json:"routeResps"`
}

type RoutesResp struct {
	MailNo string   `json:"mailNo"`
	Routes []*Route `json:"routes"`
}

type Route struct {
	AcceptAddress string `json:"acceptAddress"`
	AcceptTime    string `json:"acceptTime"`
	Remark        string `json:"remark"`
	OpCode        string `json:"opCode"`
}

type UpdateOrderReq struct {
	OrderId                 string           `json:"orderId"`                 // T 客户订单号
	DealType                int              `json:"dealType"`                // F 客户订单操作标识: 1:确认 (丰桥下订单接口默认自动确认，不需客户重复确认，该操作用在其它非自动确认的场景) 2:取消
	WaybillNoInfoList       []*WaybillNoInfo `json:"waybillNoInfoList"`       // F 顺丰运单号(如dealtype=1， 必填)
	CustomsBatchs           string           `json:"customsBatchs"`           // F 报关批次
	CollectEmpCode          string           `json:"collectEmpCode"`          // F 揽收员工号
	InProcessWaybillNo      string           `json:"inProcessWaybillNo"`      // F 头程运单号
	SourceZoneCode          string           `json:"sourceZoneCode"`          // F 原寄地网点代码
	DestZoneCode            string           `json:"destZoneCode"`            // F 目的地网点代码
	TotalWeight             float64          `json:"totalWeight"`             // F 订单货物总重量，包含子母 件，单位千克，精确到小数点 后3位，如果提供此值，必 须>0
	TotalVolume             float64          `json:"totalVolume"`             // F 订单货物总体积，单位立方厘 米，精确到小数点后3位，会 用于计抛（是否计抛具体商务 沟通中双方约定）
	ExpressTypeId           int              `json:"expressTypeId"`           // F 快件产品类别，支持附录《快 件产品类别表》的产品编码 值，仅可使用与顺丰销售约定 的快件产品
	ExtraInfoList           []*ExtraInfo     `json:"extraInfoList"`           // F 扩展属性
	TotalLength             float64          `json:"totalLength"`             // F 客户订单货物总长，单位厘米， 精确到小数点后3位，包含子 母件
	TotalWidth              float64          `json:"totalWidth"`              // F 客户订单货物总宽，单位厘米， 精确到小数点后3位，包含子 母件
	TotalHeight             float64          `json:"totalHeight"`             // F 客户订单货物总高，单位厘米， 精确到小数点后3位，包含子 母件
	ServiceList             []*Service       `json:"serviceList"`             // F 增值服务信息
	IsConfirmNew            int              `json:"isConfirmNew"`            // F 是否走新通用确认1：支持修改联系人 2：支持改其他客户订单默认0
	DestContactInfo         *ContactInfo     `json:"destContactInfo"`         // F 收件人信息
	IsDocall                int              `json:"isDocall"`                // F 是否通过手持终端通知顺丰收派员上门收件， 支持以下值：1：要求其它为不要求
	SpecialDeliveryTypeCode string           `json:"specialDeliveryTypeCode"` // F 特殊派送类型代码 身份验证 2. 极效前置单
	SpecialDeliveryValue    string           `json:"specialDeliveryValue"`    // F 特殊派件具体表述 证件类型:证件后8位 如：1:09296231（1表示身份证，暂不支持其他证件） 2>.极效前置单时:Y:若不支持则返回普通运单N:若不支持则返回错误码
	SendStartTm             string           `json:"sendStartTm"`             // F 预约时间(上门揽收时间)
	PickupAppointEndtime    string           `json:"pickupAppointEndtime"`    // F 上门揽收截止时间
}

type UpdateOrderResp struct {
	OrderId           string           `json:"orderId"`           // T 客户订单号
	WaybillNoInfoList []*WaybillNoInfo `json:"waybillNoInfoList"` // F 顺丰运单号
	ResStatus         int              `json:"resStatus"`         // T 备注 1:客户订单号与顺丰运单不匹配 2 :操作成功
	ExtraInfoList     []*ExtraInfo     `json:"extraInfoList"`     // F 扩展属性
}

type QuerySFWaybillReq struct {
	TrackingType    string `json:"trackingType"`    // T 1:表示按订单查询 2:表示按运单查询
	TrackingNum     string `json:"trackingNum"`     // T 订单号或运单号；
	Phone           string `json:"phone"`           // C 配置校验手机号时必传，支持最高6个收寄方电话后4位下发， 每个电话号码使用英文逗号分隔
	BizTemplateCode string `json:"bizTemplateCode"` // F 业务配置代码，针对客户业务需求配置的一套接口处理逻辑，一个接入编码可对应多个业务配置代码
}

type QuerySFWaybillResp struct {
	WaybillInfo    *WaybillInfo `json:"WaybillInfo"`    // 运单信息
	WaybillFeeList []WaybillFee `json:"waybillFeeList"` // 费用信息
}

type WaybillInfo struct {
	WaybillNo             string  `json:"waybillNo"`             // 运单号
	OrderId               string  `json:"orderId"`               // 客户订单号
	WaybillChilds         string  `json:"waybillChilds"`         // 子单号,英文逗号分隔
	CustomerAcctCode      string  `json:"customerAcctCode"`      // 月结账号
	MeterageWeightQty     float64 `json:"meterageWeightQty"`     // 计费重量
	RealWeightQty         float64 `json:"realWeightQty"`         //	实际重量
	ConsigneeEmpCode      string  `json:"consigneeEmpCode"`      // 收件员工号
	DeliverEmpCode        string  `json:"deliverEmpCode"`        // 派件员工号
	CargoTypeCode         string  `json:"cargoTypeCode"`         // 快件内容
	CargoTypeName         string  `json:"cargoTypeName"`         // 快件内容名称
	LimitTypeCode         string  `json:"limitTypeCode"`         // 时效类型
	LimitName             string  `json:"limitName"`             // 时效类型名称
	ExpressTypeCode       string  `json:"expressTypeCode"`       // 业务类型
	ExpressTypeName       string  `json:"expressTypeName"`       // 业务类型名称
	ConsValue             float64 `json:"consValue"`             // 声明价值
	ConsValueCurrencyCode string  `json:"consValueCurrencyCode"` // 声明价值币种
	ProductCode           string  `json:"productCode"`           // 产品代码
	ProductName           string  `json:"productName"`           // 产品名称
	JProvince             string  `json:"jProvince"`             // 寄件省份
	JCity                 string  `json:"jCity"`                 // 寄件城市
	ConsignorAddr         string  `json:"consignorAddr"`         // 寄件详细地址
	ConsignorContName     string  `json:"consignorContName"`     // 寄件联系人
	ConsignorPhone        string  `json:"consignorPhone"`        // 寄件电话号码
	ConsignorMobile       string  `json:"consignorMobile"`       // 寄件手机号
	DProvince             string  `json:"dProvince"`             // 收件省份
	DCity                 string  `json:"dCity"`                 // 收件城市
	AddresseeAddr         string  `json:"addresseeAddr"`         // 收件详细地址
	AddresseeContName     string  `json:"addresseeContName"`     // 收件联系人
	AddresseePhone        string  `json:"addresseePhone"`        // 收件电话号码
	AddresseeMobile       string  `json:"addresseeMobile"`       // 收件手机号
}

type WaybillFee struct {
	Type               string  `json:"type"`               // 费用类型: 1、运费；2、其它费用；3、保价；4、代收货款等 具体查询官方文档
	Name               string  `json:"name"`               // 费用名称
	Value              float64 `json:"value"`              // 费用金额
	PaymentTypeCode    string  `json:"paymentTypeCode"`    // 付款类型：1、寄付；2、到付；3、第三方付
	SettlementTypeCode string  `json:"settlementTypeCode"` // 结算类型：1为现结；2为月结
	ServiceProdCode    string  `json:"serviceProdCode"`    // 增值服务代码
	InsuredValue       string  `json:"insuredValue"`       // 保价金额
	CustomerAcctCode   string  `json:"customerAcctCode"`   // 月结账号
}

type SubMailNoReq struct {
	OrderId           string           `json:"orderId"`           // T 客户订单号
	ParcelQty         int64            `json:"parcelQty"`         // T 子单号数
	WaybillNoInfoList []*WaybillNoInfo `json:"waybillNoInfoList"` // F 顺丰运单号子单申请需要指定每个子包裹的长宽高提交重量时需要传入此参数
}

type SubMailNoResp struct {
	OrderId           string           `json:"orderId"`           // T 客户订单号
	ParcelQty         int64            `json:"parcelQty"`         // F 子单号数
	WaybillNoInfoList []*WaybillNoInfo `json:"waybillNoInfoList"` // T 顺丰运单号
}

type SearchOrderReq struct {
	OrderId    string `json:"orderId"`    // T 客户订单号
	SearchType string `json:"searchType"` // F 查询类型：1正向单 2退货单
	Language   string `json:"language"`   // F 语言
}

type SearchOrderResp struct {
	OrderId             string            `json:"orderId"`             // T 客户订单号
	OriginCode          string            `json:"originCode"`          // F 原寄地区域代码，可用于顺丰电子运单标签打印
	DestCode            string            `json:"destCode"`            // F 目的地区域代码，可用 于顺丰电子运单标签打印
	FilterResult        string            `json:"filterResult"`        // F 筛单结果： 1：人工确认 2：可收派 3：不可以收派
	ReturnExtraInfoList []*ExtraInfo      `json:"returnExtraInfoList"` // F 返回信息扩展属性
	WaybillNoInfoList   []*WaybillNoInfo  `json:"waybillNoInfoList"`   // F 顺丰运单号
	RouteLabelInfo      []*RouteLabelInfo `json:"routeLabelInfo"`      // F 路由标签数据
}

type DeliveryNoticeReq struct {
	WaybillNo string `json:"waybillNo"` // T 运单号
	DataType  string `json:"dataType"`  // T 71：派送通知 72：通知退回
	Language  string `json:"language"`  // F 语言：处理错误时返回的消息的语言；zh-CN，中文简体；zh-TW；zh-HK；zh-MO；US；
}

type ExchangeOrderReq struct {
	Language        string `json:"language"`        // T 返回描述语语言 0：中文 1：英文 2：繁体
	OrderId         string `json:"orderId"`         // T 客户订单号
	OrigOrderId     string `json:"origOrderId"`     //   原订单号
	OrigWaybillNo   string `json:"origWaybillNo"`   //   原运单号
	BizTemplateCode string `json:"bizTemplateCode"` // F 业务配置代码，业务配置代码指BSP针对客户业务需求配置的一套接口处理逻辑，一个接入编码可对应多个业务配置代码。
	ExchangeType    int    `json:"exchangeType"`    // T 换货类型 2:收件后换，3:退换一体
	IsCheck         string `json:"isCheck"`         //   是否验货, 1-验货，其他-不验货
	NewOrder        *Order `json:"newOrder"`        // T 新货订单
	OldOrder        *Order `json:"oldOrder"`        // T 旧货订
}

type ExchangeOrderResp struct {
	OrderId        string            `json:"orderId"`        // T 客户订单号
	NewWaybill     []*WaybillInfo    `json:"newWaybill"`     // F 新货运单号
	OldWaybill     []*WaybillInfo    `json:"oldWaybill"`     // F 旧货运单号
	NewOriginCode  string            `json:"newOriginCode"`  //   新货的原寄地代码
	NewDestCode    string            `json:"newDestCode"`    //   新货的目的地代码
	OldDestCode    string            `json:"oldDestCode"`    //   旧货的目的地代码
	FilterResult   int               `json:"filterResult"`   // F 筛单结果： 1：人工确认 2：可收派 3：不可以收派
	Remark         string            `json:"remark"`         // C 如果filter_result = 3时为必填，不可以收派的原因代码： 1：收方超范围 2：派方超范围 3-：其它原因 高峰管控提示信息 【数字】：【高峰管控提示信息】(如 4：温馨提示 ，1：春运延时)
	RouteLabelInfo []*RouteLabelInfo `json:"routeLabelInfo"` // C 下单路由标签新数据结构参考下单接口 （和标准下单返回一致）
}

type WantedInterceptReq struct {
	WaybillNo      string          `json:"waybillNo"`      // T	运单号
	ServiceCode    string          `json:"serviceCode"`    // T	1转寄2退回3优派4再派5改自取（改派-其他自取点取件）6改派送（上门派送）7更改派送时间8修改收件人信息9更改付款方式10修改代收货款12作废
	Role           string          `json:"role"`           // T	发起方1（寄方）2（收方）3（第三方）
	PayMode        string          `json:"payMode"`        // T	付款方式：1寄付 2到付 3转第三方月结 4寄付月结
	MonthlyCardNo  string          `json:"monthlyCardNo"`  // F	月结卡号如果是月结账户，必填字段。付款方式为月结时，必填字段。修改代收货款，必填字段
	ProductType    string          `json:"productType"`    // F	产品类型
	CodAmount      float64         `json:"codAmount"`      // F	代收货款金额
	DeliverDate    string          `json:"deliverDate"`    // F 派送日期，形如“2018-05-02”
	DeliverTimeMin string          `json:"deliverTimeMin"` // F 派送最早时间，形如“09:00”
	DeliverTimeMax string          `json:"deliverTimeMax"` // F 派送最晚时间，形如“12:00”
	SelfPickPoint  string          `json:"selfPickPoint"`  // F 自取点 1-合作点 2-快递柜 3-网点 4中转场
	NewDestAddress *NewDestAddress `json:"newDestAddress"` // F 新目的地改联系人业务请携带姓名及手机号转寄、退回业务必须携带，撤销转寄、退回不携带改自取必须携带
}

type NewDestAddress struct {
	Country      string `json:"country"`      // F 国家，默认中国
	CountryCode  string `json:"countryCode"`  // F 国家编码，默认CN
	Province     string `json:"province"`     // T 省，如广东省
	City         string `json:"city"`         // T	市，如深圳市
	County       string `json:"county"`       // T 区，如南山区
	Address      string `json:"address"`      // T 详细地址，不包含省市区，如粤海街道高新区1C栋二楼
	Company      string `json:"company"`      // F 公司
	Contact      string `json:"contact"`      // F 联系人姓名，如黄飞鸿
	Phone        string `json:"phone"`        // F 电话，如18012345678
	AreaCode     string `json:"areaCode"`     // F 网点代码，如755WQ
	LocationCode string `json:"locationCode"` // F 城市代码
}

type ReverseOrder struct {
	Language        string         `json:"language"`        // F 响应报文的语言，缺省值为zh-CN，目前支持以下值zh-CN表示中文简体，zh-TW或zh-HK或zh-MO表示中文繁体，en表示英文
	OrderId         string         `json:"orderId"`         // T 客户订单号
	CargoDetails    []*CargoDetail `json:"cargoDetails"`    // T 拖寄物信息
	OrigOrderId     string         `json:"origOrderId"`     // F 原订单号
	OldClientCode   []interface{}  `json:"oldClientCode"`   // F 原顾客编码
	ServiceList     []*Service     `json:"serviceList"`     // F 增值服务信息
	ContactInfoList []*ContactInfo `json:"contactInfoList"` //	 收寄双方信息
	MonthlyCard     string         `json:"monthlyCard"`     // C 顺丰月结卡号
	PayMethod       int            `json:"payMethod"`       // F 付款方式，支持以下值： 1:寄方付 2:收方付 3:第三方付
	ExpressTypeId   int            `json:"expressTypeId"`   // F 快件产品类别，支持附录《快件产品类别表》的产品编码值，仅可使用与顺丰销售约定的快件产品。
	SendStartTm     string         `json:"sendStartTm"`     // F 要求上门取件开始时间，格式：YYYY-MM-DD HH24:MM:SS，示例：2012-7-30 09:30:00。
	IsCheck         string         `json:"isCheck"`         // F 是否验货, 1-验货，其他-不验货
	ShopName        string         `json:"shopName"`        // F 商家店铺名称
	OrigWaybillNo   string         `json:"origWaybillNo"`   // C 原运单号
	OrderType       string         `json:"orderType"`       // F 订单类型：退货订单，维修订单
	BizTemplateCode string         `json:"bizTemplateCode"` // F 业务配置代码，业务配置代码指BSP针对客户业务需求配置的一套接口处理逻辑，一个接入编码可对应多个业务配置代码。
	RefundAmount    float64        `json:"refundAmount"`    // F 退款金额
	Volume          float64        `json:"volume"`          // F 托寄物体积
	Remark          string         `json:"remark"`          // F 备注
}
