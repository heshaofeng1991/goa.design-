// Code generated by goa v3.6.2, DO NOT EDIT.
//
// platform_product HTTP server encoders and decoders
//
// Command:
// $ goa gen goa/design -o ./

package server

import (
	"context"
	"errors"
	platformproduct "goa/gen/platform_product"
	"io"
	"net/http"
	"strconv"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodePlatformProductResponse returns an encoder for responses returned by
// the platform_product platform_product endpoint.
func EncodePlatformProductResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*platformproduct.MultiPlatformProductRsp)
		enc := encoder(ctx, w)
		body := NewPlatformProductResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodePlatformProductRequest returns a decoder for requests sent to the
// platform_product platform_product endpoint.
func DecodePlatformProductRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			platformStatus *int
			name           *string
			pageSize       *int
			current        *int
			id             *int32
			listingSku     *string
			sku            *string
			isMapping      *int
			authorization  *string
			token          *string
			err            error
		)
		{
			platformStatusRaw := r.URL.Query().Get("platform_status")
			if platformStatusRaw != "" {
				v, err2 := strconv.ParseInt(platformStatusRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("platformStatus", platformStatusRaw, "integer"))
				}
				pv := int(v)
				platformStatus = &pv
			}
		}
		nameRaw := r.URL.Query().Get("name")
		if nameRaw != "" {
			name = &nameRaw
		}
		{
			pageSizeRaw := r.URL.Query().Get("page_size")
			if pageSizeRaw != "" {
				v, err2 := strconv.ParseInt(pageSizeRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("pageSize", pageSizeRaw, "integer"))
				}
				pv := int(v)
				pageSize = &pv
			}
		}
		if pageSize != nil {
			if !(*pageSize == 10 || *pageSize == 20 || *pageSize == 50 || *pageSize == 100) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("pageSize", *pageSize, []interface{}{10, 20, 50, 100}))
			}
		}
		{
			currentRaw := r.URL.Query().Get("current")
			if currentRaw != "" {
				v, err2 := strconv.ParseInt(currentRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("current", currentRaw, "integer"))
				}
				pv := int(v)
				current = &pv
			}
		}
		if current != nil {
			if *current < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("current", *current, 1, true))
			}
		}
		{
			idRaw := r.URL.Query().Get("id")
			if idRaw != "" {
				v, err2 := strconv.ParseInt(idRaw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
				}
				pv := int32(v)
				id = &pv
			}
		}
		listingSkuRaw := r.URL.Query().Get("listing_sku")
		if listingSkuRaw != "" {
			listingSku = &listingSkuRaw
		}
		skuRaw := r.URL.Query().Get("sku")
		if skuRaw != "" {
			sku = &skuRaw
		}
		{
			isMappingRaw := r.URL.Query().Get("is_mapping")
			if isMappingRaw != "" {
				v, err2 := strconv.ParseInt(isMappingRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("isMapping", isMappingRaw, "integer"))
				}
				pv := int(v)
				isMapping = &pv
			}
		}
		authorizationRaw := r.Header.Get("Authorization")
		if authorizationRaw != "" {
			authorization = &authorizationRaw
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewPlatformProductGetListing(platformStatus, name, pageSize, current, id, listingSku, sku, isMapping, authorization, token)
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

// EncodePlatformProductError returns an encoder for errors returned by the
// platform_product platform_product endpoint.
func EncodePlatformProductError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewPlatformProductUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeConvertResponse returns an encoder for responses returned by the
// platform_product convert endpoint.
func EncodeConvertResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*platformproduct.ConvertPlatformProductsRes)
		enc := encoder(ctx, w)
		body := NewConvertResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeConvertRequest returns a decoder for requests sent to the
// platform_product convert endpoint.
func DecodeConvertRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body ConvertRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateConvertRequestBody(&body)
		if err != nil {
			return nil, err
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
		payload := NewConvertPlatformProductsReq(&body, authorization, token)
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

// EncodeConvertError returns an encoder for errors returned by the convert
// platform_product endpoint.
func EncodeConvertError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewConvertUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeMappingsResponse returns an encoder for responses returned by the
// platform_product mappings endpoint.
func EncodeMappingsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*platformproduct.UpdateMappingsRes)
		enc := encoder(ctx, w)
		body := NewMappingsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeMappingsRequest returns a decoder for requests sent to the
// platform_product mappings endpoint.
func DecodeMappingsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MappingsRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMappingsRequestBody(&body)
		if err != nil {
			return nil, err
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
		payload := NewMappingsUpdateMappingsReq(&body, authorization, token)
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

// EncodeMappingsError returns an encoder for errors returned by the mappings
// platform_product endpoint.
func EncodeMappingsError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewMappingsUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalPlatformproductMultiPlatformProductToMultiPlatformProductResponseBody
// builds a value of type *MultiPlatformProductResponseBody from a value of
// type *platformproduct.MultiPlatformProduct.
func marshalPlatformproductMultiPlatformProductToMultiPlatformProductResponseBody(v *platformproduct.MultiPlatformProduct) *MultiPlatformProductResponseBody {
	if v == nil {
		return nil
	}
	res := &MultiPlatformProductResponseBody{}
	if v.Meta != nil {
		res.Meta = marshalPlatformproductMetaDataToMetaDataResponseBody(v.Meta)
	}
	if v.List != nil {
		res.List = make([]*ListingResponseBody, len(v.List))
		for i, val := range v.List {
			res.List[i] = marshalPlatformproductListingToListingResponseBody(val)
		}
	}

	return res
}

// marshalPlatformproductMetaDataToMetaDataResponseBody builds a value of type
// *MetaDataResponseBody from a value of type *platformproduct.MetaData.
func marshalPlatformproductMetaDataToMetaDataResponseBody(v *platformproduct.MetaData) *MetaDataResponseBody {
	res := &MetaDataResponseBody{
		Current:  v.Current,
		PageSize: v.PageSize,
		Total:    v.Total,
	}

	return res
}

// marshalPlatformproductListingToListingResponseBody builds a value of type
// *ListingResponseBody from a value of type *platformproduct.Listing.
func marshalPlatformproductListingToListingResponseBody(v *platformproduct.Listing) *ListingResponseBody {
	res := &ListingResponseBody{
		ListingSku:     v.ListingSku,
		Barcode:        v.Barcode,
		Name:           v.Name,
		Vendor:         v.Vendor,
		PlatformStatus: v.PlatformStatus,
		ID:             v.ID,
		IsMapping:      v.IsMapping,
	}
	if v.Images != nil {
		res.Images = make([]string, len(v.Images))
		for i, val := range v.Images {
			res.Images[i] = val
		}
	}
	if v.Store != nil {
		res.Store = marshalPlatformproductStoreToStoreResponseBody(v.Store)
	}
	if v.Mappings != nil {
		res.Mappings = make([]*MappingResponseBody, len(v.Mappings))
		for i, val := range v.Mappings {
			res.Mappings[i] = marshalPlatformproductMappingToMappingResponseBody(val)
		}
	}

	return res
}

// marshalPlatformproductStoreToStoreResponseBody builds a value of type
// *StoreResponseBody from a value of type *platformproduct.Store.
func marshalPlatformproductStoreToStoreResponseBody(v *platformproduct.Store) *StoreResponseBody {
	if v == nil {
		return nil
	}
	res := &StoreResponseBody{
		ID:       v.ID,
		Name:     v.Name,
		Platform: v.Platform,
	}

	return res
}

// marshalPlatformproductMappingToMappingResponseBody builds a value of type
// *MappingResponseBody from a value of type *platformproduct.Mapping.
func marshalPlatformproductMappingToMappingResponseBody(v *platformproduct.Mapping) *MappingResponseBody {
	if v == nil {
		return nil
	}
	res := &MappingResponseBody{
		ID:                v.ID,
		PlatformProductID: v.PlatformProductID,
		ProductID:         v.ProductID,
		ProductSku:        v.ProductSku,
		Qty:               v.Qty,
		ProductName:       v.ProductName,
	}
	if v.Images != nil {
		res.Images = make([]string, len(v.Images))
		for i, val := range v.Images {
			res.Images[i] = val
		}
	}

	return res
}

// marshalPlatformproductConvertPlatformProductsResDataToConvertPlatformProductsResDataResponseBody
// builds a value of type *ConvertPlatformProductsResDataResponseBody from a
// value of type *platformproduct.ConvertPlatformProductsResData.
func marshalPlatformproductConvertPlatformProductsResDataToConvertPlatformProductsResDataResponseBody(v *platformproduct.ConvertPlatformProductsResData) *ConvertPlatformProductsResDataResponseBody {
	if v == nil {
		return nil
	}
	res := &ConvertPlatformProductsResDataResponseBody{
		ID:      v.ID,
		Success: v.Success,
		Message: v.Message,
	}

	return res
}

// unmarshalUpdateMappingsProductRequestBodyToPlatformproductUpdateMappingsProduct
// builds a value of type *platformproduct.UpdateMappingsProduct from a value
// of type *UpdateMappingsProductRequestBody.
func unmarshalUpdateMappingsProductRequestBodyToPlatformproductUpdateMappingsProduct(v *UpdateMappingsProductRequestBody) *platformproduct.UpdateMappingsProduct {
	res := &platformproduct.UpdateMappingsProduct{
		ProductID: *v.ProductID,
		Qty:       *v.Qty,
	}

	return res
}
