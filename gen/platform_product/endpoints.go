// Code generated by goa v3.6.2, DO NOT EDIT.
//
// platform_product endpoints
//
// Command:
// $ goa gen goa/design -o ./

package platformproduct

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "platform_product" service endpoints.
type Endpoints struct {
	PlatformProduct goa.Endpoint
	Convert         goa.Endpoint
	Mappings        goa.Endpoint
}

// NewEndpoints wraps the methods of the "platform_product" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		PlatformProduct: NewPlatformProductEndpoint(s, a.JWTAuth),
		Convert:         NewConvertEndpoint(s, a.JWTAuth),
		Mappings:        NewMappingsEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "platform_product" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.PlatformProduct = m(e.PlatformProduct)
	e.Convert = m(e.Convert)
	e.Mappings = m(e.Mappings)
}

// NewPlatformProductEndpoint returns an endpoint function that calls the
// method "platform_product" of service "platform_product".
func NewPlatformProductEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetListing)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.PlatformProduct(ctx, p)
	}
}

// NewConvertEndpoint returns an endpoint function that calls the method
// "convert" of service "platform_product".
func NewConvertEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ConvertPlatformProductsReq)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Convert(ctx, p)
	}
}

// NewMappingsEndpoint returns an endpoint function that calls the method
// "mappings" of service "platform_product".
func NewMappingsEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateMappingsReq)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Mappings(ctx, p)
	}
}