/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    outboundorder
	@Date    2022/2/10 11:04 上午:00
	@Desc
*/

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

// OutboundOrder 出库订单请求体
var OutboundOrder = Type("OutboundOrder", func() {
	Description("OutboundOrders describes an order to outbound order.")
	Field(1, "customer_order_id", String, "customer order id", func() {
		Example("YT000001")
	})
	Field(2, "customer_code", String, "customer code", func() {
		MinLength(1)
		MaxLength(50)
		Example("YT")
	})
	Field(3, "total_price", Float64, "total price", func() {
		Example(980.67)
	})
	Field(4, "currency", String, "currency", func() {
		Example("USD")
	})
	Field(5, "customer_tariff_number", String, "customer tariff number", func() {
		Example("xxx")
	})
	Field(6, "customer_tariff_number_type", Int, "customer tariff number of type", func() {
		Example(0)
	})
	Field(7, "enable_prepay_tariff", Boolean, "enable prepay tariff", func() {
		Example(false)
	})
	Field(8, "shipping_type", Int, "shipping type(1 Economic, 2 Fastest, 3 Recommended)", func() {
		Enum(1, 2, 3)
		Example(1)
	})
	Field(9, "receiver_info", ShippingAddress, "receiver info")
	Field(10, "items", ArrayOf(OutboundOrderItem), "order items", func() {
		MinLength(1)
	})
	Field(12, "channel_id", Int, "channel id", func() {
		Example(1)
	})
	Field(13, "description", String, "description", func() {
		Example("description")
	})
	Field(15, "inbound_order_id", Int64, "入库单ID", func() {
		Example(1)
	})
	Field(16, "type", Int, "delivery mode(1 direct，2 warehouse)", func() {
		Enum(1, 2)
		Example(1)
	})
	Field(17, "id", Int32, "outbound order id", func() {
		Example(1)
	})
	Field(18, "package_id", Int64, "package id", func() {
		Example(1)
	})

	Required("customer_order_id", "customer_code", "total_price", "currency",
		"customer_tariff_number", "customer_tariff_number_type", "enable_prepay_tariff",
		"shipping_type", "receiver_info", "items", "channel_id", "description",
		"inbound_order_id", "type")
})

// OutboundOrderItem 出库订单返回出参.
var OutboundOrderItem = Type("OutboundOrderItem", func() {
	Field(1, "product_name", String, "product name", func() {
		Example("NSS mate40")
	})
	Field(2, "product_sku", String, "product SKU", func() {
		Example("xxxx")
	})
	Field(3, "product_price", Float64, "product price", func() {
		Example(10.3)
	})
	Field(4, "qty", Int, "产品数量", func() {
		Example(1)
	})
	Field(5, "hs_code", String, "hs code", func() {
		Example("xxx")
	})
	Field(6, "declared_cn_name", String, "declared cn name", func() {
		Example("NSS")
	})
	Field(7, "declared_en_name", String, "declared en name", func() {
		Example("NSS")
	})
	Field(8, "declared_value_in_usd", Float64, "declared value in usd", func() {
		Minimum(0)
		Example(10.07)
	})
	Field(9, "product_weight", Float64, "product weight", func() {
		Example(10)
	})
	Field(10, "product_length", Int, "product length", func() {
		Example(10)
	})
	Field(11, "product_width", Int, "product width", func() {
		Example(10)
	})
	Field(12, "product_height", Int, "product height", func() {
		Example(10)
	})
	Field(13, "product_attributes", ArrayOf(String), "product attributes", func() {
		Example([]string{"battery", "cosmetic", "liquid", "magnetic"})
	})

	Field(14, "barcode", String, "产品barcode", func() {
		Example("xxx")
	})

	Field(15, "declared_value_in_eur", Float64, "declared value in eur（€）", func() {
		Minimum(0)
		Example(10.07)
	})

	Required("product_name", "product_sku", "product_price", "qty", "hs_code",
		"declared_cn_name", "declared_en_name", "declared_value_in_usd",
		"product_weight", "product_length", "product_width", "product_height",
		"product_attributes", "barcode", "declared_value_in_eur")
})

// OutboundOrderRsp 入库订单返回出参.
var OutboundOrderRsp = Type("OutboundOrderRsp", func() {
	Description("OutboundOrderRsp describes the outbound order return value")
	Field(1, "outbound_order_id", Int64, "outbound order id", func() {
		Example(1)
	})
	Field(2, "tracking_number", String, "tracking number", func() {
		Example("xxx")
	})
})

var OrderRsp = Type("OrderRsp", func() {
	Description("OrderRsp describes the order info")
	Field(1, "client_order_id", String, "client order id", func() {
		Example("xxx1234")
		MaxLength(50)
	})
	Field(2, "status", Int, "order status", func() {
		Example(1)
	})
	Field(3, "platform_order_id", Int64, "platform order id", func() {
		Example(123)
	})
	Field(4, "tracking_number", String, "tracking number")
	Field(5, "tracking_url", String, "tracking url")
	Field(6, "items", ArrayOf(OutboundOrderItem), "items")

	Required("client_order_id", "status", "platform_order_id", "tracking_number", "tracking_url", "items")
})
