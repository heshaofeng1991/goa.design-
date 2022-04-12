package model

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl"
)

var AuthorizePlatform = Type("Authorize", func() {
	Description("Apply for a token from platform")
	Field(1, "platform", String, "platform", func() {
		Example("woocommerce")
		MaxLength(50)
	})
	Field(2, "host", String, "host", func() {
		Example("127.0.0.1")
		MaxLength(50)
	})
	Extend(model.AuthToken)
	Required("platform", "host")
})

var AuthorizePlatformRspData = Type("AuthorizePlatformRspData", func() {
	Description("Response data for AuthorizePlatform")
	Field(1, "redirect_url", String, "redirect_url", func() {
		Example("http://wp-woo.nextsmartship.com/wc-auth/v1/authorize?app_name=NextSmartShip&scope=read_write&user_id=123&return_url=http://app.com/return-page&callback_url=https://app.com/callback-endpoint")
	})
	Required("redirect_url")
})

var AuthorizePlatformRsp = Type("AuthorizePlatformRsp", func() {
	Description("Apply for a token from platform response")
	Extend(model.BaseResponse)
	Field(1, "data", AuthorizePlatformRspData, "data")
})

var IntegrationListResult = Type("IntegrationListResult", func() {
	Description("Integration list result")
	Extend(model.BaseResponse)
	Field(1, "data", ArrayOf(Integration), "data")
})

var Integration = Type("Integration", func() {
	Description("Integration")
	Field(1, "id", Int32, "id", func() {
		Example(1)
	})
	Field(2, "name", String, "name", func() {
		Example("woocommerce")
		MaxLength(50)
	})
	Field(3, "status", Int, "status", func() {
		Enum(0, 1)
		Example(1)
	})
})
