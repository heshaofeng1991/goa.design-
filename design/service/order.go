// Package design @Author johnny 2022/1/13 1:39 PM:00
package service

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("order", func() {
	Description("The order service performs operations on order")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1")
	})

	Method("create_inbound_order", func() {
		Payload(model.InboundOrder)
		Result(model.InboundOrderRsp)
		HTTP(func() {
			POST("/inbound-orders")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_inbound_order", func() {
		Payload(model.InboundOrder)
		Result(model.UpdateResponse)
		HTTP(func() {
			PUT("/inbound-orders")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("create_pickup_order", func() {
		Payload(model.PickupOrder)
		Result(model.PickupOrderRsp)
		HTTP(func() {
			POST("/pickup-orders")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("batch_query_inbound_order", func() {
		Payload(model.GetOrder)
		Result(model.InboundOrderResponse)
		HTTP(func() {
			GET("/inbound-orders")
			Header("Authorization:Authorization")
			Params(func() {
				Param("order_numbers:order_numbers")
				Param("status:status")
				Param("current:current")
				Param("page_size:page_size")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_inbound_order", func() {
		Payload(model.QueryOrder)
		Result(model.QueryInboundOrderRsp)
		HTTP(func() {
			GET("/inbound-orders/{order_number}")
			Header("Authorization:Authorization")
			Params(func() {
				Param("order_number:order_number")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("create_outbound_order", func() {
		Payload(model.OutboundOrder)
		Result(model.OutboundOrderRsp)
		HTTP(func() {
			POST("/outbound-orders")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_outbound_order", func() {
		Payload(model.OutboundOrderUpdateRequest)
		Result(model.UpdateResponse)
		HTTP(func() {
			PUT("/outbound-orders/{id}")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("batch_update_outbound_order", func() {
		Payload(model.BatchUpdateOrderRequest)
		Result(model.BatchUpdateOrderResponse)
		HTTP(func() {
			PUT("/outbound-orders")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("create_outbound_order_item", func() {
		Payload(model.OutboundOrderItemCreateRequest)
		Result(model.BaseResponse)
		HTTP(func() {
			POST("/outbound-order-items")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_outbound_order_item", func() {
		Payload(model.OutboundOrderItemUpdateRequest)
		Result(model.BaseResponse)
		HTTP(func() {
			PUT("/outbound-order-items/{id}")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("delete_outbound_order_item", func() {
		Payload(model.DeleteOutboundItemRequest)
		Result(model.BaseResponse)
		HTTP(func() {
			DELETE("/outbound-order-items/{id}")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("batch_query_outbound_order", func() {
		Payload(model.GetOrder)
		Result(model.OrderRsp)
		HTTP(func() {
			GET("/outbound-orders")
			Header("Authorization:Authorization")
			Params(func() {
				Param("order_numbers:order_numbers")
				Param("status:status")
				Param("current:current")
				Param("page_size:page_size")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_outbound_order", func() {
		Payload(model.QueryOutOrder)
		Result(model.QueryOrderRsp)
		HTTP(func() {
			GET("/outbound-orders/{id}")
			Header("Authorization:Authorization")
			Params(func() {
				Param("id:id")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_outbound_order_list_filters", func() {
		Payload(model.AuthToken)
		Result(model.OrderListFiltersResult)
		HTTP(func() {
			GET("/outbound-orders/filters")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_outbound_order_count", func() {
		Payload(model.OrderQueryPayload)
		Result(model.OrderCountResult)
		HTTP(func() {
			GET("/outbound-orders/counts")
			Header("Authorization:Authorization")
			Params(model.OrderQueryParams)
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_outbound_order_list", func() {
		Payload(model.OrderQueryPayload)
		Result(model.GetOrderListResult)
		HTTP(func() {
			GET("/outbound-orders/list")
			Header("Authorization:Authorization")
			Params(model.OrderQueryParams)
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("upload_outbound_orders", func() {
		Payload(model.UploadOrdersPayload)
		Result(model.UploadOrdersResult)
		HTTP(func() {
			POST("/outbound-orders/upload")
			Header("Authorization:Authorization")
			MultipartRequest()
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("export_outbound_orders", func() {
		Payload(model.OrderQueryPayload)
		Result(model.ExportOrderResult)

		Error("internal_error", ErrorResult, "Fault while processing download.")

		HTTP(func() {
			GET("/outbound-orders/export")
			Header("Authorization:Authorization")
			Params(model.OrderQueryParams)
			SkipResponseBodyEncodeDecode()
			Response(func() {
				Header("length:Content-Length")
			})
			Response("Unauthorized", StatusUnauthorized)
			Response("internal_error", StatusInternalServerError)
		})
	})
})
