package service

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("track", func() {
	Description("The track service performs operations on track order status")

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1")
	})

	Method("batch_query_track_info", func() {
		Payload(model.BatchQueryTrackPayload)
		Result(model.QueryTrackRsp)
		HTTP(func() {
			GET("/tracks")
			Params(func() {
				Param("tracking_numbers:tracking_numbers")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_track", func() {
		Payload(model.QueryTrackPayload)
		Result(model.TrackRsp)
		HTTP(func() {
			GET("/tracks/{tracking_number}")
			Params(func() {
				Param("tracking_number:tracking_number")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
