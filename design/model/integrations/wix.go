package model

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl"
)

var WixReturnArgs = Type("WixReturnArgs", func() {
	Description("Wix Return Arguments")
	Field(1, "token", String, "token", func() {
		Example("XXX")
	})
})

var WixReturnResult = Type("WixReturnResult", func() {
	Description("Wix Return Result")
	Field(1, "success", String, "Success", func() {
		Example("0")
	})
})

var WixCallbackArgs = Type("WixCallbackArgs", func() {
	Description("Wix callback args")
	Field(1, "code", String, "code")
	Field(2, "state", String, "state")
	Field(3, "instanceId", String, "instanceId")
	Required("instanceId", "code")
})

var WixCallbackRsp = Type("WixCallbackRsp", func() {
	Description("Wix callback response")
	Field(1, "success", Boolean, "success")
	Required("success")
})

// GetWix 请求入参.
var GetWix = Type("GetWix", func() {
	Field(1, "start_date", String, "start_date")
	Field(2, "end_date", String, "end_date")
	Extend(model.AuthToken)
})

// Wix 返回出参.
var Wix = Type("Wix", func() {
	Description("WixOrders describes an order to outbound order.")
	Field(1, "data", String, "result", func() {
		Example("Success")
	})

	Required("data")
})

var WebHooksCallBackProducts = Type("WebHooksCallBackProducts", func() {
	Field(1, "start_date", String, "start_date")
	Field(2, "end_date", String, "end_date")
	// Extend(model.AuthToken)
})

var WebHooksCallBackProductsResp = Type("WebHooksCallBackProductsResp", func() {
	Description("Wix callback response")
	Field(1, "success", Boolean, "success")
	Required("success")
})

var GetProductsListReq = Type("GetProductsListReq", func() {
	Field(1, "start_date", String, "start_date")
	Field(2, "end_date", String, "end_date")
	// Extend(model.AuthToken)
})

var GetProductsListResp = Type("GetProductsListResp", func() {
	Description("Wix product list response")
	Field(1, "success", Boolean, "success")
	Required("success")
})

var GetOrdersListReq = Type("GetOrdersListReq", func() {
	Field(1, "start_date", String, "start_date")
	Field(2, "end_date", String, "end_date")
	// Extend(model.AuthToken)
})

var GetOrdersListResp = Type("GetOrdersListResp", func() {
	Description("Wix oder list response")
	Field(1, "success", Boolean, "success")
	Required("success")
})
