package sf

import (
	"fmt"
	"log"
	"testing"
)

func TestClient_FilterOrderBSP(t *testing.T) {
	data, err := c.FilterOrderBSP(&FilterOrderBSPReq{
		ContactInfos: []*FilterAddrInfo{
			{
				Address:     "创业总部基地B07二楼",
				City:        "天津市",
				ContactType: 2,
				Country:     "中国",
				County:      "武清区",
				Province:    "天津市",
				Tel:         "19851401196",
			},
			{
				Address:     "测试订单，请不要接单",
				City:        "成都市",
				ContactType: 1,
				Country:     "中国",
				County:      "锦江区",
				Province:    "四川省",
			},
		},
		FilterType: 1,
		OrderId:    "QIAO-20200728-006",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_ValidateWayBillNo(t *testing.T) {
	data, err := c.ValidateWayBillNo(&ValidateWayBillNoReq{
		WaybillNo: "SF1040275268927",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_QueryGISDepartment(t *testing.T) {
	data, err := c.QueryGISDepartment(&QueryGISDepartmentReq{
		Address:  "广东省深圳市宝安区劳动路",
		X:        "100",
		Y:        "60",
		Opt:      "dq0",
		DeptType: "1|2|5",
		ServType: "1|2|5",
		Distance: 1000,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}
