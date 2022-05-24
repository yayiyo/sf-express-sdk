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
