/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    product
	@Date    2022/2/10 11:02 上午:00
	@Desc
*/

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

var Product = Type("Product", func() {
	Description("Product describes the create product")
	Field(1, "product_image", ArrayOf(String), "product image", func() {
		Example([]string{"url"})
	})
	Field(2, "product_sku", String, "product sku", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(3, "product_name", String, "product name", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(4, "declared_en_name", String, "declared en name", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(5, "declared_cn_name", String, "declared cn name", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(6, "declared_value_in_usd", Float64, "declared value in usd（$）", func() {
		Minimum(0)
		Example(50.0)
	})
	Field(7, "product_weight", Float64, "product weight（g）", func() {
		Example(10.0)
	})
	Field(8, "product_length", Float64, "product length（mm）", func() {
		Example(1.0)
	})
	Field(9, "product_width", Float64, "product width（mm）", func() {
		Example(1.0)
	})
	Field(10, "product_height", Float64, "product height（mm）", func() {
		Example(2.0)
	})
	Field(11, "hs_code", String, "hs code", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(12, "product_barcode", String, "product_barcode", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(13, "product_attributes", ArrayOf(String), "product attributes", func() {
		Example([]string{"liquid", "battery", "cosmetic", "magnetic"})
	})
	Field(14, "qty", Int, "quality", func() {
		Example(1)
	})
	Field(15, "enabled_nss_barcode", Boolean, "need use nss barcode", func() {
		Example(false)
		Enum(false, true)
	})
	Field(16, "declared_value_in_eur", Float64, "declared value in eur（€）", func() {
		Minimum(0)
		Example(50.0)
	})
	Field(17, "customer_code", String, "customer code", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(18, "id", Int32, "id", func() {
		Example(1)
		Minimum(1)
	})
	Field(19, "barcode_service", Boolean, "barcode_service", func() {
		Example(false)
		Enum(false, true)
	})

	Required("product_sku", "product_name", "declared_en_name", "declared_cn_name",
		"declared_value_in_usd", "product_weight", "product_length", "product_width",
		"product_height", "hs_code", "product_barcode", "product_attributes", "customer_code",
		"declared_value_in_eur", "barcode_service")
})

var MultiProduct = Type("MultiProduct", func() {
	Description("MultiProduct describes the create multi product")
	Field(1, "products", ArrayOf(Product), "product info", func() {
		MinLength(1)
	})

	Required("products")
})

var MultiProductRsp = Type("MultiProductRsp", func() {
	Description("MultiProductRsp describes the multi product info")
	Field(1, "results", ArrayOf(MultiProductData), "results")

	Required("results")
})

var MultiProductData = Type("MultiProductData", func() {
	Description("MultiProductRsp describes the multi product info")
	Field(1, "product_id", Int64, "product id")
	Field(2, "status", String, "status")
	Field(3, "error_msg", String, "error msg")
	Field(4, "barcode", String, "barcode")
	Field(5, "sku", String, "sku")
	Field(5, "product_name", String, "product_name")

	Required("product_id", "status", "error_msg", "barcode", "sku", "product_name")
})
