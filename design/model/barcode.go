/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    barcode
	@Date    2022/2/10 10:55 上午:00
	@Desc
*/

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

var BarCodeRsp = Type("BarCodeRsp", func() {
	Description("BarCodeRsp describes the response of generate barcode")
	Extend(BaseResponse)
	Field(1, "data", BarCodeData, "data")
})

var BarCodeData = Type("BarCodeData", func() {
	Description("BarCodeData describes the response of generate barcode")
	Field(1, "barcode", String, "barcode")

	Required("barcode")
})
