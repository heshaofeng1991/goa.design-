// Code generated by goa v3.6.2, DO NOT EDIT.
//
// product HTTP server
//
// Command:
// $ goa gen goa/design -o ./

package server

import (
	"context"
	product "goa/gen/product"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the product service endpoint HTTP handlers.
type Server struct {
	Mounts               []*MountPoint
	BatchesCreateProduct http.Handler
	UpdateProduct        http.Handler
	GenerateBarcode      http.Handler
	GenerateToken        http.Handler
	CORS                 http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the product service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *product.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"BatchesCreateProduct", "POST", "/products"},
			{"UpdateProduct", "PUT", "/products"},
			{"GenerateBarcode", "POST", "/barcode"},
			{"GenerateToken", "POST", "/token"},
			{"CORS", "OPTIONS", "/products"},
			{"CORS", "OPTIONS", "/barcode"},
			{"CORS", "OPTIONS", "/token"},
		},
		BatchesCreateProduct: NewBatchesCreateProductHandler(e.BatchesCreateProduct, mux, decoder, encoder, errhandler, formatter),
		UpdateProduct:        NewUpdateProductHandler(e.UpdateProduct, mux, decoder, encoder, errhandler, formatter),
		GenerateBarcode:      NewGenerateBarcodeHandler(e.GenerateBarcode, mux, decoder, encoder, errhandler, formatter),
		GenerateToken:        NewGenerateTokenHandler(e.GenerateToken, mux, decoder, encoder, errhandler, formatter),
		CORS:                 NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "product" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.BatchesCreateProduct = m(s.BatchesCreateProduct)
	s.UpdateProduct = m(s.UpdateProduct)
	s.GenerateBarcode = m(s.GenerateBarcode)
	s.GenerateToken = m(s.GenerateToken)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the product endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountBatchesCreateProductHandler(mux, h.BatchesCreateProduct)
	MountUpdateProductHandler(mux, h.UpdateProduct)
	MountGenerateBarcodeHandler(mux, h.GenerateBarcode)
	MountGenerateTokenHandler(mux, h.GenerateToken)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the product endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountBatchesCreateProductHandler configures the mux to serve the "product"
// service "batches_create_product" endpoint.
func MountBatchesCreateProductHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleProductOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/products", f)
}

// NewBatchesCreateProductHandler creates a HTTP handler which loads the HTTP
// request and calls the "product" service "batches_create_product" endpoint.
func NewBatchesCreateProductHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeBatchesCreateProductRequest(mux, decoder)
		encodeResponse = EncodeBatchesCreateProductResponse(encoder)
		encodeError    = EncodeBatchesCreateProductError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "batches_create_product")
		ctx = context.WithValue(ctx, goa.ServiceKey, "product")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateProductHandler configures the mux to serve the "product" service
// "update_product" endpoint.
func MountUpdateProductHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleProductOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/products", f)
}

// NewUpdateProductHandler creates a HTTP handler which loads the HTTP request
// and calls the "product" service "update_product" endpoint.
func NewUpdateProductHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateProductRequest(mux, decoder)
		encodeResponse = EncodeUpdateProductResponse(encoder)
		encodeError    = EncodeUpdateProductError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update_product")
		ctx = context.WithValue(ctx, goa.ServiceKey, "product")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGenerateBarcodeHandler configures the mux to serve the "product"
// service "generate_barcode" endpoint.
func MountGenerateBarcodeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleProductOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/barcode", f)
}

// NewGenerateBarcodeHandler creates a HTTP handler which loads the HTTP
// request and calls the "product" service "generate_barcode" endpoint.
func NewGenerateBarcodeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGenerateBarcodeRequest(mux, decoder)
		encodeResponse = EncodeGenerateBarcodeResponse(encoder)
		encodeError    = EncodeGenerateBarcodeError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "generate_barcode")
		ctx = context.WithValue(ctx, goa.ServiceKey, "product")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGenerateTokenHandler configures the mux to serve the "product" service
// "generate_token" endpoint.
func MountGenerateTokenHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleProductOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/token", f)
}

// NewGenerateTokenHandler creates a HTTP handler which loads the HTTP request
// and calls the "product" service "generate_token" endpoint.
func NewGenerateTokenHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGenerateTokenRequest(mux, decoder)
		encodeResponse = EncodeGenerateTokenResponse(encoder)
		encodeError    = EncodeGenerateTokenError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "generate_token")
		ctx = context.WithValue(ctx, goa.ServiceKey, "product")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service product.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleProductOrigin(h)
	mux.Handle("OPTIONS", "/products", h.ServeHTTP)
	mux.Handle("OPTIONS", "/barcode", h.ServeHTTP)
	mux.Handle("OPTIONS", "/token", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleProductOrigin applies the CORS response headers corresponding to the
// origin for the service product.
func HandleProductOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Headers", "Authorization")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}