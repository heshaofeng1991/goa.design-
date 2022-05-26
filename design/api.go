package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
	_ "goa/design/service"
)

var _ = API("test", func() {
	Title("test")
	Description("test")

	cors.Origin("*", func() {
		cors.Headers("X-Requested-With, Content-Type, Accept, Origin, Authorization, X-Api-Version")
		cors.Credentials()
		cors.Methods("GET", "POST", "OPTIONS", "PUT", "DELETE", "PATCH")
	})

	// Server describes a single process listening for client requests. The DSL
	// defines the set of services that the server hosts as well as hosts details.
	Server("test", func() {
		Description("tets")

		// List the services hosted by this server.
		Services([]string{
			"channel",
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
