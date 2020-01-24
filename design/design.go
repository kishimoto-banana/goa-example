package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("example-api", func() {
	Title("exmaple api")
	Description("Service for example")
	Server("example-api", func() {
		Host("localhost", func() {
			URI("http://localhost:8000") 
		})
	})
})

var _ = Service("ping", func() {
	Description("Ping/Pong service")
	Method("Return pong", func() {
		Result(String)
		HTTP(func() {
			GET("/ping")
			Response(StatusOK, func() {
				ContentType("text/html")
			})
		})
	})
})
