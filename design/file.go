/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    upload
	@Date    2022/1/23 11:22 PM:00
	@Desc
*/

package design

import (
	. "goa.design/goa/v3/dsl" //nolint:revive
	"goa/design/model"
)

var _ = Service("file", func() {
	Description("The file service performs operations on file")

	Error("Unauthorized")

	HTTP(func() {
		Path("/")
	})

	Method("upload_image", func() {
		Payload(model.ImageFile)
		Result(model.ImageUrl)
		HTTP(func() {
			POST("/upload_image")
			MultipartRequest()
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
