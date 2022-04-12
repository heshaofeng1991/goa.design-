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
	Field(6, "country_code", String, "country code", func() {
		Example("US")
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
	Field(15, "inbound_order_number", String, "inbound number", func() {
		Example("246938764")
	})
	Field(16, "type", Int, "delivery mode(1 direct，2 warehouse, 3 platform)", func() {
		Enum(1, 2, 3)
		Example(1)
	})
	Field(17, "id", Int32, "outbound order id", func() {
		Example(1)
	})
	Field(18, "package_id", String, "package id", func() {
		Example("xxx")
	})

	Field(19, "store_id", Int32, "store id", func() {
		Example(0)
	})
	Field(20, "estimated_weight", Int64, "estimated weight", func() {
		Example(100)
	})
	Field(21, "platform_order_no", String, "platform order no", func() {
		Example("00001")
	})
	Field(22, "platform_created_at", String, "Platform created at", func() {
		Example("2022-03-10T06:35:33")
	})
	Field(23, "request_shipping_at", String, "request shipping at", func() {
		Example("2022-03-10 06:35:33")
	})
	Field(24, "remark", String, "remark", func() {
		Example("remark")
	})
	Field(25, "vat_number", String, "var number", func() {
		Example("var number")
	})
	Extend(AuthToken)

	Required("customer_order_id", "customer_code", "total_price", "currency",
		"enable_prepay_tariff", "shipping_type", "receiver_info", "items", "channel_id",
		"type")
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
	Field(9, "product_weight", Int, "product weight", func() {
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

	Field(14, "product_barcode", String, "产品barcode", func() {
		Example("xxx")
	})
	Field(15, "declared_value_in_eur", Float64, "declared value in eur（€）", func() {
		Minimum(0)
		Example(10.07)
	})
	Field(18, "requires_shipping", Boolean, "requires shipping", func() {
		Example(true)
	})
	Field(16, "ext_order_item_id", String, "Ext. order item id")
	Field(17, "ext_product_id", String, "Ext. order item id")
	Field(18, "platform_product_id", Int32, "Ext. order item id")

	Field(19, "material", String, "Material", func() {
		Default("")
		Example("")
	})
	Field(20, "purpose", String, "Purpose", func() {
		Default("")
		Example("")
	})

	Required("product_name", "product_sku", "product_price", "qty", "hs_code",
		"declared_cn_name", "declared_en_name", "declared_value_in_usd",
		"product_weight", "product_length", "product_width", "product_height",
		"product_attributes", "product_barcode", "declared_value_in_eur")
})

// OutboundOrderRsp 入库订单返回出参.
var OutboundOrderRsp = Type("OutboundOrderRsp", func() {
	Description("OutboundOrderRsp describes the outbound order return value")
	Field(1, "data", OutboundOrderRspData, "data")
	Extend(BaseResponse)
})

var OutboundOrderRspData = Type("OutboundOrderRspData", func() {
	Description("OutboundOrderRspData describes the outbound order return value")
	Field(1, "order_number", String, "order number", func() {
		Example("37942132")
	})
	Field(2, "tracking_number", String, "tracking number", func() {
		Example("xxx")
	})
	Required("order_number", "tracking_number")
})

var OrderRsp = Type("OrderRsp", func() {
	Description("OrderRsp describes the order info")
	Extend(BaseResponse)
	Field(1, "data", OrderInfo, "data")
})

var QueryOrder = Type("QueryOrder", func() {
	Field(1, "order_number", String, "inbound order number", func() {
		MinLength(1)
		Example("xxx1234")
	})
	Extend(AuthToken)
	Required("order_number")
})

var QueryOutOrder = Type("QueryOutOrder", func() {
	Field(1, "id", String, "outbound order number", func() {
		MinLength(1)
		Example("xxx1234")
	})
	Extend(AuthToken)
	Required("id")
})

var QueryOrderRsp = Type("QueryOrderRsp", func() {
	Description("QueryOrderRsp describes the order info")
	Extend(BaseResponse)
	Field(1, "data", OrderData, "data")
})

var OrderInfo = Type("OrderInfo", func() {
	Description("OrderInfo describes the order info")
	Field(1, "list", ArrayOf(OrderData), "outbounds data")
	Field(2, "meta", MetaData, "MetaData info")
	Required("list", "meta")
})

var OrderData = Type("OrderData", func() {
	Field(1, "customer_order_id", String, "customer order id", func() {
		Example("xxx1234")
		MaxLength(50)
	})
	Field(2, "status", Int, "order status", func() {
		Example(1)
	})
	Field(3, "order_number", String, "outbound order number", func() {
		Example("123")
	})
	Field(4, "tracking_number", String, "tracking number")
	Field(5, "tracking_url", String, "tracking url")
	Field(6, "items", ArrayOf(OutboundOrderItem), "items")

	Required("customer_order_id", "status", "order_number", "tracking_number", "tracking_url", "items")
})

var OutboundOrderItemCreateRequest = Type("OutboundOrderItemCreateRequest", func() {
	Field(1, "outbound_order_id", Int64, "outbound order id", func() {
		Example(1)
	})
	Field(2, "product_sku", String, "product SKU", func() {
		Example("xxxx")
	})
	Field(3, "ext_order_item_id", String, "ext order item id", func() {
		Example("xxxx")
	})
	Field(4, "product_name", String, "product name", func() {
		Example("NSS mate40")
	})
	Field(3, "product_price", Float64, "product price", func() {
		Example(10.3)
	})
	Field(5, "barcode", String, "产品barcode", func() {
		Example("xxx")
	})
	Field(6, "qty", Int, "产品数量", func() {
		Example(1)
	})
	Field(7, "hs_code", String, "hs code", func() {
		Example("xxx")
	})
	Field(8, "declared_cn_name", String, "declared cn name", func() {
		Example("NSS")
	})
	Field(9, "declared_en_name", String, "declared en name", func() {
		Example("NSS")
	})
	Field(10, "declared_value_in_usd", Float64, "declared value in usd", func() {
		Minimum(0)
		Example(10.07)
	})
	Field(11, "declared_value_in_eur", Float64, "declared value in eur", func() {
		Minimum(0)
		Example(9.8)
	})
	Field(12, "product_weight", Float64, "claim weight", func() {
		Example(10)
	})
	Field(13, "product_attributes", ArrayOf(String), "product attributes", func() {
		Example([]string{"battery", "cosmetic", "liquid", "magnetic"})
	})
	Field(14, "customer_code", String, "customer code", func() {
		Example("xxx")
	})
	Field(15, "images", ArrayOf(String), "images", func() {
		Example([]string{"url1", "url2"})
	})
	Field(16, "material", String, "material", func() {
		Example("xxx")
	})
	Field(17, "purpose", String, "purpose", func() {
		Example("xxx")
	})
	Field(18, "requires_shipping", Boolean, "requires shipping", func() {
		Example(true)
	})
	Field(19, "product_length", Float64, "length", func() {
		Example(10)
	})
	Field(20, "product_width", Float64, "width", func() {
		Example(10)
	})
	Field(21, "product_height", Float64, "height", func() {
		Example(10)
	})

	Required("outbound_order_id")
	Extend(AuthToken)
})

// OutboundOrderItemUpdateRequest 出库订单产品入参.
var OutboundOrderItemUpdateRequest = Type("OutboundOrderItemUpdateRequest", func() {
	Field(1, "outbound_order_id", Int64, "outbound order id", func() {
		Example(1)
	})
	Field(2, "product_sku", String, "product SKU", func() {
		Example("xxxx")
	})
	Field(3, "ext_order_item_id", String, "ext order item id", func() {
		Example("xxxx")
	})
	Field(4, "product_name", String, "product name", func() {
		Example("NSS mate40")
	})
	Field(3, "product_price", Float64, "product price", func() {
		Example(10.3)
	})
	Field(5, "barcode", String, "产品barcode", func() {
		Example("xxx")
	})
	Field(6, "qty", Int, "产品数量", func() {
		Example(1)
	})
	Field(7, "hs_code", String, "hs code", func() {
		Example("xxx")
	})
	Field(8, "declared_cn_name", String, "declared cn name", func() {
		Example("NSS")
	})
	Field(9, "declared_en_name", String, "declared en name", func() {
		Example("NSS")
	})
	Field(10, "declared_value_in_usd", Float64, "declared value in usd", func() {
		Minimum(0)
		Example(10.07)
	})
	Field(11, "declared_value_in_eur", Float64, "declared value in eur", func() {
		Minimum(0)
		Example(9.8)
	})
	Field(12, "product_weight", Float64, "claim weight", func() {
		Example(10)
	})
	Field(13, "product_attributes", ArrayOf(String), "product attributes", func() {
		Example([]string{"battery", "cosmetic", "liquid", "magnetic"})
	})
	Field(14, "customer_code", String, "customer code", func() {
		Example("xxx")
	})
	Field(15, "images", ArrayOf(String), "images", func() {
		Example([]string{"url1", "url2"})
	})
	Field(16, "material", String, "material", func() {
		Example("xxx")
	})
	Field(17, "purpose", String, "purpose", func() {
		Example("xxx")
	})
	Field(18, "requires_shipping", Boolean, "requires shipping", func() {
		Example(true)
	})
	Field(19, "product_length", Float64, "length", func() {
		Example(10)
	})
	Field(20, "product_width", Float64, "width", func() {
		Example(10)
	})
	Field(21, "product_height", Float64, "height", func() {
		Example(10)
	})
	Field(22, "ext_product_id", String, "ext order item id", func() {
		Example("xxxx")
	})
	Field(23, "platform_product_id", Int32, "ext order item id", func() {
		Example(0)
	})
	Field(24, "id", Int64, "outbound order item id", func() {
		Example(1)
	})

	Extend(AuthToken)
})

// OutboundOrderUpdateRequest 出库订单更新请求体
var OutboundOrderUpdateRequest = Type("OutboundOrderUpdateRequest", func() {
	Description("OutboundOrders describes an order to outbound order.")
	Description("OutboundOrders describes an order to outbound order.")
	Field(1, "id", Int32, "outbound order id", func() {
		Example(1)
	})
	Field(6, "warehouse_id", Int, "warehouse id", func() {
		Example(1)
	})
	Field(7, "offline", Boolean, "offline", func() {
		Example(true)
	})
	Field(8, "enable_prepay_tariff", Boolean, "enable prepay tariff", func() {
		Example(false)
	})
	Field(9, "customer_tariff_number_type", Int, "customer tariff number of type", func() {
		Example(0)
	})
	Field(10, "customer_tariff_country_code", String, "customer tariff country code", func() {
		Example("US")
	})
	Field(11, "customer_tariff_number", String, "customer tariff number", func() {
		Example("xxx")
	})
	Field(12, "receiver_info", Address, "receiver info")
	Field(13, "items", ArrayOf(OutboundOrderItemUpdateRequest), "order items")
	Field(14, "description", String, "description", func() {
		Example("description")
	})

	Required("id")
	Extend(AuthToken)
})

var BatchUpdateOrderRequest = Type("BatchUpdateOrderRequest ", func() {
	Description("Batch Update Order Request")
	Extend(AuthToken)
	Field(1, "orders", ArrayOf(OutboundOrderUpdateRequest), "Update Order Data")
})

var BatchUpdateOrderResponse = Type("BatchUpdateOrderResponse", func() {
	Description("Batch Update Order Response")
	Extend(BaseResponse)
	Field(1, "data", ArrayOf(BatchUpdateResult), "Data")
})

var BatchUpdateResult = Type("BatchUpdateResult", func() {
	Description("update status")
	Field(1, "id", Int32, "outbound order id", func() {
		Example(1)
	})
	Field(2, "success", Boolean, "success", func() {
		Example(true)
	})
	Field(3, "message", String, "message", func() {
		Example("成功")
	})

	Required("id", "success")
})

var DeleteOutboundItemRequest = Type("DeleteOutboundItemRequest", func() {
	Description("delete request")
	Field(1, "id", Int64, "id", func() {
		Example(1)
	})

	Required("id")
	Extend(AuthToken)
})
