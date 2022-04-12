package service

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("auth", func() {
	Description("The auth service")

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1")
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
