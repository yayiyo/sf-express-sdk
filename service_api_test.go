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

func TestClient_QueryDeliverTM(t *testing.T) {
	data, err := c.QueryDeliverTm(&QueryDeliverTMReq{
		BusinessType:  "2",
		ConsignedTime: "2022-05-25 15:03:22",
		DestAddress: &Address{
			Address:  "北京街道西湖路38号首层102号东南铺江博士",
			City:     "广州市",
			District: "越秀区",
			Province: "广东省",
		},
		SearchPrice: "1",
		SrcAddress: &Address{
			Address:  "琶洲街道琶洲蟠龙新街2号保利广场购物中心3层3036号江博士专卖铺",
			City:     "广州市",
			District: "海珠区",
			Province: "广东省",
		},
		Weight: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_SearchPromiseTm(t *testing.T) {
	data, err := c.SearchPromiseTm(&SearchPromiseTmReq{
		SearchNo:  "SF7444439535578",
		CheckType: 1,
		CheckNos:  []string{"4006789888"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_CheckPickUpTime(t *testing.T) {
	data, err := c.CheckPickUpTime(&CheckPickUpTimeReq{
		Address:     "深圳市南山区软件产业基地1栋A座",
		CityCode:    "755",
		AddressType: "1",
		SendTime:    "2021-03-21 12:00:00",
		SysCode:     "qiao",
		Version:     "V1.1",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}
