package service

import (
	"goa/design/model"
	inter "goa/design/model/integrations"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("integrations", func() {
	Description("Integrations of other platforms")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1/integrations")
	})

	Method("list", func() {
		Payload(model.AuthToken)
		Result(inter.IntegrationListResult)
		HTTP(func() {
			GET("/")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("authorize", func() {
		Payload(inter.AuthorizePlatform)
		Result(inter.AuthorizePlatformRsp)
		HTTP(func() {
			POST("/authorize")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
