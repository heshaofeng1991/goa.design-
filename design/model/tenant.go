/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    tenant
	@Date    2022/3/21 11:32 上午
	@Desc
*/

package model

import (
	. "goa.design/goa/v3/dsl"
)

var TenantIntegrations = Type("TenantIntegrations", func() {
	Description("A list of tenant's integrations")

	Extend(BaseResponse)

	Field(1, "data", ArrayOf(IntegrationData), "data", func() {
		Description("List of integrations")
	})
	Extend(AuthToken)
})

var IntegrationData = Type("IntegrationData", func() {
	Description("Integration")
	Field(1, "store", String, "store", func() {
		Description("NSS store")
	})
	Field(2, "platform", String, "platform", func() {
		Description("shopify")
	})
	Field(3, "created_at", String, "created_at", func() {
		Example("2020-01-01T00:00:00Z")
	})
	Field(4, "integration_at", String, "updated_at", func() {
		Example("2020-01-01T00:00:00Z")
	})
})

// TenantRsp 响应结构.
var TenantRsp = Type("TenantRsp", func() {
	Field(1, "data", TenantData, "data")
	Extend(BaseResponse)
})

// TenantData 响应data结构.
var TenantData = Type("TenantData", func() {
	Field(1, "tenant_id", Int32, "tenant id")
	Field(2, "tenant_code", String, "tenant code")
	Field(3, "currency", String, "currency")
	Field(4, "balance", Float64, "balance")
	Field(5, "handling_fee", Float64, "handling_fee")
	Field(6, "preset_channel_id", ArrayOf(Int32), "preset_channel_ids")
	Field(7, "test_channel_id", ArrayOf(Int32), "test_channel_id")
	Field(8, "first_inbound_at", String, "first_inbound_at")
	Field(9, "storage_unit_price", Float64, "storage_unit_price")
	Field(10, "shipping_option", Int32, "shipping_option")
	Field(11, "default_warehouse", Int32, "default warehouse")
	Field(12, "country_code", String, "country code of tariff number")
	Field(13, "us_tariff_number", String, "us tariff number")
	Field(14, "uk_tariff_number", String, "uk tariff number")
	Field(15, "prepay_tariff", Boolean, "prepay tariff")
})

// TenantInfo 更新tenant info结构.
var TenantInfo = Type("TenantInfo", func() {
	Field(1, "shipping_option", Int32, "shipping_option")
	Field(2, "default_warehouse", Int32, "default warehouse")
	Field(3, "country_code", String, "country code of tariff number")
	Field(4, "us_tariff_number", String, "us tariff number")
	Field(5, "uk_tariff_number", String, "uk tariff number")
	Field(6, "prepay_tariff", Boolean, "prepay tariff")
	Extend(AuthToken)
})
