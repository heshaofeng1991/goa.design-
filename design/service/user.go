/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    user
	@Date    2022/3/18 3:31 下午
	@Desc
*/

package service

import (
	"goa/design/model"
	. "goa.design/goa/v3/dsl" //nolint:revive
)

var _ = Service("user", func() {
	Description("The user service performs operations on user")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1/users")
	})

	Method("user_signup", func() {
		NoSecurity()
		Payload(model.Signup)
		Result(model.SignupRsp)
		HTTP(func() {
			POST("/signup")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("user_login", func() {
		NoSecurity()
		Payload(model.Login)
		Result(model.SignupRsp)
		HTTP(func() {
			POST("/login")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("user_modify_password", func() {
		Payload(model.ModifyPassword)
		Result(model.UserRsp)
		HTTP(func() {
			POST("/modify-password")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("user_forget_password", func() {
		NoSecurity()
		Payload(model.ForgetPassword)
		Result(model.UserRsp)
		HTTP(func() {
			POST("/forget-password")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("user_validate", func() {
		NoSecurity()
		Payload(model.ValidateEmail)
		Result(model.UserRsp)
		HTTP(func() {
			POST("/validate-email")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("user_logout", func() {
		Payload(model.Logout)
		Result(model.UserRsp)
		HTTP(func() {
			POST("/logout")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_user_info", func() {
		Payload(model.AuthToken)
		Result(model.UserInfoRsp)
		HTTP(func() {
			GET("/")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_user_info", func() {
		Payload(model.UserInfo)
		Result(model.UserRsp)
		HTTP(func() {
			PUT("/")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("permissions", func() {
		Payload(model.AuthToken)
		Result(model.PermissionsRsp)
		HTTP(func() {
			GET("/permissions")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
