package design

import (
	- "github.com/kishimoto-banana/goa-example/design/resource"
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
