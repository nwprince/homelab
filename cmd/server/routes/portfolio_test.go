package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yarf-framework/yarf"
)

func TestPortfolio(t *testing.T) {
	p := new(Portfolio)

	c := new(yarf.Context)
	c.Request, _ = http.NewRequest("GET", "/", nil)
	c.Response = httptest.NewRecorder()

	err := p.Get(c)
	if err != nil {
		t.Error(err.Error())
	}
}
