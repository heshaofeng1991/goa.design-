package service

import (
	model "goa/design/model/integrations"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("wix", func() {
	Description("integrations of Wix service")

	// Security(service.JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1/integrations/wix")
	})

	Method("callback_wix", func() {
		NoSecurity()
		Payload(model.WixCallbackArgs)
		Result(model.WixCallbackRsp)
		HTTP(func() {
			Params(func() {
				Param("code")
				Param("state")
				Param("instanceId")
			})
			GET("/callback")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("webhooks_products_wix", func() {
		NoSecurity()
		Payload(model.WebHooksCallBackProducts)
		Result(model.WebHooksCallBackProductsResp)
		HTTP(func() {
			POST("/webhooks/products")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("products_list", func() {
		NoSecurity()
		Payload(model.GetProductsListReq)
		Result(model.GetProductsListResp)
		HTTP(func() {
			POST("/products")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("orders_list", func() {
		NoSecurity()
		Payload(model.GetOrdersListReq)
		Result(model.GetOrdersListResp)
		HTTP(func() {
			POST("/orders")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
