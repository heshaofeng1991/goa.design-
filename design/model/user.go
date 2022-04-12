/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    user
	@Date    2022/3/18 3:35 下午
	@Desc
*/

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

// Signup 用户注册请求结构.
var Signup = Type("Signup", func() {
	Description("Signup describes user register")
	Field(1, "email", String, "user register email", func() {
		Example("xxx@nextsmartship.com")
	})
	Field(2, "password", String, "user register password", func() {
		Example("xxx")
	})
	Field(3, "user_name", String, "user register name", func() {
		Example("john")
	})
	Field(4, "phone", String, "user register phone", func() {
		Example("139xxxxxxxx")
	})
	Field(5, "source", String, "source form platform", func() {
		Example("shopify")
	})
	Field(6, "inviter_id", Int32, "inviter id", func() {
		Minimum(1)
		Example(1)
	})
	Field(7, "store_code", String, "store code", func() {
		MinLength(1)
		Example("xxx")
	})
	Field(8, "website", String, "website", func() {
		MinLength(1)
		Example("xxx")
	})
	Field(7, "platform", String, "platform", func() {
		MinLength(1)
		Example("xxx")
	})
	Field(7, "concerns", String, "concerns", func() {
		MinLength(1)
		Example("xxx")
	})
	Required("email", "password")
})

// SignupRsp 注册响应结构.
var SignupRsp = Type("SignupRsp", func() {
	Field(1, "data", SignupData, "data")
	Extend(BaseResponse)
})

// SignupData 注册响应data结构.
var SignupData = Type("SignupData", func() {
	Field(1, "status", Int, "status")
	Field(2, "token", String, "token")
	Required("status", "token")
})

// UserRsp 用户公共响应返回结构.
var UserRsp = Type("UserRsp", func() {
	Field(1, "data", UserData, "data")
	Extend(BaseResponse)
})

// UserData 用户公共响应data结构.
var UserData = Type("UserData", func() {
	Field(1, "status", Int, "status")
	Required("status")
})

// Login 用户注册请求结构.
var Login = Type("Login", func() {
	Description("Login describes user login")
	Field(1, "email", String, "user login email", func() {
		Example("xxx@nextsmartship.com")
	})
	Field(2, "password", String, "user login password", func() {
		Example("xxx")
	})
	Field(3, "store_code", String, "store code", func() {
		MinLength(1)
		Example("xxx")
	})
	Required("email", "password")
})

var ModifyPassword = Type("ModifyPassword", func() {
	Description("ModifyPassword describes user modify password")
	Field(1, "email", String, "user email", func() {
		Example("xxx@nextsmartship.com")
	})
	Field(2, "password", String, "user password", func() {
		Example("xxx")
	})
	Field(3, "old_password", String, "old password", func() {
		Example("xxx")
	})
	Extend(AuthToken)
	Required("email", "password", "old_password")
})

var ForgetPassword = Type("ForgetPassword", func() {
	Description("ForgetPassword describes user forget password, need set new password")
	Field(1, "email", String, "user email", func() {
		Example("xxx@nextsmartship.com")
	})
	Field(2, "password", String, "user password", func() {
		Example("xxx")
	})
	Field(3, "code", String, "email code", func() {
		Example("xxx")
	})
	Field(4, "action", String, "action", func() {
		MinLength(1)
		Example("xxx")
	})
	Required("email", "password", "code")
})

var ValidateEmail = Type("ValidateEmail", func() {
	Description("ValidateEmail describes verify email")
	Field(1, "email", String, "user email", func() {
		Example("xxx@nextsmartship.com")
	})
	Required("email")
})

var Logout = Type("Logout", func() {
	Description("Logout describes the user logout")
	Extend(AuthToken)
})

// UserInfoRsp 响应结构.
var UserInfoRsp = Type("UserInfoRsp", func() {
	Field(1, "data", UserInfoData, "data")
	Extend(BaseResponse)
})

// UserInfoData 响应data结构.
var UserInfoData = Type("UserInfoData", func() {
	Field(1, "id", Int32, "id", func() {
		Example(1)
	})
	Field(2, "user_name", String, "user name", func() {
		Example("john")
	})
	Field(3, "user_email", String, "user email", func() {
		Example("xxx@nextsmartship.com")
	})
	Field(4, "avatar", String, "user avatar", func() {
		Example("link url")
	})
	Field(4, "phone", String, "user phone", func() {
		Example("139xxxxxxxx")
	})
})

// UserInfo 更新user info结构.
var UserInfo = Type("UserInfo", func() {
	Field(1, "user_name", String, "user name")
	Extend(AuthToken)
	Required("user_name")
})

var PermissionsRsp = Type("PermissionsRsp", func() {
	Description("Permissions")
	Extend(BaseResponse)
	Field(1, "data", PermissionsData, "data")
})

var PermissionsData = Type("PermissionsData", func() {
	Description("OrderStatusData describes the details")
	Field(1, "order_status", ArrayOf(OrderStatusData), "order_status")
})

var OrderStatusData = Type("OrderStatusData", func() {
	Description("OrderStatusData describes the details")
	Field(1, "status", Int, "new")
	Field(2, "action", ArrayOf(String), "ready")
	Required("status", "action")
})
