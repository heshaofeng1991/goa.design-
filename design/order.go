// Package design @Author johnny 2022/1/13 1:39 PM:00
package design

import (
	. "goa.design/goa/v3/dsl" //nolint:revive
	"goa/design/model"
)

var _ = Service("order", func() {
	Description("The order service performs operations on order")

	Error("Unauthorized")

	HTTP(func() {
		Path("/")
	})

	Method("create_inbound_order", func() {
		Payload(model.InboundOrder)
		Result(model.InboundOrderRsp)
		HTTP(func() {
			POST("/inbound-orders")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_inbound_order", func() {
		Payload(model.InboundOrder)
		Result(model.UpdateResponse)
		HTTP(func() {
			PUT("/inbound-orders")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("create_outbound_order", func() {
		Payload(model.OutboundOrder)
		Result(model.OutboundOrderRsp)
		HTTP(func() {
			POST("/outbound-orders")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_outbound_order", func() {
		Payload(model.OutboundOrder)
		Result(model.UpdateResponse)
		HTTP(func() {
			PUT("/outbound-orders")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("create_pickup_order", func() {
		Payload(model.PickupOrder)
		Result(model.PickupOrderRsp)
		HTTP(func() {
			POST("/pickup-orders")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_inbound_order", func() {
		Payload(model.GetOrder)
		Result(ArrayOf(model.InboundOrderResponse))
		HTTP(func() {
			GET("/inbound-orders")
			Params(func() {
				Param("client_order_id:client_order_id")
				Required("client_order_id")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_outbound_order", func() {
		Payload(model.GetOrder)
		Result(ArrayOf(model.OrderRsp))
		HTTP(func() {
			GET("/outbound-orders")
			Params(func() {
				Param("client_order_id:client_order_id")
				Required("client_order_id")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
