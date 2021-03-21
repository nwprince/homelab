package panic

import "github.com/yarf-framework/yarf"

func Exception(c *yarf.Context) {
	c.Render("This is a custom exception handler")
}
