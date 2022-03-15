package design

import (
	. "goa.design/goa/v3/dsl"
	"goa/design/model"
) //nolint:revive

var _ = Service("quote", func() {
	Description("The quote service performs operations on quotation")

	Error("Unauthorized")

	HTTP(func() {
		Path("/shipping-estimate")
	})

	Method("get", func() {
		Payload(model.GetQuote)
		Result(ArrayOf(model.Quote))
		HTTP(func() {
			GET("/")
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
					"weight", "length", "width", "height", "product_attributes")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
