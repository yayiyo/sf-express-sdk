package sf

import (
	"encoding/json"
	"fmt"
)

// CreateOrder 功能描述
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=819535&interName=%E4%B8%8B%E8%AE%A2%E5%8D%95%E6%8E%A5%E5%8F%A3-EXP_RECE_CREATE_ORDER
// 下订单接口根据客户需要,可提供以下四个功能:
//  ● 客户系统向顺丰下发订单
//  ● 为订单分配运单号
//  ● 筛单
//  ● 路由轨迹(可选)
func (c *Client) CreateOrder(req *CreateOrderReq) (*Order, error) {
	data, err := c.Do(req, serviceCodeCreateOrder)
	if err != nil {
		return nil, err
	}

	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	o := &Order{}
	err = json.Unmarshal(d, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// SearchRoutes 路由查询
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=221529&interName=%E8%B7%AF%E7%94%B1%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3-EXP_RECE_SEARCH_ROUTES
/**
 客户可通过此接口查询顺丰运单路由，顺丰会在响应Json报文返回当时点要求的全部路由节点信息。
 此路由查询接口支持三类查询方式:
		1)根据顺丰运单号查询：查询请求中提供接入编码与运单号，
	验证接入编码与所有请求运单号的归属关系，系统只返回具有正确归属关系的运单路由信息。
		2)根据客户订单号查询：查询请求中提供接入编码与订单号，
	验证接入编码与所有请求订单号的归属关系，对于归属关系正确的订单号，找到对应的运单号，
	然后返回订单对应运单号的路由信息。适用于通过下单的客户订单。
		3)根据运单号+运单对应的收寄人任一方电话号码后4位(参数checkPhoneNo中传入)查询,
	系统校验信息匹配将返回对应运单路由信息。
*/
func (c *Client) SearchRoutes(s *SearchRoutesReq) ([]*RoutesResp, error) {
	data, err := c.Do(s, serviceCodeSearchRoutes)
	if err != nil {
		return nil, err
	}

	resp := data["routeResps"]
	fmt.Println(resp)
	d, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	routes := make([]*RoutesResp, 0)
	err = json.Unmarshal(d, &routes)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

// UpdateOrder 订单确认/取消接口
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=470176&interName=%E8%AE%A2%E5%8D%95%E7%A1%AE%E8%AE%A4%2F%E5%8F%96%E6%B6%88%E6%8E%A5%E5%8F%A3-EXP_RECE_UPDATE_ORDER
/**
接口用于以下场景:
	(1)客户在确定将货物交付给顺丰托运后，将运单上的一些重要信息，如快件重量通过此接口发送给顺丰。
	(2)客户在发货前取消订单。
	注意：订单取消之后，订单号也是不能重复利用的。
*/
func (c *Client) UpdateOrder(u *UpdateOrderReq) (*UpdateOrderResp, error) {
	data, err := c.Do(u, serviceCodeUpdateOrder)
	if err != nil {
		return nil, err
	}

	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	o := &UpdateOrderResp{}
	err = json.Unmarshal(d, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// QuerySFWaybill 清单运费查询
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=925884&interName=%E6%B8%85%E5%8D%95%E8%BF%90%E8%B4%B9%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3-EXP_RECE_QUERY_SFWAYBILL
/**
 * 此功能主要是根据订单号或者运单号查询清单运费信息。
 */
func (c *Client) QuerySFWaybill(q *QuerySFWaybillReq) (*QuerySFWaybillResp, error) {
	data, err := c.Do(q, serviceCodeQuerySFWayBill)
	if err != nil {
		return nil, err
	}

	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res := &QuerySFWaybillResp{}
	err = json.Unmarshal(d, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSubMailNo 子单号申请接口
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=816719&interName=%E5%AD%90%E5%8D%95%E5%8F%B7%E7%94%B3%E8%AF%B7%E6%8E%A5%E5%8F%A3-EXP_RECE_GET_SUB_MAILNO
/**
 * 客户在下单成功后 ，因业务场景需要可以调用此接口获取更多的子单号数。但不能超过配置的最大数/1200个
 */
func (c *Client) GetSubMailNo(s *SubMailNoReq) (*SubMailNoResp, error) {
	data, err := c.Do(s, serviceCodeGetSubMailNo)
	if err != nil {
		return nil, err
	}

	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res := &SubMailNoResp{}
	err = json.Unmarshal(d, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SearchOrderResp 订单结果查询接口
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=299038&interName=%E8%AE%A2%E5%8D%95%E7%BB%93%E6%9E%9C%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3-EXP_RECE_SEARCH_ORDER_RESP
/**
 * 因Internet环境下，网络不是绝对可靠，用户系统下订单到顺丰后，
 * 不一定可以收到顺丰系统返回的数据，此接口用于在未收到返回数据时，
 * 查询订单创建接口客户订单当前的处理情况。
 */
func (c *Client) SearchOrderResp(s *SearchOrderReq) (*SearchOrderResp, error) {
	data, err := c.Do(s, serviceCodeSearchOrderResp)
	if err != nil {
		return nil, err
	}

	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res := &SearchOrderResp{}
	err = json.Unmarshal(d, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeliveryNotice 派送通知接口
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=184259&interName=%E6%B4%BE%E9%80%81%E9%80%9A%E7%9F%A5%E6%8E%A5%E5%8F%A3-EXP_RECE_DELIVERY_NOTICE
/**
 * 派件通知接口，双11等大促活动前，电商平台有预付款功能。
 * 消费者只需要提前缴纳就可以在双十一当前抵扣，获取更大的优惠。
 * 预付完成后商家会提前将产品交给顺丰，保存到离收货地址最近的仓。
 * 大促活动当天消费者全款支付完毕后，商家会通知顺丰送达客户
 */
func (c *Client) DeliveryNotice(d *DeliveryNoticeReq) error {
	_, err := c.Do(d, serviceCodeDeliveryNotice)
	if err != nil {
		return err
	}
	return nil
}

// CreateExchangeOrder 换货下单接口
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=428991&interName=%E6%8D%A2%E8%B4%A7%E4%B8%8B%E5%8D%95%E6%8E%A5%E5%8F%A3-EXP_RECE_CREATE_EXCHANGE_ORDER
/**
 *	客户系统向顺丰下发订单
 *	为订单分配运单号
 *	筛单
 */
func (c *Client) CreateExchangeOrder(ceo *ExchangeOrderReq) (*ExchangeOrderResp, error) {
	data, err := c.Do(ceo, serviceCodeCreateExchangeOrder)
	if err != nil {
		return nil, err
	}

	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res := &ExchangeOrderResp{}
	err = json.Unmarshal(d, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
