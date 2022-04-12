// Code generated by goa v3.6.2, DO NOT EDIT.
//
// user HTTP client encoders and decoders
//
// Command:
// $ goa gen goa/design -o ./

package client

import (
	"bytes"
	"context"
	user "goa/gen/user"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildUserSignupRequest instantiates a HTTP request object with method and
// path set to call the "user" service "user_signup" endpoint
func (c *Client) BuildUserSignupRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserSignupUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "user_signup", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUserSignupRequest returns an encoder for requests sent to the user
// user_signup server.
func EncodeUserSignupRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.Signup)
		if !ok {
			return goahttp.ErrInvalidType("user", "user_signup", "*user.Signup", v)
		}
		body := NewUserSignupRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "user_signup", err)
		}
		return nil
	}
}

// DecodeUserSignupResponse returns a decoder for responses returned by the
// user user_signup endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUserSignupResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeUserSignupResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UserSignupResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_signup", err)
			}
			err = ValidateUserSignupResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_signup", err)
			}
			res := NewUserSignupSignupRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UserSignupUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_signup", err)
			}
			err = ValidateUserSignupUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_signup", err)
			}
			return nil, NewUserSignupUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "user_signup", resp.StatusCode, string(body))
		}
	}
}

// BuildUserLoginRequest instantiates a HTTP request object with method and
// path set to call the "user" service "user_login" endpoint
func (c *Client) BuildUserLoginRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserLoginUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "user_login", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUserLoginRequest returns an encoder for requests sent to the user
// user_login server.
func EncodeUserLoginRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.Login)
		if !ok {
			return goahttp.ErrInvalidType("user", "user_login", "*user.Login", v)
		}
		body := NewUserLoginRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "user_login", err)
		}
		return nil
	}
}

// DecodeUserLoginResponse returns a decoder for responses returned by the user
// user_login endpoint. restoreBody controls whether the response body should
// be restored after having been read.
// DecodeUserLoginResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeUserLoginResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UserLoginResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_login", err)
			}
			err = ValidateUserLoginResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_login", err)
			}
			res := NewUserLoginSignupRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UserLoginUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_login", err)
			}
			err = ValidateUserLoginUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_login", err)
			}
			return nil, NewUserLoginUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "user_login", resp.StatusCode, string(body))
		}
	}
}

// BuildUserModifyPasswordRequest instantiates a HTTP request object with
// method and path set to call the "user" service "user_modify_password"
// endpoint
func (c *Client) BuildUserModifyPasswordRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserModifyPasswordUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "user_modify_password", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUserModifyPasswordRequest returns an encoder for requests sent to the
// user user_modify_password server.
func EncodeUserModifyPasswordRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.ModifyPassword)
		if !ok {
			return goahttp.ErrInvalidType("user", "user_modify_password", "*user.ModifyPassword", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUserModifyPasswordRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "user_modify_password", err)
		}
		return nil
	}
}

// DecodeUserModifyPasswordResponse returns a decoder for responses returned by
// the user user_modify_password endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeUserModifyPasswordResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeUserModifyPasswordResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UserModifyPasswordResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_modify_password", err)
			}
			err = ValidateUserModifyPasswordResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_modify_password", err)
			}
			res := NewUserModifyPasswordUserRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UserModifyPasswordUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_modify_password", err)
			}
			err = ValidateUserModifyPasswordUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_modify_password", err)
			}
			return nil, NewUserModifyPasswordUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "user_modify_password", resp.StatusCode, string(body))
		}
	}
}

// BuildUserForgetPasswordRequest instantiates a HTTP request object with
// method and path set to call the "user" service "user_forget_password"
// endpoint
func (c *Client) BuildUserForgetPasswordRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserForgetPasswordUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "user_forget_password", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUserForgetPasswordRequest returns an encoder for requests sent to the
// user user_forget_password server.
func EncodeUserForgetPasswordRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.ForgetPassword)
		if !ok {
			return goahttp.ErrInvalidType("user", "user_forget_password", "*user.ForgetPassword", v)
		}
		body := NewUserForgetPasswordRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "user_forget_password", err)
		}
		return nil
	}
}

// DecodeUserForgetPasswordResponse returns a decoder for responses returned by
// the user user_forget_password endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeUserForgetPasswordResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeUserForgetPasswordResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UserForgetPasswordResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_forget_password", err)
			}
			err = ValidateUserForgetPasswordResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_forget_password", err)
			}
			res := NewUserForgetPasswordUserRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UserForgetPasswordUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_forget_password", err)
			}
			err = ValidateUserForgetPasswordUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_forget_password", err)
			}
			return nil, NewUserForgetPasswordUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "user_forget_password", resp.StatusCode, string(body))
		}
	}
}

// BuildUserValidateRequest instantiates a HTTP request object with method and
// path set to call the "user" service "user_validate" endpoint
func (c *Client) BuildUserValidateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserValidateUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "user_validate", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUserValidateRequest returns an encoder for requests sent to the user
// user_validate server.
func EncodeUserValidateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.ValidateEmail)
		if !ok {
			return goahttp.ErrInvalidType("user", "user_validate", "*user.ValidateEmail", v)
		}
		body := NewUserValidateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "user_validate", err)
		}
		return nil
	}
}

// DecodeUserValidateResponse returns a decoder for responses returned by the
// user user_validate endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUserValidateResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeUserValidateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UserValidateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_validate", err)
			}
			err = ValidateUserValidateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_validate", err)
			}
			res := NewUserValidateUserRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UserValidateUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_validate", err)
			}
			err = ValidateUserValidateUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_validate", err)
			}
			return nil, NewUserValidateUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "user_validate", resp.StatusCode, string(body))
		}
	}
}

// BuildUserLogoutRequest instantiates a HTTP request object with method and
// path set to call the "user" service "user_logout" endpoint
func (c *Client) BuildUserLogoutRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserLogoutUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "user_logout", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUserLogoutRequest returns an encoder for requests sent to the user
// user_logout server.
func EncodeUserLogoutRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.Logout)
		if !ok {
			return goahttp.ErrInvalidType("user", "user_logout", "*user.Logout", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeUserLogoutResponse returns a decoder for responses returned by the
// user user_logout endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUserLogoutResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeUserLogoutResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UserLogoutResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_logout", err)
			}
			err = ValidateUserLogoutResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_logout", err)
			}
			res := NewUserLogoutUserRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UserLogoutUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "user_logout", err)
			}
			err = ValidateUserLogoutUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "user_logout", err)
			}
			return nil, NewUserLogoutUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "user_logout", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserInfoRequest instantiates a HTTP request object with method and
// path set to call the "user" service "get_user_info" endpoint
func (c *Client) BuildGetUserInfoRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserInfoUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "get_user_info", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetUserInfoRequest returns an encoder for requests sent to the user
// get_user_info server.
func EncodeGetUserInfoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.AuthToken)
		if !ok {
			return goahttp.ErrInvalidType("user", "get_user_info", "*user.AuthToken", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeGetUserInfoResponse returns a decoder for responses returned by the
// user get_user_info endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeGetUserInfoResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeGetUserInfoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetUserInfoResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "get_user_info", err)
			}
			err = ValidateGetUserInfoResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "get_user_info", err)
			}
			res := NewGetUserInfoUserInfoRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body GetUserInfoUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "get_user_info", err)
			}
			err = ValidateGetUserInfoUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "get_user_info", err)
			}
			return nil, NewGetUserInfoUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "get_user_info", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateUserInfoRequest instantiates a HTTP request object with method
// and path set to call the "user" service "update_user_info" endpoint
func (c *Client) BuildUpdateUserInfoRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateUserInfoUserPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "update_user_info", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateUserInfoRequest returns an encoder for requests sent to the user
// update_user_info server.
func EncodeUpdateUserInfoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.UserInfo)
		if !ok {
			return goahttp.ErrInvalidType("user", "update_user_info", "*user.UserInfo", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateUserInfoRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "update_user_info", err)
		}
		return nil
	}
}

// DecodeUpdateUserInfoResponse returns a decoder for responses returned by the
// user update_user_info endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeUpdateUserInfoResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeUpdateUserInfoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateUserInfoResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "update_user_info", err)
			}
			err = ValidateUpdateUserInfoResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "update_user_info", err)
			}
			res := NewUpdateUserInfoUserRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UpdateUserInfoUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "update_user_info", err)
			}
			err = ValidateUpdateUserInfoUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "update_user_info", err)
			}
			return nil, NewUpdateUserInfoUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "update_user_info", resp.StatusCode, string(body))
		}
	}
}

// BuildPermissionsRequest instantiates a HTTP request object with method and
// path set to call the "user" service "permissions" endpoint
func (c *Client) BuildPermissionsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PermissionsUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "permissions", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePermissionsRequest returns an encoder for requests sent to the user
// permissions server.
func EncodePermissionsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.AuthToken)
		if !ok {
			return goahttp.ErrInvalidType("user", "permissions", "*user.AuthToken", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodePermissionsResponse returns a decoder for responses returned by the
// user permissions endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodePermissionsResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodePermissionsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body PermissionsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "permissions", err)
			}
			err = ValidatePermissionsResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "permissions", err)
			}
			res := NewPermissionsRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body PermissionsUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "permissions", err)
			}
			err = ValidatePermissionsUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "permissions", err)
			}
			return nil, NewPermissionsUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "permissions", resp.StatusCode, string(body))
		}
	}
}

// unmarshalSignupDataResponseBodyToUserSignupData builds a value of type
// *user.SignupData from a value of type *SignupDataResponseBody.
func unmarshalSignupDataResponseBodyToUserSignupData(v *SignupDataResponseBody) *user.SignupData {
	if v == nil {
		return nil
	}
	res := &user.SignupData{
		Status: *v.Status,
		Token:  *v.Token,
	}

	return res
}

// unmarshalUserDataResponseBodyToUserUserData builds a value of type
// *user.UserData from a value of type *UserDataResponseBody.
func unmarshalUserDataResponseBodyToUserUserData(v *UserDataResponseBody) *user.UserData {
	if v == nil {
		return nil
	}
	res := &user.UserData{
		Status: *v.Status,
	}

	return res
}

// unmarshalUserInfoDataResponseBodyToUserUserInfoData builds a value of type
// *user.UserInfoData from a value of type *UserInfoDataResponseBody.
func unmarshalUserInfoDataResponseBodyToUserUserInfoData(v *UserInfoDataResponseBody) *user.UserInfoData {
	if v == nil {
		return nil
	}
	res := &user.UserInfoData{
		ID:        v.ID,
		UserName:  v.UserName,
		UserEmail: v.UserEmail,
		Avatar:    v.Avatar,
		Phone:     v.Phone,
	}

	return res
}

// unmarshalPermissionsDataResponseBodyToUserPermissionsData builds a value of
// type *user.PermissionsData from a value of type *PermissionsDataResponseBody.
func unmarshalPermissionsDataResponseBodyToUserPermissionsData(v *PermissionsDataResponseBody) *user.PermissionsData {
	if v == nil {
		return nil
	}
	res := &user.PermissionsData{}
	if v.OrderStatus != nil {
		res.OrderStatus = make([]*user.OrderStatusData, len(v.OrderStatus))
		for i, val := range v.OrderStatus {
			res.OrderStatus[i] = unmarshalOrderStatusDataResponseBodyToUserOrderStatusData(val)
		}
	}

	return res
}

// unmarshalOrderStatusDataResponseBodyToUserOrderStatusData builds a value of
// type *user.OrderStatusData from a value of type *OrderStatusDataResponseBody.
func unmarshalOrderStatusDataResponseBodyToUserOrderStatusData(v *OrderStatusDataResponseBody) *user.OrderStatusData {
	if v == nil {
		return nil
	}
	res := &user.OrderStatusData{
		Status: *v.Status,
	}
	res.Action = make([]string, len(v.Action))
	for i, val := range v.Action {
		res.Action[i] = val
	}

	return res
}