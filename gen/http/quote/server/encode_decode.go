// Code generated by goa v3.6.2, DO NOT EDIT.
//
// quote HTTP server encoders and decoders
//
// Command:
// $ goa gen goa/design -o ./

package server

import (
	"context"
	"errors"
	quote "goa/gen/quote"
	"io"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetResponse returns an encoder for responses returned by the quote get
// endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*quote.QuoteRsp)
		enc := encoder(ctx, w)
		body := NewGetResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRequest returns a decoder for requests sent to the quote get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			originCountry     string
			destCountry       string
			destState         string
			destZipCode       string
			weight            int
			length            int
			width             int
			height            int
			productAttributes []string
			factory           *string
			date              *string
			authorization     *string
			token             *string
			err               error
		)
		originCountry = r.URL.Query().Get("origin_country")
		if originCountry == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("origin_country", "query string"))
		}
		if utf8.RuneCountInString(originCountry) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("originCountry", originCountry, utf8.RuneCountInString(originCountry), 2, true))
		}
		if utf8.RuneCountInString(originCountry) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("originCountry", originCountry, utf8.RuneCountInString(originCountry), 2, false))
		}
		destCountry = r.URL.Query().Get("dest_country")
		if destCountry == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("dest_country", "query string"))
		}
		if utf8.RuneCountInString(destCountry) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("destCountry", destCountry, utf8.RuneCountInString(destCountry), 2, true))
		}
		if utf8.RuneCountInString(destCountry) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("destCountry", destCountry, utf8.RuneCountInString(destCountry), 2, false))
		}
		destState = r.URL.Query().Get("dest_state")
		if destState == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("dest_state", "query string"))
		}
		destZipCode = r.URL.Query().Get("dest_zip_code")
		if destZipCode == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("dest_zip_code", "query string"))
		}
		{
			weightRaw := r.URL.Query().Get("weight")
			if weightRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("weight", "query string"))
			}
			v, err2 := strconv.ParseInt(weightRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("weight", weightRaw, "integer"))
			}
			weight = int(v)
		}
		if weight < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("weight", weight, 1, true))
		}
		{
			lengthRaw := r.URL.Query().Get("length")
			if lengthRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("length", "query string"))
			}
			v, err2 := strconv.ParseInt(lengthRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("length", lengthRaw, "integer"))
			}
			length = int(v)
		}
		if length < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("length", length, 1, true))
		}
		{
			widthRaw := r.URL.Query().Get("width")
			if widthRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("width", "query string"))
			}
			v, err2 := strconv.ParseInt(widthRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("width", widthRaw, "integer"))
			}
			width = int(v)
		}
		if width < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("width", width, 1, true))
		}
		{
			heightRaw := r.URL.Query().Get("height")
			if heightRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("height", "query string"))
			}
			v, err2 := strconv.ParseInt(heightRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("height", heightRaw, "integer"))
			}
			height = int(v)
		}
		if height < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("height", height, 1, true))
		}
		productAttributes = r.URL.Query()["product_attributes"]
		factoryRaw := r.URL.Query().Get("factory")
		if factoryRaw != "" {
			factory = &factoryRaw
		}
		dateRaw := r.URL.Query().Get("date")
		if dateRaw != "" {
			date = &dateRaw
		}
		if date != nil {
			if utf8.RuneCountInString(*date) > 30 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("date", *date, utf8.RuneCountInString(*date), 30, false))
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
		payload := NewGetQuote(originCountry, destCountry, destState, destZipCode, weight, length, width, height, productAttributes, factory, date, authorization, token)
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

// EncodeGetError returns an encoder for errors returned by the get quote
// endpoint.
func EncodeGetError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewGetUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodePostResponse returns an encoder for responses returned by the quote
// post endpoint.
func EncodePostResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*quote.UserRsp)
		enc := encoder(ctx, w)
		body := NewPostResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodePostRequest returns a decoder for requests sent to the quote post
// endpoint.
func DecodePostRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body PostRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidatePostRequestBody(&body)
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
		payload := NewPostQuote(&body, authorization, token)
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

// EncodePostError returns an encoder for errors returned by the post quote
// endpoint.
func EncodePostError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewPostUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalQuoteQuoteInfoToQuoteInfoResponseBody builds a value of type
// *QuoteInfoResponseBody from a value of type *quote.QuoteInfo.
func marshalQuoteQuoteInfoToQuoteInfoResponseBody(v *quote.QuoteInfo) *QuoteInfoResponseBody {
	if v == nil {
		return nil
	}
	res := &QuoteInfoResponseBody{}
	if v.List != nil {
		res.List = make([]*QuoteResponseBody, len(v.List))
		for i, val := range v.List {
			res.List[i] = marshalQuoteQuoteToQuoteResponseBody(val)
		}
	}

	return res
}

// marshalQuoteQuoteToQuoteResponseBody builds a value of type
// *QuoteResponseBody from a value of type *quote.Quote.
func marshalQuoteQuoteToQuoteResponseBody(v *quote.Quote) *QuoteResponseBody {
	res := &QuoteResponseBody{
		ChannelName:   v.ChannelName,
		ChannelID:     v.ChannelID,
		Type:          v.Type,
		MinNormalDays: v.MinNormalDays,
		MaxNormalDays: v.MaxNormalDays,
		TotalCost:     v.TotalCost,
		Currency:      v.Currency,
		Weight:        v.Weight,
	}

	return res
}

// marshalQuoteUserDataToUserDataResponseBody builds a value of type
// *UserDataResponseBody from a value of type *quote.UserData.
func marshalQuoteUserDataToUserDataResponseBody(v *quote.UserData) *UserDataResponseBody {
	if v == nil {
		return nil
	}
	res := &UserDataResponseBody{
		Status: v.Status,
	}

	return res
}
