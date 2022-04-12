// Code generated by goa v3.6.2, DO NOT EDIT.
//
// quote endpoints
//
// Command:
// $ goa gen goa/design -o ./

package quote

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "quote" service endpoints.
type Endpoints struct {
	Get  goa.Endpoint
	Post goa.Endpoint
}

// NewEndpoints wraps the methods of the "quote" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Get:  NewGetEndpoint(s, a.JWTAuth),
		Post: NewPostEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "quote" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Get = m(e.Get)
	e.Post = m(e.Post)
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "quote".
func NewGetEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetQuote)
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
		return s.Get(ctx, p)
	}
}

// NewPostEndpoint returns an endpoint function that calls the method "post" of
// service "quote".
func NewPostEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PostQuote)
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
		return s.Post(ctx, p)
	}
}
