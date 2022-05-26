package service

import (
	. "goa.design/goa/v3/dsl"
	"goa/design/model"
)

var _ = Service("quote", func() {
	Description("The quote service performs operations on quotation")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/")
	})

	Method("UpdateChannelCostStatus", func() {
		Payload(model.UpdateChannelCostStatusReq)
		Result(model.UpdateChannelCostStatusRsp)
		HTTP(func() {
			GET("/costs")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
