package service

import (
	model "goa/design/model/integrations"
	"goa/design/service"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("woocommerce", func() {
	Description("integrations of Woocommerce service")

	Security(service.JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1/integrations/woocommerce")
	})

	Method("return_woocommerce", func() {
		NoSecurity()
		Payload(model.WoocommerceReturnArgs)
		Result(model.WoocommerceReturnResult)
		HTTP(func() {
			GET("/return")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("callback_woocommerce", func() {
		NoSecurity()
		Payload(model.WoocommerceCallbackArgs)
		Result(model.WoocommerceCallbackRsp)
		HTTP(func() {
			POST("/callback")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("retrieve_orders", func() {
		Payload(model.GetWoocommerce)
		Result(model.Woocommerce)
		HTTP(func() {
			GET("/retrieve/orders")
			Header("Authorization:Authorization")
			Params(func() {
				Param("store_id")
				Param("platform_ref_ids")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("retrieve_products", func() {
		Payload(model.GetWoocommerce)
		Result(model.Woocommerce)
		HTTP(func() {
			GET("/retrieve/products")
			Header("Authorization:Authorization")
			Params(func() {
				Param("store_id")
				Param("platform_ref_ids")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
