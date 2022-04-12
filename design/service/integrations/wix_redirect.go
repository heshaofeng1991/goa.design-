package service

import (
	"goa/design/service"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("wix_redirect", func() {
	Description("integrations of Wix service")

	Security(service.JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1/integrations/wix")
	})

	Method("return_wix", func() {
		NoSecurity()
		HTTP(func() {
			GET("/return")
		})
	})
})
