package design

import (
	. "goa.design/goa/v3/dsl" //nolint:revive
	"goa/design/model"
)

var _ = Service("track", func() {
	Description("The track service performs operations on track order status")

	Error("Unauthorized")

	HTTP(func() {
		Path("/track")
	})

	Method("get", func() {
		Payload(model.GetTrack)
		Result(ArrayOf(model.Track))
		HTTP(func() {
			GET("/")
			Params(func() {
				Param("tracking_number:tracking_number")
				Param("type:type")
				Required("tracking_number", "type")
			})
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
