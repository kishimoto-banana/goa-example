package design

import (
	. "goa.design/goa/v3/dsl"
)

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
