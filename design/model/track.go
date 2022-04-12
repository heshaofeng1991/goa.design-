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

// BatchQueryTrackPayload 物流轨迹请求入参.
var BatchQueryTrackPayload = Type("BatchQueryTrackPayload", func() {
	Field(1, "tracking_numbers", ArrayOf(String), "tracking number", func() {
		MinLength(1)
		MaxLength(50)
	})
	Required("tracking_numbers")
})

// QueryTrackPayload 物流轨迹请求入参.
var QueryTrackPayload = Type("QueryTrackPayload", func() {
	Field(1, "tracking_number", String, "tracking number", func() {
		MinLength(1)
	})
	Required("tracking_number")
})

// QueryTrackRsp 物流轨迹返回出参.
var QueryTrackRsp = Type("QueryTrackRsp", func() {
	Extend(BaseResponse)
	Field(1, "data", TrackInfo, "data")
})

var TrackInfo = Type("TrackInfo", func() {
	Field(1, "list", ArrayOf(Track), "tracks")
	Required("list")
})

// Track 物流轨迹返回出参Data.
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
	Required("tracking_number", "tracking_url", "details", "status")
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

// TrackRsp 物流轨迹返回出参.
var TrackRsp = Type("TrackRsp", func() {
	Extend(BaseResponse)
	Field(1, "data", Track, "data")
})
