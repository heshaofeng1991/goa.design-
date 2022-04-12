package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

var Listing = Type("Listing", func() {
	Description("Listing describes the listing")
	Field(1, "listing_sku", String, "listing sku", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(2, "barcode", String, "barcode", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(3, "name", String, "name", func() {
		Example("xxx")
		MaxLength(50)
	})
	Field(4, "images", ArrayOf(String), "images", func() {
		Example([]string{"xxx"})
	})
	Field(5, "vendor", String, "vendor", func() {
		Example("xxx")
	})
	Field(6, "platform_status", Int, "platform status", func() {
		Example(1)
	})
	Field(7, "id", Int32, "id", func() {
		Example(1)
		Minimum(1)
	})
	Field(8, "store", Store, "store")
	Field(9, "mappings", ArrayOf(Mapping), "mappings")
	Field(10, "is_mapping", Boolean, "is_mapping")
})

var Mapping = Type("Mapping", func() {
	Description("Mappings describes the Mappings")
	Field(1, "id", Int32, "id", func() {
		Example(1)
	})
	Field(2, "platform_product_id", Int32, "platform_product_id", func() {
		Example(1)
	})
	Field(3, "product_id", Int32, "product_id", func() {
		Example(1)
	})
	Field(4, "product_sku", String, "product_sku", func() {
		Example("xxx")
	})
	Field(5, "qty", Int32, "qty", func() {
		Example(1)
	})
	Field(6, "product_name", String, "product_name", func() {
		Example("xxx")
	})
	Field(7, "images", ArrayOf(String), "images", func() {
		Example([]string{"xxx"})
	})
	Required("id", "platform_product_id", "product_id",
		"product_sku", "qty", "product_name", "images")
})

var MultiListing = Type("MultiPlatformProduct", func() {
	Description("MultiProduct describes the create multi product")
	Field(1, "meta", MetaData, "MetaData info")
	Field(2, "list", ArrayOf(Listing), "product info")

	Required("meta", "list")
})

var GetListing = Type("GetListing", func() {
	Description("GetOrder describes the order info")
	Field(1, "current", Int, "current", func() {
		Minimum(1)
		Example(1)
	})
	Field(2, "page_size", Int, "page_size", func() {
		Example(10)
		Enum(10, 20, 50, 100)
	})
	Field(3, "name", String, "name", func() {
		Example("xxx")
	})
	Field(4, "platform_status", Int, "platform_status", func() {
		Example(1)
	})
	Field(5, "id", Int32, "id", func() {
		Example(1)
	})
	Field(6, "listing_sku", String, "listing_sku", func() {
		Example("xxx")
	})
	Field(7, "sku", String, "sku", func() {
		Example("xxx")
	})
	Field(8, "is_mapping", Int, "is_mapping", func() {
		Example(1)
	})
	Extend(AuthToken)
})

var MultiListingRsp = Type("MultiPlatformProductRsp", func() {
	Description("MultiProductRsp describes the multi product info")
	Extend(BaseResponse)
	Field(3, "data", MultiListing, "data")
})

var ConvertPlatformProductsReq = Type("ConvertPlatformProductsReq", func() {
	Description("ConvertPlatformProductsReq describes ConvertPlatformProductsReq")
	Field(1, "platform_product_ids", ArrayOf(Int32), "platform_product_ids", func() {
		Example([]int{1})
	})
	Extend(AuthToken)
	Required("platform_product_ids")
})

var ConvertPlatformProductsResData = Type("ConvertPlatformProductsResData", func() {
	Description("ConvertPlatformProductsResData")
	Field(1, "id", Int32, "id", func() {
		Example(1)
	})
	Field(2, "success", Boolean, "success", func() {
		Example(true)
	})
	Field(2, "message", String, "message", func() {
		Example("xxx")
	})
	Required("id", "success", "message")
})

var ConvertPlatformProductsRes = Type("ConvertPlatformProductsRes", func() {
	Description("MultiProductRsp describes the multi product info")
	Extend(BaseResponse)
	Field(3, "data", ArrayOf(ConvertPlatformProductsResData), "data")
})

var UpdateMappingsProduct = Type("UpdateMappingsProduct", func() {
	Description("UpdateMappingsProduct")
	Field(1, "product_id", Int32, "product_id", func() {
		Example(1)
	})
	Field(2, "qty", Int32, "qty", func() {
		Example(1)
	})
	Required("product_id", "qty")
})

var UpdateMappingsReq = Type("UpdateMappingsReq", func() {
	Description("UpdateMappingsReq")
	Field(1, "platform_product_ids", ArrayOf(Int32), "platform_product_ids", func() {
		Example([]int32{1})
	})
	Field(2, "products", ArrayOf(UpdateMappingsProduct), "products")
	Extend(AuthToken)
	Required("platform_product_ids", "products")
})

var UpdateMappingsRes = Type("UpdateMappingsRes", func() {
	Description("UpdateMappingsRes")
	Extend(BaseResponse)
	Field(3, "data", Any, "data")
})
