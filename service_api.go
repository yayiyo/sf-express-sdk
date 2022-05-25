package sf

import (
	"encoding/json"
	"fmt"
)

// 服务查询类API

// FilterOrderBSP 订单筛选
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=474424&interName=%E8%AE%A2%E5%8D%95%E7%AD%9B%E9%80%89%E6%8E%A5%E5%8F%A3-EXP_RECE_FILTER_ORDER_BSP
/**
 * 客户系统通过此接口向顺丰系统发送主动的筛单请求,用于判断客户的收、派地址是否属于顺丰的收派范围
 */
func (c *Client) FilterOrderBSP(f *FilterOrderBSPReq) (*FilterOrderBSPResp, error) {
	data, err := c.Do(f, serviceCodeFilterOrderBSP)
	if err != nil {
		return nil, err
	}

	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res := &FilterOrderBSPResp{}
	err = json.Unmarshal(d, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ValidateWayBillNo 运单号合法性校验
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=430610&interName=%E8%BF%90%E5%8D%95%E5%8F%B7%E5%90%88%E6%B3%95%E6%80%A7%E6%A0%A1%E9%AA%8C%E6%8E%A5%E5%8F%A3-EXP_RECE_VALIDATE_WAYBILLNO
/**
 * 提供顺丰运单号合法性校验功能
 */
func (c *Client) ValidateWayBillNo(v *ValidateWayBillNoReq) (bool, error) {
	data, err := c.Do(v, serviceCodeValidateWayBillNo)
	if err != nil {
		return false, err
	}

	if b, ok := data.(bool); ok {
		return b, nil
	}
	return false, fmt.Errorf("the response %v is not bool type", data)
}

// QueryGISDepartment 服务网点查询
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=312430&interName=%E9%A1%BA%E4%B8%B0%E6%9C%8D%E5%8A%A1%E7%BD%91%E7%82%B9%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3-EXP_RECE_QUERY_GIS_DEPARTMENT
/**
 * 客户可通过经纬度，查询附近的10个网点信息
 */
func (c *Client) QueryGISDepartment(q *QueryGISDepartmentReq) (*QueryGISDepartmentResp, error) {
	data, err := c.Do(q, serviceCodeQueryGISDepartment)
	if err != nil {
		return nil, err
	}
	res := &QueryGISDepartmentResp{}
	err = json.Unmarshal([]byte(data.(string)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// QueryDeliverTm 时效标准及价格查询
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=279416&interName=%E6%97%B6%E6%95%88%E6%A0%87%E5%87%86%E5%8F%8A%E4%BB%B7%E6%A0%BC%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3-EXP_RECE_QUERY_DELIVERTM
/**
 * 客户可通过接口查询从特定原寄地寄特定目的地的时效和价格
 */
func (c *Client) QueryDeliverTm(q *QueryDeliverTMReq) (*QueryDeliverTMResp, error) {
	data, err := c.Do(q, serviceCodeQueryDeliverTm)
	if err != nil {
		return nil, err
	}
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res := &QueryDeliverTMResp{}
	err = json.Unmarshal(d, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SearchPromiseTm 预计派送时间查询
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=226162&interName=%E9%A2%84%E8%AE%A1%E6%B4%BE%E9%80%81%E6%97%B6%E9%97%B4%E6%8E%A5%E5%8F%A3-EXP_RECE_SEARCH_PROMITM
/**
 * 此功能主要是根据运单号查询预计派件时间
 */
func (c *Client) SearchPromiseTm(q *SearchPromiseTmReq) (*SearchPromiseTmResp, error) {
	data, err := c.Do(q, serviceCodeSearchPromiseTm)
	if err != nil {
		return nil, err
	}
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res := &SearchPromiseTmResp{}
	err = json.Unmarshal(d, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CheckPickUpTime 揽件服务时间校验
// 文档地址 https://open.sf-express.com/Api/ApiDetails?level3=888570&interName=%E6%8F%BD%E4%BB%B6%E6%9C%8D%E5%8A%A1%E6%97%B6%E9%97%B4%E6%A0%A1%E9%AA%8C%E6%8E%A5%E5%8F%A3-EXP_EXCE_CHECK_PICKUP_TIME
/**
 * 客户系统通过该接口检查是否为揽件服务时间范围
 */
func (c *Client) CheckPickUpTime(cp *CheckPickUpTimeReq) (*CheckPickUpTimeResp, error) {
	data, err := c.Do(cp, serviceCodeCheckPickUpTime)
	if err != nil {
		return nil, err
	}

	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res := &CheckPickUpTimeResp{}
	var rd interface{}
	// 返回值是状态（bool）
	if cp.Version == "V1.0" {
		rd = &res.Status
	} else {
		rd = res
	}
	err = json.Unmarshal(d, rd)
	if err != nil {
		return nil, err
	}
	return res, nil
}
