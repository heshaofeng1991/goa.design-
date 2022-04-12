/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    upload
	@Date    2022/1/23 11:22 PM:00
	@Desc
*/

package service

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("file", func() {
	Description("The file service performs operations on file")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1")
	})

	Method("upload_image", func() {
		Payload(model.UploadFile)
		Result(model.UploadUrl)
		HTTP(func() {
			POST("/upload_image")
			Header("Authorization:Authorization")
			MultipartRequest()
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
