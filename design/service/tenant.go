/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    tenant
	@Date    2022/3/21 10:40 上午
	@Desc
*/

package service

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl"
)

var _ = Service("tenant", func() {
	Description("The tenant service performs operations on tenant")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1/tenants")
	})

	Method("integrations", func() {
		Payload(model.AuthToken)
		Result(model.TenantIntegrations)
		HTTP(func() {
			GET("/integrations")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_tenant_info", func() {
		Payload(model.AuthToken)
		Result(model.TenantRsp)
		HTTP(func() {
			GET("/")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_tenant_info", func() {
		Payload(model.TenantInfo)
		Result(model.UserRsp)
		HTTP(func() {
			PUT("/")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
