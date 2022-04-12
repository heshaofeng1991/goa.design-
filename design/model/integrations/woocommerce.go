package model

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl"
)

var WoocommerceReturnArgs = Type("WoocommerceReturnArgs", func() {
	Description("Woocommerce Return Arguments")
	Field(1, "user_id", String, "User ID", func() {
		Example("1")
	})
	Field(2, "success", String, "Success", func() {
		Example("0")
	})
})

var WoocommerceReturnResult = Type("WoocommerceReturnResult", func() {
	Description("Woocommerce Return Result")
	Field(1, "success", String, "Success", func() {
		Example("0")
	})
})

var WoocommerceCallbackArgs = Type("WoocommerceCallbackArgs", func() {
	Description("WooCommerce callback args")
	Field(1, "key_id", Int, "key_id")
	Field(2, "user_id", String, "user_id")
	Field(2, "consumer_key", String, "consumer_key")
	Field(2, "consumer_secret", String, "consumer_secret")
	Field(2, "key_permissions", String, "key_permissions")
	Required("key_id", "user_id", "consumer_key", "consumer_secret", "key_permissions")
})

var WoocommerceCallbackRsp = Type("WoocommerceCallbackRsp", func() {
	Description("WooCommerce callback response")
	Field(1, "success", Boolean, "success")
	Required("success")
})

// GetWoocommerce 请求入参.
var GetWoocommerce = Type("GetWoocommerce", func() {
	Field(3, "store_id", Int32, "store_id")
	Field(4, "platform_ref_ids", ArrayOf(String), "platform_ref_ids")

	Extend(model.AuthToken)
})

// Woocommerce 返回出参.
var Woocommerce = Type("Woocommerce", func() {
	Description("WooCommerceOrders response")
	Field(1, "data", Any, "data")
	Extend(model.BaseResponse)
})
