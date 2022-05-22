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
		TrackingNumber: []string{"444003077898", "441003077850"},
		MethodType:     "1",
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
				Amount: 308.0,
				Count:  1.0,
				Name:   "君宝牌地毯",
				Unit:   "件",
				Volume: 0,
				Weight: 6.6,
			},
		},
		ContactInfoList: []*ContactInfo{
			{
				Address:     "十堰市丹江口市公园路155号",
				City:        "十堰市",
				Company:     "清雅轩保健品专营店",
				Contact:     "张三丰",
				ContactType: 1,
				County:      "武当山风景区",
				Mobile:      "17006805888",
				Province:    "湖北省",
			}, {
				Address:     "湖北省襄阳市襄城区环城东路122号",
				City:        "襄阳市",
				Contact:     "郭襄阳",
				County:      "襄城区",
				ContactType: 2,
				Mobile:      "18963828829",
				Province:    "湖北省",
			},
		},
		CustomsInfo:     &CustomsInfo{},
		ExpressTypeId:   1,
		ExtraInfoList:   []*ExtraInfo{},
		IsOneselfPickup: 0,
		Language:        "zh-CN",
		MonthlyCard:     "7551234567",
		OrderId:         "QIAO-20200728-004",
		ParcelQty:       1,
		PayMethod:       1,
		TotalWeight:     6,
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
		OrderId:    "QIAO-20200728-004",
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
