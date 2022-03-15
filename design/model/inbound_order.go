/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    inboundorder
	@Date    2022/2/10 11:04 上午:00
	@Desc
*/

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

// InboundOrder 入库预报请求入参.
var InboundOrder = Type("InboundOrder", func() {
	Description("InboundOrder describes the inbound info")
	Field(1, "customer_order_id", String, "Customer Order ID", func() {
		Example("xxx1234")
		MinLength(1)
		MaxLength(50)
	})
	Field(2, "warehouse_id", Int64, "warehouse id", func() {
		Example(1)
		Minimum(1)
	})
	Field(3, "customer_code", String, "customer code", func() {
		Example("1013")
		MinLength(1)
		MaxLength(50)
	})
	Field(4, "tracking_number", String, "tracking number", func() {
		Example("YT000001")
		MaxLength(50)
	})
	Field(5, "requested_pickup_at", String, "requested pickup at", func() {
		Example("2021-01-12 09:34:09")
		MaxLength(30)
	})
	Field(6, "estimated_arrival_at", String, "estimated arrival at", func() {
		Example("2021-01-12 09:34:09")
		MaxLength(30)
	})
	Field(7, "items", ArrayOf(InboundOrderItems), "inbound order items")
	Field(8, "type", Int, "delivery mode(1 direct，2 warehouse)", func() {
		Enum(1, 2)
		Example(1)
	})
	Field(10, "address", ShippingAddress, "address")
	Field(11, "is_pickup", Boolean, "is pickup", func() {
		Example(true)
	})
	Field(12, "description", String, "description", func() {
		Example("description")
	})
	Field(14, "id", Int32, "inbound order id", func() {
		Example(1)
	})

	Required("customer_order_id", "warehouse_id", "customer_code", "tracking_number",
		"requested_pickup_at", "estimated_arrival_at", "items", "type", "address", "is_pickup", "description")
})

// InboundOrderItems 入库订单items请求消息体.
var InboundOrderItems = Type("Item", func() {
	Description("InboundOrderItems describes the inbound order items")
	Field(1, "product_name", String, "product name", func() {
		Example("NSS Mate 40E")
		MinLength(1)
		MaxLength(100)
	})
	Field(2, "product_sku", String, "product sku", func() {
		Example("YCrankshaft")
		MinLength(1)
		MaxLength(50)
	})
	Field(3, "barcode", String, "barcode", func() {
		Example("YCrankshaft")
		MaxLength(50)
	})
	Field(4, "qty", Int, "quality", func() {
		Example(3)
	})
	Required("product_name", "product_sku", "barcode", "qty")
})

// InboundOrderRsp 入库订单返回出参.
var InboundOrderRsp = Type("InboundOrderRsp", func() {
	Description("InboundOrderRsp describes the inbound order return value")
	Field(1, "inbound_order_id", Int64, "inbound order id", func() {
		Example(1)
	})
	Field(2, "label_url", String, "label url", func() {
		Example("https://example.com")
	})
})

var InboundOrderResponse = Type("InboundOrderResponse", func() {
	Description("OrderRsp describes the order info")
	Field(1, "client_order_id", String, "client order id", func() {
		Example("xxx1234")
		MaxLength(50)
	})
	Field(2, "status", Int, "order status(1 准备揽件 2 运输中 3 已到库)", func() {
		Example(1)
	})
	Field(3, "platform_order_id", Int64, "platform order id", func() {
		Example(123)
	})
	Field(4, "tracking_number", String, "tracking number")
	Field(5, "tracking_url", String, "tracking url")
	Field(6, "items", ArrayOf(InboundOrderItems), "items")
	Field(7, "timestamp", String, "timestamp", func() {
		Example("2022-01-21 02:26:54")
	})

	Required("client_order_id", "status", "platform_order_id", "tracking_number", "tracking_url", "items", "timestamp")
})
