package service

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl"
) //nolint:revive

var _ = Service("quote", func() {
	Description("The quote service performs operations on quotation")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1/shipping-estimate")
	})

	Method("get", func() {
		Payload(model.GetQuote)
		Result(model.QuoteRsp)
		HTTP(func() {
			GET("/")
			Header("Authorization:Authorization")
			Params(func() {
				Param("origin_country:origin_country")
				Param("dest_country:dest_country")
				Param("dest_state:dest_state")
				Param("dest_zip_code:dest_zip_code")
				Param("weight:weight")
				Param("length:length")
				Param("width:width")
				Param("height:height")
				Param("product_attributes:product_attributes")
				Param("factory:factory")
				Param("date:date")
				Required("origin_country", "dest_country", "dest_state", "dest_zip_code",
					"weight", "length", "width", "height")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("post", func() {
		Payload(model.PostQuote)
		Result(model.UserRsp)
		HTTP(func() {
			POST("/")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
