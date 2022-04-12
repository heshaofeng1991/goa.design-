package design

import (
	_ "goa/design/service"
	_ "goa/design/service/integrations"
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
) //nolint:nolintlint

var _ = API("openapi", func() {
	Title("NextSmartShip openapi")
	Description("NextSmartShip openapi for oms and wms")

	cors.Origin("*", func() {
		cors.Headers("X-Requested-With, Content-Type, Accept, Origin, Authorization, X-Api-Version, x-nss-tenant-id")
		cors.Credentials()
		cors.Methods("GET", "POST", "OPTIONS", "PUT", "DELETE", "PATCH")
	})

	// Server describes a single process listening for client requests. The DSL
	// defines the set of services that the server hosts as well as hosts details.
	Server("oms", func() {
		Description("Order Management Service")

		// List the services hosted by this server.
		Services([]string{
			"healthy", "quote", "swagger", "track", "order",
			"product", "integrations", "auth", "file", "user",
			"woocommerce", "tenant", "platform_product", "wix", "wix_redirect",
		}...)

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
