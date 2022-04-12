// Code generated by goa v3.6.2, DO NOT EDIT.
//
// HTTP request path constructors for the user service.
//
// Command:
// $ goa gen goa/design -o ./

package server

// UserSignupUserPath returns the URL path to the user service user_signup HTTP endpoint.
func UserSignupUserPath() string {
	return "/v1/users/signup"
}

// UserLoginUserPath returns the URL path to the user service user_login HTTP endpoint.
func UserLoginUserPath() string {
	return "/v1/users/login"
}

// UserModifyPasswordUserPath returns the URL path to the user service user_modify_password HTTP endpoint.
func UserModifyPasswordUserPath() string {
	return "/v1/users/modify-password"
}

// UserForgetPasswordUserPath returns the URL path to the user service user_forget_password HTTP endpoint.
func UserForgetPasswordUserPath() string {
	return "/v1/users/forget-password"
}

// UserValidateUserPath returns the URL path to the user service user_validate HTTP endpoint.
func UserValidateUserPath() string {
	return "/v1/users/validate-email"
}

// UserLogoutUserPath returns the URL path to the user service user_logout HTTP endpoint.
func UserLogoutUserPath() string {
	return "/v1/users/logout"
}

// GetUserInfoUserPath returns the URL path to the user service get_user_info HTTP endpoint.
func GetUserInfoUserPath() string {
	return "/v1/users"
}

// UpdateUserInfoUserPath returns the URL path to the user service update_user_info HTTP endpoint.
func UpdateUserInfoUserPath() string {
	return "/v1/users"
}

// PermissionsUserPath returns the URL path to the user service permissions HTTP endpoint.
func PermissionsUserPath() string {
	return "/v1/users/permissions"
}
