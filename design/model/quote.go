/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    quote
	@Date    2022/2/10 10:50 上午:00
	@Desc
*/

// Package model API in/out parameter define of shipping-estimate
package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

// GetQuote 运费试算请求入参.
var GetQuote = Type("GetQuote", func() {
	Field(1, "origin_country", String, "Origin country code (two-letter codes)", func() {
		Example("CN")
		MaxLength(2)
		MinLength(2)
	})
	Field(2, "dest_country", String, "Destination country code (two-letter codes)", func() {
		Example("US")
		MaxLength(2)
		MinLength(2)
	})
	Field(3, "dest_state", String, "State of destination country", func() {
		Example("NY")
	})
	Field(4, "dest_zip_code", String, "Destination post code", func() {
		Example("10016")
	})
	Field(5, "weight", Int, "total weight of the package, unit(g)", func() {
		Example(3000)
		Minimum(1)
	})

	Field(6, "length", Int, "the length of the package, unit(mm)", func() {
		Example(100)
		Minimum(1)
	})
	Field(7, "width", Int, "the width of the package, unit(mm)", func() {
		Example(200)
		Minimum(1)
	})
	Field(8, "height", Int, "the height of the package, unit(mm)", func() {
		Example(300)
		Minimum(1)
	})
	Field(9, "product_attributes", ArrayOf(String), "Product attributes", func() {
		Example([]string{"battery", "cosmetic", "liquid", "magnetic"})
	})
	Field(10, "factory", String, "address of factory", func() {
		Example("xxx")
	})
	Field(11, "date", String, "date", func() {
		Example("2022-02-14")
		MaxLength(30)
	})
	Extend(AuthToken)

	Required("origin_country", "dest_country", "dest_state", "dest_zip_code", "weight",
		"length", "width", "height")
})

// Quote 运费试算返回出参.
var Quote = Type("Quote", func() {
	Field(1, "channel_name", String, "Channel Display Name", func() {
		Example("NSS")
	})
	Field(2, "channel_id", Int32, "Channel ID", func() {
		Example(1)
	})
	Field(3, "type", Int, "Channel type", func() {
		Example(1)
	})
	Field(4, "min_normal_days", Int32, "Min Normal Days", func() {
		Example(3)
	})
	Field(5, "max_normal_days", Int32, "Max Normal Days", func() {
		Example(5)
	})
	Field(6, "total_cost", Float64, "Total cost (decimal(15,2))", func() {
		Example(45.89)
	})
	Field(7, "currency", String, func() {
		Example("USD")
	})
	Field(8, "weight", Int, "weight(unit g)", func() {
		Example(1000)
	})

	Required("channel_name", "channel_id", "type", "min_normal_days", "max_normal_days", "total_cost", "currency", "weight")
})

var QuoteInfo = Type("QuoteInfo", func() {
	Field(1, "list", ArrayOf(Quote), "fees")
	Required("list")
})

var QuoteRsp = Type("QuoteRsp", func() {
	Description("QuoteRsp describes the fee info")
	Extend(BaseResponse)
	Field(1, "data", QuoteInfo, "data")
})

var PostQuote = Type("PostQuote", func() {
	Field(1, "id", Int32, "channel cost id")
	Field(4, "delivery_area", Int, "delivery area")
	Field(5, "delivery_country_code", String, "delivery country code")
	Field(6, "delivery_country_name", String, "delivery country name")
	Field(7, "delivery_province_code", String, "delivery province code")
	Field(8, "delivery_province_name", String, "delivery province name")
	Field(9, "delivery_city_code", String, "delivery city code")
	Field(10, "delivery_city_name", String, "delivery city name")
	Field(11, "dest_area", Int, "dest area")
	Field(12, "dest_country_code", String, "dest country code")
	Field(13, "dest_country_name", String, "dest country name")
	Field(14, "dest_province_code", String, "dest province code")
	Field(15, "dest_province_name", String, "dest province name")
	Field(16, "dest_city_code", String, "dest city code")
	Field(17, "dest_city_name", String, "dest city name")
	Extend(AuthToken)

	Required("id", "delivery_area", "delivery_country_code",
		"delivery_country_name", "dest_area",
		"dest_country_code", "dest_country_name")
})
