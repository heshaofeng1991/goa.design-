package model

import (
	. "goa.design/goa/v3/dsl"
)

var GenerateTokenReq = Type("GenerateTokenReq", func() {
	Description("GenerateTokenReq describes the generate token")

	Field(1, "id", Int64, "user_id", func() {
		Example(1)
		Minimum(1)
	})
	Field(1, "tenant_id", Int64, "tenantID", func() {
		Example(1)
		Minimum(1)
	})

	Required("id", "tenant_id")
})

var GenerateTokenRsp = Type("GenerateTokenRsp", func() {
	Description("GenerateTokenRsp describes the response of generate token")
	Extend(BaseResponse)
	Field(1, "data", GenerateTokenData, "data")
})

var GenerateTokenData = Type("GenerateTokenData", func() {
	Description("GenerateTokenData describes the response of generate token")
	Field(1, "token", String, "token")
	Required("token")
})
