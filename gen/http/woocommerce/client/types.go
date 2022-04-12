// Code generated by goa v3.6.2, DO NOT EDIT.
//
// woocommerce HTTP client types
//
// Command:
// $ goa gen goa/design -o ./

package client

import (
	woocommerce "goa/gen/woocommerce"

	goa "goa.design/goa/v3/pkg"
)

// ReturnWoocommerceRequestBody is the type of the "woocommerce" service
// "return_woocommerce" endpoint HTTP request body.
type ReturnWoocommerceRequestBody struct {
	// User ID
	UserID *string `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// Success
	Success *string `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// CallbackWoocommerceRequestBody is the type of the "woocommerce" service
// "callback_woocommerce" endpoint HTTP request body.
type CallbackWoocommerceRequestBody struct {
	// key_id
	KeyID int `form:"key_id" json:"key_id" xml:"key_id"`
	// user_id
	UserID string `form:"user_id" json:"user_id" xml:"user_id"`
	// consumer_key
	ConsumerKey string `form:"consumer_key" json:"consumer_key" xml:"consumer_key"`
	// consumer_secret
	ConsumerSecret string `form:"consumer_secret" json:"consumer_secret" xml:"consumer_secret"`
	// key_permissions
	KeyPermissions string `form:"key_permissions" json:"key_permissions" xml:"key_permissions"`
}

// ReturnWoocommerceResponseBody is the type of the "woocommerce" service
// "return_woocommerce" endpoint HTTP response body.
type ReturnWoocommerceResponseBody struct {
	// Success
	Success *string `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// CallbackWoocommerceResponseBody is the type of the "woocommerce" service
// "callback_woocommerce" endpoint HTTP response body.
type CallbackWoocommerceResponseBody struct {
	// success
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// RetrieveOrdersResponseBody is the type of the "woocommerce" service
// "retrieve_orders" endpoint HTTP response body.
type RetrieveOrdersResponseBody struct {
	// data
	Data interface{} `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code *int `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RetrieveProductsResponseBody is the type of the "woocommerce" service
// "retrieve_products" endpoint HTTP response body.
type RetrieveProductsResponseBody struct {
	// data
	Data interface{} `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code *int `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReturnWoocommerceUnauthorizedResponseBody is the type of the "woocommerce"
// service "return_woocommerce" endpoint HTTP response body for the
// "Unauthorized" error.
type ReturnWoocommerceUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CallbackWoocommerceUnauthorizedResponseBody is the type of the "woocommerce"
// service "callback_woocommerce" endpoint HTTP response body for the
// "Unauthorized" error.
type CallbackWoocommerceUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// RetrieveOrdersUnauthorizedResponseBody is the type of the "woocommerce"
// service "retrieve_orders" endpoint HTTP response body for the "Unauthorized"
// error.
type RetrieveOrdersUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// RetrieveProductsUnauthorizedResponseBody is the type of the "woocommerce"
// service "retrieve_products" endpoint HTTP response body for the
// "Unauthorized" error.
type RetrieveProductsUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewReturnWoocommerceRequestBody builds the HTTP request body from the
// payload of the "return_woocommerce" endpoint of the "woocommerce" service.
func NewReturnWoocommerceRequestBody(p *woocommerce.WoocommerceReturnArgs) *ReturnWoocommerceRequestBody {
	body := &ReturnWoocommerceRequestBody{
		UserID:  p.UserID,
		Success: p.Success,
	}
	return body
}

// NewCallbackWoocommerceRequestBody builds the HTTP request body from the
// payload of the "callback_woocommerce" endpoint of the "woocommerce" service.
func NewCallbackWoocommerceRequestBody(p *woocommerce.WoocommerceCallbackArgs) *CallbackWoocommerceRequestBody {
	body := &CallbackWoocommerceRequestBody{
		KeyID:          p.KeyID,
		UserID:         p.UserID,
		ConsumerKey:    p.ConsumerKey,
		ConsumerSecret: p.ConsumerSecret,
		KeyPermissions: p.KeyPermissions,
	}
	return body
}

// NewReturnWoocommerceWoocommerceReturnResultOK builds a "woocommerce" service
// "return_woocommerce" endpoint result from a HTTP "OK" response.
func NewReturnWoocommerceWoocommerceReturnResultOK(body *ReturnWoocommerceResponseBody) *woocommerce.WoocommerceReturnResult {
	v := &woocommerce.WoocommerceReturnResult{
		Success: body.Success,
	}

	return v
}

// NewReturnWoocommerceUnauthorized builds a woocommerce service
// return_woocommerce endpoint Unauthorized error.
func NewReturnWoocommerceUnauthorized(body *ReturnWoocommerceUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCallbackWoocommerceWoocommerceCallbackRspOK builds a "woocommerce"
// service "callback_woocommerce" endpoint result from a HTTP "OK" response.
func NewCallbackWoocommerceWoocommerceCallbackRspOK(body *CallbackWoocommerceResponseBody) *woocommerce.WoocommerceCallbackRsp {
	v := &woocommerce.WoocommerceCallbackRsp{
		Success: *body.Success,
	}

	return v
}

// NewCallbackWoocommerceUnauthorized builds a woocommerce service
// callback_woocommerce endpoint Unauthorized error.
func NewCallbackWoocommerceUnauthorized(body *CallbackWoocommerceUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewRetrieveOrdersWoocommerceOK builds a "woocommerce" service
// "retrieve_orders" endpoint result from a HTTP "OK" response.
func NewRetrieveOrdersWoocommerceOK(body *RetrieveOrdersResponseBody) *woocommerce.Woocommerce {
	v := &woocommerce.Woocommerce{
		Data:    body.Data,
		Code:    *body.Code,
		Message: *body.Message,
	}

	return v
}

// NewRetrieveOrdersUnauthorized builds a woocommerce service retrieve_orders
// endpoint Unauthorized error.
func NewRetrieveOrdersUnauthorized(body *RetrieveOrdersUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewRetrieveProductsWoocommerceOK builds a "woocommerce" service
// "retrieve_products" endpoint result from a HTTP "OK" response.
func NewRetrieveProductsWoocommerceOK(body *RetrieveProductsResponseBody) *woocommerce.Woocommerce {
	v := &woocommerce.Woocommerce{
		Data:    body.Data,
		Code:    *body.Code,
		Message: *body.Message,
	}

	return v
}

// NewRetrieveProductsUnauthorized builds a woocommerce service
// retrieve_products endpoint Unauthorized error.
func NewRetrieveProductsUnauthorized(body *RetrieveProductsUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateCallbackWoocommerceResponseBody runs the validations defined on
// callback_woocommerce_response_body
func ValidateCallbackWoocommerceResponseBody(body *CallbackWoocommerceResponseBody) (err error) {
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateRetrieveOrdersResponseBody runs the validations defined on
// retrieve_orders_response_body
func ValidateRetrieveOrdersResponseBody(body *RetrieveOrdersResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRetrieveProductsResponseBody runs the validations defined on
// retrieve_products_response_body
func ValidateRetrieveProductsResponseBody(body *RetrieveProductsResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateReturnWoocommerceUnauthorizedResponseBody runs the validations
// defined on return_woocommerce_Unauthorized_response_body
func ValidateReturnWoocommerceUnauthorizedResponseBody(body *ReturnWoocommerceUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCallbackWoocommerceUnauthorizedResponseBody runs the validations
// defined on callback_woocommerce_Unauthorized_response_body
func ValidateCallbackWoocommerceUnauthorizedResponseBody(body *CallbackWoocommerceUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateRetrieveOrdersUnauthorizedResponseBody runs the validations defined
// on retrieve_orders_Unauthorized_response_body
func ValidateRetrieveOrdersUnauthorizedResponseBody(body *RetrieveOrdersUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateRetrieveProductsUnauthorizedResponseBody runs the validations
// defined on retrieve_products_Unauthorized_response_body
func ValidateRetrieveProductsUnauthorizedResponseBody(body *RetrieveProductsUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
