/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    product
	@Date    2022/1/23 10:42 PM:00
	@Desc
*/

package service

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("product", func() {
	Description("The product service performs operations on product")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1")
	})

	Method("batches_create_product", func() {
		Payload(model.MultiProduct)
		Result(model.MultiProductRsp)
		HTTP(func() {
			POST("/products")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_product", func() {
		Payload(model.Product)
		Result(model.UpdateResponse)
		HTTP(func() {
			PUT("/products")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("export_product", func() {
		Payload(model.ProductQueryPayload)
		Result(model.ExportProductResult)

		Error("internal_error", ErrorResult, "Fault while processing download.")

		HTTP(func() {
			GET("/products/export")
			Header("Authorization:Authorization")
			Params(func() {
				Param("id")
				Param("sku")
				Param("barcode")
				Param("status")
				Param("attributes")
				Param("name")
				Param("inventory")
				Param("current")
				Param("page_size")
			})
			SkipResponseBodyEncodeDecode()

			Response(func() {
				Header("length:Content-Length")
			})
			Response("Unauthorized", StatusUnauthorized)
			Response("internal_error", StatusInternalServerError)
		})
	})

	Method("download_templates", func() {
		Payload(model.DownloadTemplatesReq)
		Result(model.ExportProductResult)

		Error("internal_error", ErrorResult, "Fault while processing download.")

		HTTP(func() {
			GET("/download-templates/{template}")
			Header("Authorization:Authorization")
			SkipResponseBodyEncodeDecode()
			Response(func() {
				Header("length:Content-Length")
			})
			Response("Unauthorized", StatusUnauthorized)
			Response("internal_error", StatusInternalServerError)
		})
	})

	Method("upload_product", func() {
		Payload(model.UploadProductPayload)
		Result(model.UploadProductResponse)
		HTTP(func() {
			POST("/products/upload")
			Header("Authorization:Authorization")
			Response(StatusOK)
			MultipartRequest()
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("upload_product_update", func() {
		Payload(model.UploadProductPayload)
		Result(model.UploadProductResponse)
		HTTP(func() {
			PUT("/products/upload")
			Header("Authorization:Authorization")
			Response(StatusOK)
			MultipartRequest()
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("generate_barcode", func() {
		Payload(model.AuthToken)
		Result(model.BarCodeRsp)
		HTTP(func() {
			GET("/barcode")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("products_query", func() {
		Payload(model.ProductsQueryReq)
		Result(model.ProductsQueryRsp)
		HTTP(func() {
			GET("/products")
			Header("Authorization:Authorization")
			Params(func() {
				Param("name")
				Param("sku")
				Param("barcode")
				Param("status")
				Param("attributes")
				Param("inventory")
				Param("current")
				Param("page_size")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("product_detail", func() {
		Payload(model.ProductDetailReq)
		Result(model.ProductDetailRsp)
		HTTP(func() {
			GET("/products/{id}")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
