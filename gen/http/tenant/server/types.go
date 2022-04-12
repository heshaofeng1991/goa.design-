// Code generated by goa v3.6.2, DO NOT EDIT.
//
// tenant HTTP server types
//
// Command:
// $ goa gen goa/design -o ./

package server

import (
	tenant "goa/gen/tenant"

	goa "goa.design/goa/v3/pkg"
)

// UpdateTenantInfoRequestBody is the type of the "tenant" service
// "update_tenant_info" endpoint HTTP request body.
type UpdateTenantInfoRequestBody struct {
	// shipping_option
	ShippingOption *int32 `form:"shipping_option,omitempty" json:"shipping_option,omitempty" xml:"shipping_option,omitempty"`
	// default warehouse
	DefaultWarehouse *int32 `form:"default_warehouse,omitempty" json:"default_warehouse,omitempty" xml:"default_warehouse,omitempty"`
	// country code of tariff number
	CountryCode *string `form:"country_code,omitempty" json:"country_code,omitempty" xml:"country_code,omitempty"`
	// us tariff number
	UsTariffNumber *string `form:"us_tariff_number,omitempty" json:"us_tariff_number,omitempty" xml:"us_tariff_number,omitempty"`
	// uk tariff number
	UkTariffNumber *string `form:"uk_tariff_number,omitempty" json:"uk_tariff_number,omitempty" xml:"uk_tariff_number,omitempty"`
	// prepay tariff
	PrepayTariff *bool `form:"prepay_tariff,omitempty" json:"prepay_tariff,omitempty" xml:"prepay_tariff,omitempty"`
}

// IntegrationsResponseBody is the type of the "tenant" service "integrations"
// endpoint HTTP response body.
type IntegrationsResponseBody struct {
	// List of integrations
	Data []*IntegrationDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code int `form:"code" json:"code" xml:"code"`
	// message
	Message string `form:"message" json:"message" xml:"message"`
	// Authorization
	Authorization *string `form:"Authorization,omitempty" json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// JWT used for authentication
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// GetTenantInfoResponseBody is the type of the "tenant" service
// "get_tenant_info" endpoint HTTP response body.
type GetTenantInfoResponseBody struct {
	// data
	Data *TenantDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code int `form:"code" json:"code" xml:"code"`
	// message
	Message string `form:"message" json:"message" xml:"message"`
}

// UpdateTenantInfoResponseBody is the type of the "tenant" service
// "update_tenant_info" endpoint HTTP response body.
type UpdateTenantInfoResponseBody struct {
	// data
	Data *UserDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code int `form:"code" json:"code" xml:"code"`
	// message
	Message string `form:"message" json:"message" xml:"message"`
}

// IntegrationsUnauthorizedResponseBody is the type of the "tenant" service
// "integrations" endpoint HTTP response body for the "Unauthorized" error.
type IntegrationsUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetTenantInfoUnauthorizedResponseBody is the type of the "tenant" service
// "get_tenant_info" endpoint HTTP response body for the "Unauthorized" error.
type GetTenantInfoUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateTenantInfoUnauthorizedResponseBody is the type of the "tenant" service
// "update_tenant_info" endpoint HTTP response body for the "Unauthorized"
// error.
type UpdateTenantInfoUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// IntegrationDataResponseBody is used to define fields on response body types.
type IntegrationDataResponseBody struct {
	// NSS store
	Store *string `form:"store,omitempty" json:"store,omitempty" xml:"store,omitempty"`
	// shopify
	Platform *string `form:"platform,omitempty" json:"platform,omitempty" xml:"platform,omitempty"`
	// created_at
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// updated_at
	IntegrationAt *string `form:"integration_at,omitempty" json:"integration_at,omitempty" xml:"integration_at,omitempty"`
}

// TenantDataResponseBody is used to define fields on response body types.
type TenantDataResponseBody struct {
	// tenant id
	TenantID *int32 `form:"tenant_id,omitempty" json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// tenant code
	TenantCode *string `form:"tenant_code,omitempty" json:"tenant_code,omitempty" xml:"tenant_code,omitempty"`
	// currency
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// balance
	Balance *float64 `form:"balance,omitempty" json:"balance,omitempty" xml:"balance,omitempty"`
	// handling_fee
	HandlingFee *float64 `form:"handling_fee,omitempty" json:"handling_fee,omitempty" xml:"handling_fee,omitempty"`
	// preset_channel_ids
	PresetChannelID []int32 `form:"preset_channel_id,omitempty" json:"preset_channel_id,omitempty" xml:"preset_channel_id,omitempty"`
	// test_channel_id
	TestChannelID []int32 `form:"test_channel_id,omitempty" json:"test_channel_id,omitempty" xml:"test_channel_id,omitempty"`
	// first_inbound_at
	FirstInboundAt *string `form:"first_inbound_at,omitempty" json:"first_inbound_at,omitempty" xml:"first_inbound_at,omitempty"`
	// storage_unit_price
	StorageUnitPrice *float64 `form:"storage_unit_price,omitempty" json:"storage_unit_price,omitempty" xml:"storage_unit_price,omitempty"`
	// shipping_option
	ShippingOption *int32 `form:"shipping_option,omitempty" json:"shipping_option,omitempty" xml:"shipping_option,omitempty"`
	// default warehouse
	DefaultWarehouse *int32 `form:"default_warehouse,omitempty" json:"default_warehouse,omitempty" xml:"default_warehouse,omitempty"`
	// country code of tariff number
	CountryCode *string `form:"country_code,omitempty" json:"country_code,omitempty" xml:"country_code,omitempty"`
	// us tariff number
	UsTariffNumber *string `form:"us_tariff_number,omitempty" json:"us_tariff_number,omitempty" xml:"us_tariff_number,omitempty"`
	// uk tariff number
	UkTariffNumber *string `form:"uk_tariff_number,omitempty" json:"uk_tariff_number,omitempty" xml:"uk_tariff_number,omitempty"`
	// prepay tariff
	PrepayTariff *bool `form:"prepay_tariff,omitempty" json:"prepay_tariff,omitempty" xml:"prepay_tariff,omitempty"`
}

// UserDataResponseBody is used to define fields on response body types.
type UserDataResponseBody struct {
	// status
	Status int `form:"status" json:"status" xml:"status"`
}

// NewIntegrationsResponseBody builds the HTTP response body from the result of
// the "integrations" endpoint of the "tenant" service.
func NewIntegrationsResponseBody(res *tenant.TenantIntegrations) *IntegrationsResponseBody {
	body := &IntegrationsResponseBody{
		Code:          res.Code,
		Message:       res.Message,
		Authorization: res.Authorization,
		Token:         res.Token,
	}
	if res.Data != nil {
		body.Data = make([]*IntegrationDataResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalTenantIntegrationDataToIntegrationDataResponseBody(val)
		}
	}
	return body
}

// NewGetTenantInfoResponseBody builds the HTTP response body from the result
// of the "get_tenant_info" endpoint of the "tenant" service.
func NewGetTenantInfoResponseBody(res *tenant.TenantRsp) *GetTenantInfoResponseBody {
	body := &GetTenantInfoResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Data != nil {
		body.Data = marshalTenantTenantDataToTenantDataResponseBody(res.Data)
	}
	return body
}

// NewUpdateTenantInfoResponseBody builds the HTTP response body from the
// result of the "update_tenant_info" endpoint of the "tenant" service.
func NewUpdateTenantInfoResponseBody(res *tenant.UserRsp) *UpdateTenantInfoResponseBody {
	body := &UpdateTenantInfoResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Data != nil {
		body.Data = marshalTenantUserDataToUserDataResponseBody(res.Data)
	}
	return body
}

// NewIntegrationsUnauthorizedResponseBody builds the HTTP response body from
// the result of the "integrations" endpoint of the "tenant" service.
func NewIntegrationsUnauthorizedResponseBody(res *goa.ServiceError) *IntegrationsUnauthorizedResponseBody {
	body := &IntegrationsUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetTenantInfoUnauthorizedResponseBody builds the HTTP response body from
// the result of the "get_tenant_info" endpoint of the "tenant" service.
func NewGetTenantInfoUnauthorizedResponseBody(res *goa.ServiceError) *GetTenantInfoUnauthorizedResponseBody {
	body := &GetTenantInfoUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateTenantInfoUnauthorizedResponseBody builds the HTTP response body
// from the result of the "update_tenant_info" endpoint of the "tenant" service.
func NewUpdateTenantInfoUnauthorizedResponseBody(res *goa.ServiceError) *UpdateTenantInfoUnauthorizedResponseBody {
	body := &UpdateTenantInfoUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewIntegrationsAuthToken builds a tenant service integrations endpoint
// payload.
func NewIntegrationsAuthToken(authorization *string, token *string) *tenant.AuthToken {
	v := &tenant.AuthToken{}
	v.Authorization = authorization
	v.Token = token

	return v
}

// NewGetTenantInfoAuthToken builds a tenant service get_tenant_info endpoint
// payload.
func NewGetTenantInfoAuthToken(authorization *string, token *string) *tenant.AuthToken {
	v := &tenant.AuthToken{}
	v.Authorization = authorization
	v.Token = token

	return v
}

// NewUpdateTenantInfoTenantInfo builds a tenant service update_tenant_info
// endpoint payload.
func NewUpdateTenantInfoTenantInfo(body *UpdateTenantInfoRequestBody, authorization *string, token *string) *tenant.TenantInfo {
	v := &tenant.TenantInfo{
		ShippingOption:   body.ShippingOption,
		DefaultWarehouse: body.DefaultWarehouse,
		CountryCode:      body.CountryCode,
		UsTariffNumber:   body.UsTariffNumber,
		UkTariffNumber:   body.UkTariffNumber,
		PrepayTariff:     body.PrepayTariff,
	}
	v.Authorization = authorization
	v.Token = token

	return v
}
