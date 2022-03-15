// Package design API in/out parameter define

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

// ShippingAddress 地址信息.
var ShippingAddress = Type("ShippingAddress", func() {
	Description("ShippingAddress describes the receiver address")
	Field(1, "first_name", String, "First name", func() {
		Example("He")
	})
	Field(2, "last_name", String, "Last name", func() {
		Example("John")
	})
	Field(3, "phone_number", String, "Phone number", func() {
		Example("150xxxxxxxx")
	})
	Field(4, "country_name", String, "Country Name", func() {
		Example("US")
	})
	Field(5, "country_code", String, "Country code", func() {
		Example("US")
	})
	Field(6, "state_name", String, "State Name", func() {
		Example("U")
	})
	Field(7, "state_code", String, "State code", func() {
		Example("S")
	})
	Field(8, "address1", String, "Address Line 1", func() {
		Example("address1")
	})
	Field(9, "address2", String, "Address Line 2", func() {
		Example("address2")
	})
	Field(10, "city_name", String, "City Name", func() {
		Example("SZ")
	})
	Field(11, "zip_code", String, "ZIP code", func() {
		Example("10016")
	})
	Field(12, "name", String, "name", func() {
		Example("He")
	})

	Required("first_name", "last_name", "phone_number", "country_name",
		"country_code", "state_name", "state_code", "address1", "address2", "city_name", "zip_code")
})

var GetOrder = Type("GetOrder", func() {
	Description("GetOrder describes the order info")
	Field(1, "client_order_id", String, "client_order_id", func() {
		MinLength(1)
		Example("xxx1234")
	})

	Required("client_order_id")
})

var UpdateResponse = Type("UpdateResponse", func() {
	Description("UpdateResponse describes update status")
	Field(1, "status", Int32, "status", func() {
		Enum(0, 1)
		Example(1)
	})

	Required("status")
})
