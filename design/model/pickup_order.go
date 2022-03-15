/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    pickuporder
	@Date    2022/2/10 11:05 上午:00
	@Desc
*/

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

// PickupOrder 揽件信息请求入参.
var PickupOrder = Type("PickupOrder", func() {
	Description("PickupOrder describes the pickup info")
	Field(1, "requested_pickup_at", String, "shipping at", func() {
		Example("2021-01-12 09:34:09")
		MaxLength(30)
	})
	Field(2, "items", ArrayOf(InboundOrderItems), "inbound order items", func() {
		MinLength(1)
	})
	Field(4, "address", ShippingAddress, "address")
	Field(6, "type", Int, "delivery mode(1 direct，2 warehouse)", func() {
		Enum(1, 2)
		Example(1)
	})
	Field(7, "customer_code", String, "customer code", func() {
		Example("1013")
		MaxLength(50)
	})

	Required("requested_pickup_at", "items", "address", "type", "customer_code")
})

// PickupOrderRsp 揽件信息返回出参.
var PickupOrderRsp = Type("PickupOrderRsp", func() {
	Description("OutboundOrderRsp describes the outbound order return value")
	Field(1, "error_msg", String, "error message", func() {
		Example("xxx")
	})

	Required("error_msg")
})
