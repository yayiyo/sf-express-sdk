package sf

// 必填 条件：C 是：T 否：F

type FilterOrderBSPReq struct {
	FilterType   int               `json:"filterType"`   // T 筛单类别: 1:自动筛单(系统根据地址库进行判断,并返回结果,系统无法检索到可派送的将返回不可派送) 2:可人工筛单(系统首先根据地址库判断,如果无法自动判断是否收派,系统将生成需要人工判断的任务,后续由人工处理,处理结束后,顺丰可主动推送给客户系统)
	OrderId      string            `json:"orderId"`      // F 客户订单号
	MonthlyCard  string            `json:"monthlyCard"`  // F 月结卡号
	ContactInfos []*FilterAddrInfo `json:"contactInfos"` // F 地址信息
}

type FilterAddrInfo struct {
	ContactType int    `json:"contactType"` // F 地址类型 1:寄件方 2:到件方
	Tel         string `json:"tel"`         // F 联系电话
	Country     string `json:"country"`     // F 国家或地区代码 2位
	Province    string `json:"province"`    // F 省级行政区名称
	City        string `json:"city"`        // F 地级行政区名称
	County      string `json:"county"`      // F 县/区级行政区名称
	Address     string `json:"address"`     // F 详细地址
	PostCode    string `json:"postCode"`    // C 邮编，跨境件必填
}

type FilterOrderBSPResp struct {
	OrderId      string `json:"orderId"`      // T 客户订单号
	FilterResult int    `json:"filterResult"` // T 筛单结果： 1：人工确认 2：可收派 3：不可以收派 4 : 地址无法确认 当filter_type=1时，不存在orderId
	OriginCode   string `json:"originCode"`   // C 原寄地区域代码，如果可收 派，此项不能为空
	DestCode     string `json:"destCode"`     // C 目的地区域代码，如果可收 派，此项不能为空
	Remark       string `json:"remark"`       // C 如果filter_result=3时为必 填，不可以收派的原因代码： 1：收方超范围 2：派方超范围 3：其它原因
}

type ValidateWayBillNoReq struct {
	WaybillNo string `json:"waybillNo"` // T 运单号 （支持12位或者15位运单号）
}

type QueryGISDepartmentReq struct {
	Address  string `json:"address"`  // F 详细地址
	X        string `json:"x"`        // F 经度
	Y        string `json:"y"`        // F 纬度
	Opt      string `json:"opt"`      // F dq0：显示所有接口返回原始信息；dq1：显示格式化后信息（默认）
	DeptType string `json:"deptType"` // F 网点类型： 1自营服务点, 2合作商家店, 3嘿客店/顺丰优选 4顺丰站 5丰巢柜（允许多选，多个值之间用竖线隔开，如：1|2|5）
	ServType string `json:"servType"` // F 服务类型：1、自寄服务2、自取服务 3、寄、取件服务4、个人地址服务 5、便民服务6自寄自取优惠服务（允许多选，多个值之间用竖线隔开，如：1|2|5）
	Count    int    `json:"count"`    // F 数量
	Distance int    `json:"distance"` // F 查询范围，默认为1000米，单位：米
	City     string `json:"city"`     // F 城市：调用香港，需输入city = 香港
}

type QueryGISDepartmentResp struct {
	Status int       `json:"status"` // 状态 0成功 其他：失败
	Count  int       `json:"count"`  // 	返回条数据
	Src    string    `json:"src"`    // 来源
	Result []*Result `json:"result"` // 网点信息结果
}

type Result struct {
	Address    string `json:"address"`
	Distance   int    `json:"distance"`
	ServerType string `json:"servertype"`
	Latitude   int    `json:"latitude"`
	Name       string `json:"name"`
	Servtime   string `json:"servtime"`
	Telephone  string `json:"telephone"`
	Id         string `json:"id"`
	Longitude  int    `json:"longitude"`
}

type QueryDeliverTMReq struct {
	BusinessType  string   `json:"businessType"`  // T 快件产品：可以为空，为空时查询顺丰支持的所有业务类型。不为空时以数字代码业务类型，例如：1：表示“”2：表示“顺丰特惠”5：表示“顺丰次晨”6：表示“即日件
	Weight        float64  `json:"weight"`        // F 货物总重量，包含子母件，单位千克，精确到小数点后2位，如果提供此值，必须>0。
	Volume        float64  `json:"volume"`        // F 货物的体积（长、宽、高分别以厘米为单位计算体积），精确到小数点后2位。
	ConsignedTime string   `json:"consignedTime"` // F 寄件时间，格式为YYYY-MM-DD HH24:MM:SS，示例2013-12-27 17:54:20。
	SearchPrice   string   `json:"searchPrice"`   // F 1：表示查询含价格的接口0：表示查询不含价格的接口 备注：限制只能为0, 1或者不传searchPrice，不可以为空, null
	DestAddress   *Address `json:"destAddress"`   // T 目的地信息
	SrcAddress    *Address `json:"srcAddress"`    // T 原寄地信息
}

type Address struct {
	Province string `json:"province"` // C 目的地所在省份，字段填写要求：必须是标准的省名称称谓 如：广东省; 如果是直辖市，请直接传北京、上海等，如果字段code为空时为必填。
	City     string `json:"city"`     // C 目的地所在城市，必须是标准的城市称谓 如：深圳市，如果字段code为空时为必填。
	District string `json:"district"` // F 目的地所在县/区，必须是标准的县/区称谓，示例：“福田区”。
	Address  string `json:"address"`  // F 目的地详细地址，此详细地址需包含省市信息，以提高地址识别的成功率，示例：“广东省深圳市福田区新洲十一街万基商务大厦10楼”。
	Code     string `json:"code"`     // C 目的地区域代码，如果填写了此项，则查询时忽略省市区具体地址，如果不填此项，则综合省市区具体地址识别区域代码，字段province或city为空时为必填，示例：020、755。
}

type QueryDeliverTMResp struct {
	DeliverTmDto []*DeliverTM `json:"deliverTmDto"`
}

type DeliverTM struct {
	BusinessType     string  `json:"businessType"`     // 快件产品
	BusinessTypeDesc string  `json:"businessTypeDesc"` // 快件产品描述
	DeliverTime      string  `json:"deliverTime"`      // 承诺时间
	Fee              float64 `json:"fee"`              // 价格
	SearchPrice      string  `json:"searchPrice"`      // 是否查询价格（1, 返回价格；0，不返回价格）
	CloseTime        string  `json:"closeTime"`        // 截单时间
}

type SearchPromiseTmReq struct {
	SearchNo  string   `json:"searchNo"`  // T 顺丰运单号
	CheckType int      `json:"checkType"` // F 校验类型 1，电话号码校验 2，月结卡号校验
	CheckNos  []string `json:"checkNos"`  // T 校验值 当校验类型为1时传电话号码 当校验类型为2时传月结卡号
}

type SearchPromiseTmResp struct {
	SearchNo  string `json:"searchNo"`  // T 顺丰运单号
	PromiseTm string `json:"promiseTm"` // T 预计派件时间格式:YYYY-MM-DD HH24:mm:SS
}

type CheckPickUpTimeReq struct {
	Address     string `json:"address"`     // T 地址
	CityCode    string `json:"cityCode"`    // T 城市代码
	AddressType string `json:"addressType"` // T 地址类型 （1：寄件地址，2：收件地址）
	SendTime    string `json:"sendTime"`    // T 预约上门时间(yyyy-MM-dd HH:mm:ss)
	SysCode     string `json:"sysCode"`     // T 来源系统
	Version     string `json:"version"`     // F version:不传 只返回是否在服务时间范围 1.1 返回是否在服务时间范围, 并返回服务时间段
	Province    string `json:"province"`    // F 省，如果没有cityCode，则通过省市区转换
	City        string `json:"city"`        // F 市
	County      string `json:"county"`      // F 区
}

type CheckPickUpTimeResp struct {
	Status  bool        `json:"status"`
	EndTm   string      `json:"endTm"`
	StartTm string      `json:"startTm"`
	System  interface{} `json:"system"`
}
