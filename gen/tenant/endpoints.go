// Code generated by goa v3.6.2, DO NOT EDIT.
//
// tenant endpoints
//
// Command:
// $ goa gen goa/design -o ./

package tenant

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "tenant" service endpoints.
type Endpoints struct {
	Integrations     goa.Endpoint
	GetTenantInfo    goa.Endpoint
	UpdateTenantInfo goa.Endpoint
}

// NewEndpoints wraps the methods of the "tenant" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Integrations:     NewIntegrationsEndpoint(s, a.JWTAuth),
		GetTenantInfo:    NewGetTenantInfoEndpoint(s, a.JWTAuth),
		UpdateTenantInfo: NewUpdateTenantInfoEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "tenant" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Integrations = m(e.Integrations)
	e.GetTenantInfo = m(e.GetTenantInfo)
	e.UpdateTenantInfo = m(e.UpdateTenantInfo)
}

// NewIntegrationsEndpoint returns an endpoint function that calls the method
// "integrations" of service "tenant".
func NewIntegrationsEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AuthToken)
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
		return s.Integrations(ctx, p)
	}
}

// NewGetTenantInfoEndpoint returns an endpoint function that calls the method
// "get_tenant_info" of service "tenant".
func NewGetTenantInfoEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AuthToken)
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
		return s.GetTenantInfo(ctx, p)
	}
}

// NewUpdateTenantInfoEndpoint returns an endpoint function that calls the
// method "update_tenant_info" of service "tenant".
func NewUpdateTenantInfoEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*TenantInfo)
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
		return s.UpdateTenantInfo(ctx, p)
	}
}
