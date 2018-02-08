package ocelot

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOcelot(t *testing.T) {
	t.Parallel()

	o := New()

	o.Get("/test", func(c *Context) error {
		fmt.Fprintf(c.w, "test")
		return nil
	})

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	o.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Error(rec.Code)
	}

	if rec.Body.String() != "test" {
		t.Error(rec.Body.String())
	}
}
