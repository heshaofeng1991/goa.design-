// Code generated by goa v3.6.2, DO NOT EDIT.
//
// tenant HTTP server encoders and decoders
//
// Command:
// $ goa gen goa/design -o ./

package server

import (
	"context"
	"errors"
	tenant "goa/gen/tenant"
	"io"
	"net/http"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeIntegrationsResponse returns an encoder for responses returned by the
// tenant integrations endpoint.
func EncodeIntegrationsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*tenant.TenantIntegrations)
		enc := encoder(ctx, w)
		body := NewIntegrationsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeIntegrationsRequest returns a decoder for requests sent to the tenant
// integrations endpoint.
func DecodeIntegrationsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			authorization *string
			token         *string
		)
		authorizationRaw := r.Header.Get("Authorization")
		if authorizationRaw != "" {
			authorization = &authorizationRaw
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewIntegrationsAuthToken(authorization, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeIntegrationsError returns an encoder for errors returned by the
// integrations tenant endpoint.
func EncodeIntegrationsError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "Unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewIntegrationsUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetTenantInfoResponse returns an encoder for responses returned by the
// tenant get_tenant_info endpoint.
func EncodeGetTenantInfoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*tenant.TenantRsp)
		enc := encoder(ctx, w)
		body := NewGetTenantInfoResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetTenantInfoRequest returns a decoder for requests sent to the tenant
// get_tenant_info endpoint.
func DecodeGetTenantInfoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			authorization *string
			token         *string
		)
		authorizationRaw := r.Header.Get("Authorization")
		if authorizationRaw != "" {
			authorization = &authorizationRaw
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewGetTenantInfoAuthToken(authorization, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetTenantInfoError returns an encoder for errors returned by the
// get_tenant_info tenant endpoint.
func EncodeGetTenantInfoError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "Unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetTenantInfoUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateTenantInfoResponse returns an encoder for responses returned by
// the tenant update_tenant_info endpoint.
func EncodeUpdateTenantInfoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*tenant.UserRsp)
		enc := encoder(ctx, w)
		body := NewUpdateTenantInfoResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateTenantInfoRequest returns a decoder for requests sent to the
// tenant update_tenant_info endpoint.
func DecodeUpdateTenantInfoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateTenantInfoRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			authorization *string
			token         *string
		)
		authorizationRaw := r.Header.Get("Authorization")
		if authorizationRaw != "" {
			authorization = &authorizationRaw
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewUpdateTenantInfoTenantInfo(&body, authorization, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeUpdateTenantInfoError returns an encoder for errors returned by the
// update_tenant_info tenant endpoint.
func EncodeUpdateTenantInfoError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "Unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateTenantInfoUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalTenantIntegrationDataToIntegrationDataResponseBody builds a value of
// type *IntegrationDataResponseBody from a value of type
// *tenant.IntegrationData.
func marshalTenantIntegrationDataToIntegrationDataResponseBody(v *tenant.IntegrationData) *IntegrationDataResponseBody {
	if v == nil {
		return nil
	}
	res := &IntegrationDataResponseBody{
		Store:         v.Store,
		Platform:      v.Platform,
		CreatedAt:     v.CreatedAt,
		IntegrationAt: v.IntegrationAt,
	}

	return res
}

// marshalTenantTenantDataToTenantDataResponseBody builds a value of type
// *TenantDataResponseBody from a value of type *tenant.TenantData.
func marshalTenantTenantDataToTenantDataResponseBody(v *tenant.TenantData) *TenantDataResponseBody {
	if v == nil {
		return nil
	}
	res := &TenantDataResponseBody{
		TenantID:         v.TenantID,
		TenantCode:       v.TenantCode,
		Currency:         v.Currency,
		Balance:          v.Balance,
		HandlingFee:      v.HandlingFee,
		FirstInboundAt:   v.FirstInboundAt,
		StorageUnitPrice: v.StorageUnitPrice,
		ShippingOption:   v.ShippingOption,
		DefaultWarehouse: v.DefaultWarehouse,
		CountryCode:      v.CountryCode,
		UsTariffNumber:   v.UsTariffNumber,
		UkTariffNumber:   v.UkTariffNumber,
		PrepayTariff:     v.PrepayTariff,
	}
	if v.PresetChannelID != nil {
		res.PresetChannelID = make([]int32, len(v.PresetChannelID))
		for i, val := range v.PresetChannelID {
			res.PresetChannelID[i] = val
		}
	}
	if v.TestChannelID != nil {
		res.TestChannelID = make([]int32, len(v.TestChannelID))
		for i, val := range v.TestChannelID {
			res.TestChannelID[i] = val
		}
	}

	return res
}

// marshalTenantUserDataToUserDataResponseBody builds a value of type
// *UserDataResponseBody from a value of type *tenant.UserData.
func marshalTenantUserDataToUserDataResponseBody(v *tenant.UserData) *UserDataResponseBody {
	if v == nil {
		return nil
	}
	res := &UserDataResponseBody{
		Status: v.Status,
	}

	return res
}