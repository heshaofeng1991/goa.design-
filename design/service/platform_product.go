/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    platform_product
	@Date    2022/3/21 4:36 下午
	@Desc
*/

package service

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("platform_product", func() {
	Description("The platform_product service performs operations on platform product")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1")
	})

	Method("platform_product", func() {
		Payload(model.GetListing)
		Result(model.MultiListingRsp)
		HTTP(func() {
			GET("/platform-products")
			Header("Authorization:Authorization")
			Params(func() {
				Param("platform_status")
				Param("name")
				Param("page_size")
				Param("current")
				Param("id")
				Param("listing_sku")
				Param("sku")
				Param("is_mapping")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("convert", func() {
		Payload(model.ConvertPlatformProductsReq)
		Result(model.ConvertPlatformProductsRes)
		HTTP(func() {
			POST("/platform-products/convert")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("mappings", func() {
		Payload(model.UpdateMappingsReq)
		Result(model.UpdateMappingsRes)
		HTTP(func() {
			POST("/platform-products/mappings")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
