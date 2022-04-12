// Code generated by goa v3.6.2, DO NOT EDIT.
//
// file endpoints
//
// Command:
// $ goa gen goa/design -o ./

package file

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "file" service endpoints.
type Endpoints struct {
	UploadImage goa.Endpoint
}

// NewEndpoints wraps the methods of the "file" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		UploadImage: NewUploadImageEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "file" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.UploadImage = m(e.UploadImage)
}

// NewUploadImageEndpoint returns an endpoint function that calls the method
// "upload_image" of service "file".
func NewUploadImageEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UploadFile)
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
		return s.UploadImage(ctx, p)
	}
}
