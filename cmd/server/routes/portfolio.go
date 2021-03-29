package routes

import (
	"net/http"

	"github.com/yarf-framework/yarf"
)

type Portfolio struct {
	yarf.Resource
}

func (p *Portfolio) Get(c *yarf.Context) error {
	c.Redirect("https://sites.google.com/view/nprincedev", http.StatusSeeOther)

	return nil
}
