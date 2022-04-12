/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    upload
	@Date    2022/2/10 10:58 上午:00
	@Desc
*/

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

var UploadFile = Type("UploadFile", func() {
	Description("File describes the upload image file")
	Field(1, "file", Bytes, "file", func() {
		Example("1.jpg")
		MaxLength(864000)
	})
	Field(4, "file_name", String, "file name", func() {
		Example("xxx")
		MaxLength(50)
	})
	Extend(AuthToken)
	Required("file", "file_name")
})

var UploadUrl = Type("UploadUrl", func() {
	Description("UploadUrl describes the image url")
	Extend(BaseResponse)
	Field(1, "data", UploadUrlData, "data")
})

var UploadUrlData = Type("UploadUrlData", func() {
	Description("UploadUrlData describes the file url")
	Field(1, "url", String, "file URL", func() {
		Example("http://images.example")
		MaxLength(300)
	})

	Required("url")
})
