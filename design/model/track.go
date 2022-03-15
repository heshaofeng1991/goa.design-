/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    track
	@Date    2022/2/10 11:01 上午:00
	@Desc
*/

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

// GetTrack 物流轨迹请求入参.
var GetTrack = Type("GetTrack", func() {
	Field(1, "tracking_number", String, "tracking number of order")
	Field(3, "type", Int, "type(1 tracking_number 2 order_id)", func() {
		Enum(1, 2)
		Example(1)
	})
	Required("tracking_number", "type")
})

// Track 物流轨迹返回出参.
var Track = Type("Track", func() {
	Field(1, "tracking_number", String, "tracking number of order", func() {
		Example("NSS193719374,NSS193719")
	})
	Field(2, "tracking_url", String, "tracking url", func() {
		Example("http://example.com")
	})
	Field(3, "details", ArrayOf(TrackItems), "tracking details")
	Field(4, "status", Int, "status", func() {
		Example(0)
	})
	Field(5, "type", Int, "type", func() {
		Example(1)
	})
	Field(6, "order_id", String, "order_id", func() {
		Example("xxx")
	})
	Required("tracking_number", "details", "status", "type")
})

// TrackItems 物流轨迹消息体内容.
var TrackItems = Type("TrackItem", func() {
	Field(1, "content", String, "tracking description", func() {
		Example("shipping done")
	})
	Field(2, "timestamp", String, "tracking timestamp", func() {
		Example("2022-01-27 00:00:00")
	})
})
