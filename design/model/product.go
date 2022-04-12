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
	Field(1, "images", ArrayOf(String), "images", func() {
		Example([]string{"url"})
	})
	Field(2, "sku", String, "sku", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(3, "name", String, "name", func() {
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
	Field(7, "weight", Float64, "weight（g）", func() {
		Example(10.0)
	})
	Field(8, "length", Float64, "length（mm）", func() {
		Example(1.0)
	})
	Field(9, "width", Float64, "width（mm）", func() {
		Example(1.0)
	})
	Field(10, "height", Float64, "height（mm）", func() {
		Example(2.0)
	})
	Field(11, "hs_code", String, "hs code", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(12, "barcode", String, "barcode", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(13, "attributes", ArrayOf(String), "attributes", func() {
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
	Field(20, "error_message", String, "error message")
	Field(21, "material", String, "Material", func() {
		Default("")
	})
	Field(22, "purpose", String, "Purpose", func() {
		Default("")
	})

	Extend(AuthToken)

	Required("sku", "name", "declared_en_name", "declared_cn_name",
		"declared_value_in_usd", "weight", "length", "width", "height", "hs_code",
		"barcode", "attributes", "customer_code", "declared_value_in_eur")
})

var MultiProduct = Type("MultiProduct", func() {
	Description("MultiProduct describes the create multi product")
	Field(1, "products", ArrayOf(Product), "product info", func() {
		MinLength(1)
	})
	Extend(AuthToken)

	Required("products")
})

var MultiProductRsp = Type("MultiProductRsp", func() {
	Description("MultiProductRsp describes the multi product info")
	Extend(BaseResponse)
	Field(1, "data", MultiProductInfo, "data")
})

var MultiProductInfo = Type("MultiProductInfo", func() {
	Description("MultiProductRsp describes the multi product info")
	Field(1, "products", ArrayOf(MultiProductData), "product info")
	Required("products")
})

var MultiProductData = Type("MultiProductData", func() {
	Description("MultiProductRsp describes the multi product info")
	Field(1, "product_id", Int64, "product id")
	Field(2, "status", String, "status")
	Field(3, "error_msg", String, "error msg")
	Field(4, "barcode", String, "barcode")
	Field(5, "sku", String, "sku")
	Field(6, "product_name", String, "product_name")
	Required("product_id", "status", "error_msg", "barcode", "sku", "product_name")
})

var ProductQueryPayload = Type("ProductQueryPayload", func() {
	Description("Product Query Payload")
	Field(1, "id", ArrayOf(String), "id", func() {
		Example([]string{"1", "2", "3"})
	})
	Field(2, "sku", String, "sku", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(3, "barcode", String, "barcode", func() {
		Example("xxx")
	})
	Field(4, "status", String, "status", func() {
		Example("1")
	})
	Field(5, "attributes", ArrayOf(String), "attributes", func() {
		Example([]string{"liquid", "battery", "cosmetic", "magnetic"})
	})
	Field(6, "name", String, "name", func() {
		Example("xxx")
	})
	Field(7, "inventory", String, "inventory", func() {
		Example("with")
	})
	Extend(Pagination)
	Extend(AuthToken)
})

var ExportProductResult = Type("ExportProductResult", func() {
	Description("ExportProductResult")
	Attribute("length", Int64, "Length is the downloaded content length in bytes.", func() {
		Example(4 * 1024 * 1024)
	})
	Required("length")
})

var UploadProductPayload = Type("UploadProductPayload", func() {
	Description("Upload Product By Excel")
	Field(1, "file", Bytes, "file", func() {
		Example("1.xlsx")
	})
	Extend(AuthToken)
})

var UploadProductResponse = Type("UploadProductResponse", func() {
	Description("Upload Product By Excel")
	Extend(BaseResponse)
	Field(1, "data", UploadProductResponseData, "data")
})

var UploadProductResponseData = Type("UploadProductResponseData", func() {
	Description("Upload Product Response")
	Field(1, "total_count", Int, "Total Count", func() {
		Example(100)
	})
	Field(2, "success_count", Int, "Success Count", func() {
		Example(955)
	})
	Field(3, "fail_count", Int, "Fail Count", func() {
		Example(45)
	})
	Field(4, "result_file", String, "Result File", func() {
		Example("product_result.excel (base64)")
	})
	Required("result_file")
})

var ProductsQueryReq = Type("ProductsQueryReq", func() {
	Description("ProductsQueryReq describes the query products")
	Field(1, "name", String, "name", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(2, "sku", String, "sku")
	Field(3, "barcode", String, "barcode")
	Field(4, "attributes", ArrayOf(String), "attributes", func() {
		Example([]string{"normal", "liquid", "electric", "cosmetic", "magnetic"})
	})
	Field(5, "status", Int, "status", func() {
		Example(1)
		Enum(0, 1, 10)
	})
	Field(6, "inventory", Boolean, "inventory", func() {
		Example(false)
		Enum(false, true)
	})
	Extend(Pagination)
	Extend(AuthToken)
})

var ProductsQueryRsp = Type("ProductsQueryRsp", func() {
	Description("ProductsQueryRsp describes the product query return value")
	Extend(BaseResponse)
	Field(1, "data", ProductsQueryData, "data")
})

var ProductsQueryData = Type("ProductsQueryData", func() {
	Description("ProductsQueryData describes the query products resp")
	Field(1, "list", ArrayOf(ProductItem), "list")
	Field(2, "meta", MetaData, "meta")
	Required("list")
})

var ProductItem = Type("ProductItem", func() {
	Description("ProductItem describes product item")
	Field(1, "id", Int32, "id")
	Field(2, "status", Int, "status")
	Field(3, "barcode", String, "barcode")
	Field(4, "sku", String, "sku")
	Field(5, "name", String, "name")
	Field(6, "attributes", ArrayOf(String), "attributes")
	Field(7, "images", ArrayOf(String), "images")
	Field(8, "inventory", Int, "inventory")
	Field(9, "weight", Float64, "weight")
	Field(10, "inbound_weight", Float64, "inbound_ eight")
	Field(11, "length", Float64, "length")
	Field(12, "width", Float64, "width")
	Field(13, "height", Float64, "height")

	Required("id", "status", "barcode", "sku", "name", "attributes", "images",
		"inventory", "weight", "inbound_weight", "length", "width", "height")
})

var ProductDetailReq = Type("ProductDetailReq", func() {
	Description("ProductDetailReq describes the query product detail")
	Field(1, "id", Int32, "id")
	Extend(AuthToken)
})

var ProductDetailRsp = Type("ProductDetailRsp", func() {
	Description("ProductsQueryRsp describes the product query return value")
	Extend(BaseResponse)
	Field(1, "data", ProductDetailData, "data")
})

var ProductDetailData = Type("ProductDetailData", func() {
	Description("ProductItem describes product item")
	Field(1, "id", Int32, "id")
	Field(2, "status", Int, "status")
	Field(3, "barcode", String, "barcode")
	Field(4, "sku", String, "sku")
	Field(5, "name", String, "name")
	Field(6, "attributes", ArrayOf(String), "attributes")
	Field(7, "images", ArrayOf(String), "images")
	Field(8, "weight", Float64, "weight")
	Field(9, "hs_code", String, "hs_code")
	Field(10, "declared_cn_name", String, "declared_cn_name")
	Field(11, "declared_en_name", String, "declare_en_name")
	Field(12, "declared_value_in_usd", Float64, "declared_value_in_usd")
	Field(13, "declared_value_in_eur", Float64, "declared_value_in_eur")
	Field(14, "barcode_service", Boolean, "barcode_service")

	Required("id", "status", "barcode", "sku", "name", "attributes", "images",
		"hs_code", "declared_cn_name", "declared_en_name", "declared_value_in_usd",
		"declared_value_in_eur", "barcode_service", "weight")
})

var DownloadTemplatesReq = Type("DownloadTemplatesReq", func() {
	Field(1, "template", String, "template")
	Extend(AuthToken)
})
