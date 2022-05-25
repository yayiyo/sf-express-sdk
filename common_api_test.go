package sf

import (
	"fmt"
	"log"
	"testing"
)

var c = NewClient("https://sfapi-sbox.sf-express.com/std/service", "LBLERiwZfH", "oNv0SZTAyeUwRCTUNGvswI1D5oQPHU90")

func TestSearchRoutes(t *testing.T) {
	data, err := c.SearchRoutes(&SearchRoutesReq{
		Language:       "0",
		TrackingType:   "1",
		TrackingNumber: []string{"SF1424638720536"},
		MethodType:     "1",
		CheckPhoneNo:   "8648",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_CreateOrder(t *testing.T) {
	data, err := c.CreateOrder(&CreateOrderReq{
		CargoDetails: []*CargoDetail{
			{
				Count:      2.365,
				Unit:       "个",
				Weight:     6.1,
				Amount:     100.5111,
				Currency:   "HKD",
				Name:       "护肤品1",
				SourceArea: "CHN",
			},
		},
		ContactInfoList: []*ContactInfo{
			{
				Address:     "广东省深圳市南山区软件产业基地11栋",
				Contact:     "小曾",
				ContactType: 1,
				Country:     "CN",
				PostCode:    "580058",
				Tel:         "17600668648",
			},
			{
				Address:     "广东省广州市白云区湖北大厦",
				Company:     "顺丰速运",
				Contact:     "小邱",
				ContactType: 2,
				Country:     "CN",
				PostCode:    "580058",
				Tel:         "18688806057",
			},
		},
		Language:      "zh_CN",
		OrderId:       "QIAO-20200728-012",
		ExpressTypeId: 2,
		ParcelQty:     1,
		PayMethod:     1,
		TotalWeight:   6,
		IsDocall:      1,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_UpdateOrder(t *testing.T) {
	data, err := c.UpdateOrder(&UpdateOrderReq{
		DealType: 2,
		OrderId:  "QIAO-20200728-003",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_QuerySFWaybill(t *testing.T) {
	data, err := c.QuerySFWaybill(&QuerySFWaybillReq{
		TrackingType: "1",
		TrackingNum:  "QIAO-20200728-003",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_GetSubMailNo(t *testing.T) {
	data, err := c.GetSubMailNo(&SubMailNoReq{
		ParcelQty: 3,
		OrderId:   "QIAO-20200728-004",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_SearchOrderResp(t *testing.T) {
	data, err := c.SearchOrderResp(&SearchOrderReq{
		SearchType: "1",
		OrderId:    "QIAO-20200728-006",
		Language:   "zh-cn",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_DeliveryNotice(t *testing.T) {
	err := c.DeliveryNotice(&DeliveryNoticeReq{
		WaybillNo: "QIAO-20200728-004",
		DataType:  "71",
		Language:  "zh-cn",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_CreateExchangeOrder(t *testing.T) {
	data, err := c.CreateExchangeOrder(&ExchangeOrderReq{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestClient_CreateReverseOrder(t *testing.T) {
	data, err := c.CreateReverseOrder(&ReverseOrder{
		Language: "zh_CN",
		OrderId:  "F2_20200604180946",
		CargoDetails: []*CargoDetail{
			{
				Amount:   100.5111,
				Count:    2.365,
				Currency: "HKD",
				CargoSku: "AAAA004",
				Name:     "护肤品1",
				Unit:     "个",
				Weight:   6.1,
			},
		},
		ServiceList: []*Service{
			{
				Name:  "INSURE",
				Value: "3000",
			},
		},
		ContactInfoList: []*ContactInfo{
			{
				Address:     "软件产业基地11栋",
				City:        "深圳市",
				Contact:     "小曾",
				ContactType: 1,
				Country:     "CN",
				County:      "南山区",
				Mobile:      "13480155048",
				PostCode:    "580058",
				Province:    "广东省",
				Tel:         "4006789888",
			}, {
				Address:     "广东省广州市白云区湖北大厦",
				City:        "",
				Company:     "顺丰速运",
				Contact:     "小邱",
				ContactType: 2,
				Country:     "CN",
				County:      "",
				Mobile:      "13925211148",
				PostCode:    "580058",
				Province:    "",
				Tel:         "18688806057",
			},
		},
		MonthlyCard:   "",
		PayMethod:     1,
		ExpressTypeId: 1,
		Volume:        8.0,
		//SendStartTm:     time.Now(),
		RefundAmount:    8.0,
		IsCheck:         "1",
		BizTemplateCode: "2020312_order",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}
