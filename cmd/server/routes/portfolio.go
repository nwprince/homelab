package routes

import (
	"github.com/yarf-framework/yarf"
)

type Portfolio struct {
	yarf.Resource
}

func (p *Portfolio) Get(c *yarf.Context) error {
	c.Render("Portfolio endpoint hit")

	return nil
}
