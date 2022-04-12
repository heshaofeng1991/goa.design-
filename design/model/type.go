// Package design API in/out parameter define

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

var BaseResponse = Type("BaseResponse", func() {
	Field(1, "code", Int, "code", func() {
		Example(0)
	})
	Field(2, "message", String, "message", func() {
		Example("description error information")
	})
	Required("code", "message")
})

// Address 地址信息.
var Address = Type("Address", func() {
	Description("Address")
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
	Field(13, "company", String, "company", func() {
		Example("China")
	})
	Field(14, "email", String, "email", func() {
		Example("123@test.com")
	})
	Field(15, "certificate_type", String, "certificate type", func() {
		Enum("ID", "PP")
		Example("ID")
	})
	Field(16, "certificate_code", String, "certificate code", func() {
		Example("3455233")
	})
	Field(17, "certificate_period", String, "certificate period", func() {
		Example("2028-01-01")
	})
})

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
	Field(13, "company", String, "company", func() {
		Example("China")
	})
	Field(14, "email", String, "email", func() {
		Example("123@test.com")
	})
	Field(15, "certificate_type", String, "certificate type", func() {
		Enum("ID", "PP")
		Example("ID")
	})
	Field(16, "certificate_code", String, "certificate code", func() {
		Example("3455233")
	})
	Field(17, "certificate_period", String, "certificate period", func() {
		Example("2028-01-01")
	})

	Required("first_name", "last_name", "phone_number", "country_name",
		"country_code", "state_name", "state_code", "address1", "address2", "city_name", "zip_code")
})

var GetOrder = Type("GetOrder", func() {
	Description("GetOrder describes the order info")
	Field(1, "order_numbers", ArrayOf(String), "order number", func() {
		MaxLength(50)
		Example([]string{"xxx1234"})
	})
	Field(2, "status", Int, "status", func() {
		Minimum(0)
		Maximum(100)
		Example(1)
	})
	Extend(Pagination)
	Extend(AuthToken)
})

var UpdateResponse = Type("UpdateResponse", func() {
	Description("UpdateResponse describes update status")
	Extend(BaseResponse)
	Field(3, "data", UpdateResponseData, "data")
})

var UpdateResponseData = Type("UpdateResponseData", func() {
	Description("UpdateResponse describes update status")
	Field(1, "status", Int32, "status", func() {
		Enum(0, 1)
		Example(1)
	})
	Required("status")
})

var Store = Type("Store", func() {
	Description("Store describes the store")
	Field(1, "id", Int64, "id", func() {
		Minimum(1)
		Example(1)
	})
	Field(2, "name", String, "name", func() {
		Example("xxx")
	})
	Field(3, "platform", String, "platform", func() {
		Example("xxx")
	})

	Required("id", "name", "platform")
})

var MetaData = Type("MetaData", func() {
	Description("MetaData describes the MetaData")
	Field(1, "current", Int, "current", func() {
		Minimum(1)
		Example(1)
	})
	Field(2, "page_size", Int, "page_size", func() {
		Minimum(1)
		Example(1)
	})
	Field(3, "total", Int, "total", func() {
		Example(1)
	})

	Required("current", "page_size", "total")
})

var Pagination = Type("Pagination", func() {
	Field(1, "current", Int, "current", func() {
		Minimum(1)
		Example(1)
	})
	Field(2, "page_size", Int, "page_size", func() {
		Minimum(1)
		Maximum(50)
		Example(1)
	})
})

var AuthToken = Type("AuthToken", func() {
	Field(1, "Authorization", String, "Authorization")
	TokenField(2, "token", String, func() {
		Description("JWT used for authentication")
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
})
