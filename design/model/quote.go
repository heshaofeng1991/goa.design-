/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    quote
	@Date    2022/2/10 10:50 上午:00
	@Desc
*/

// Package model API in/out parameter define of shipping-estimate
package model

import (
	. "goa.design/goa/v3/dsl"
)

var UpdateChannelCostStatusReq = Type("UpdateChannelCostStatusReq", func() {
	Field(1, "ids", ArrayOf(Int64), "渠道ID", func() {
		Example([]int64{1, 2})
	})
	Field(2, "country_codes", ArrayOf(String), "排除国际二字码", func() {
		Example([]string{"US", "UK"})
	})
	Field(3, "status", Boolean, "状态（0 不启用 1 启用）", func() {
		Enum(false, true)
		Example(true)
	})
	Extend(AuthToken)
	Required("ids", "country_codes", "status")
})

var UpdateCustomerConfigData = Type("UpdateCustomerConfigData", func() {
	Field(1, "status", Int32, "状态（0 更新成功 1 更新失败）", func() {
		Example(1)
	})
	Required("status")
})

var UpdateChannelCostStatusRsp = Type("UpdateChannelCostStatusRsp", func() {
	Field(1, "data", UpdateCustomerConfigData, "data")
	Extend(BaseResponse)
})
