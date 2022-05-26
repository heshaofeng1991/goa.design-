// Package design API in/out parameter define

package model

import (
	. "goa.design/goa/v3/dsl"
)

var BaseResponse = Type("BaseResponse", func() {
	Field(1, "code", Int, "code", func() {
		Example(0)
	})
	Field(2, "message", String, "message", func() {
		Example("description error information")
	})
	Required("code", "message")
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
