package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
) //nolint:revive

var _ = API("openapi", func() {
	Title("NextSmartShip openapi")
	Description("NextSmartShip openapi for oms and wms")

	cors.Origin("*", func() {
		cors.Headers("Authorization")
	})

	// Server describes a single process listening for client requests. The DSL
	// defines the set of services that the server hosts as well as hosts details.
	Server("oms", func() {
		Description("Order Management Service")

		// List the services hosted by this server.
		Services([]string{"healthy", "quote", "swagger", "track", "order", "product", "file"}...)

		// List the Hosts and their transport URLs.
		Host("localhost", func() {
			// Transport specific URLs, supported schemes are:
			// 'http', 'https', 'grpc' and 'grpcs' with the respective default
			// ports: 80, 443, 8080, 8443.
			URI("http://0.0.0.0:80")
			// Variable describes a URI variable.
			Variable("version", String, "API version", func() {
				// URL parameters must have a default value and/or an
				// enum validation.
				Default("v1")
			})
		})
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint`)
	Scope("api:read")  // Enforce presence of both "api:read"
	Scope("api:write") // and "api:write" scopes in OAuth2 claims.
})

// APIKeyAuth defines a security scheme that uses API keys.
var APIKeyAuth = APIKeySecurity("api_key", func() {
	Description("Secures endpoint by requiring an API key.")
})

// BasicAuth defines a security scheme using basic authentication. The scheme
// protects the "signin" action used to create JWTs.
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during signin")
})
