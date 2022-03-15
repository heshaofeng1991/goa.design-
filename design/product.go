/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    product
	@Date    2022/1/23 10:42 PM:00
	@Desc
*/

package design

import (
	. "goa.design/goa/v3/dsl" //nolint:revive
	"goa/design/model"
)

var _ = Service("product", func() {
	Description("The product service performs operations on product")

	Error("Unauthorized")

	HTTP(func() {
		Path("/")
	})

	Method("batches_create_product", func() {
		Payload(model.MultiProduct)
		Result(model.MultiProductRsp)
		HTTP(func() {
			POST("/products")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_product", func() {
		Payload(model.Product)
		Result(model.UpdateResponse)
		HTTP(func() {
			PUT("/products")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("generate_barcode", func() {
		Payload(model.BarCode)
		Result(model.BarCodeRsp)
		HTTP(func() {
			POST("/barcode")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("generate_token", func() {
		Payload(model.GenerateTokenReq)
		Result(model.GenerateTokenRsp)
		HTTP(func() {
			POST("/token")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
