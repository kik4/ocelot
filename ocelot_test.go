package ocelot

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOcelot(t *testing.T) {
	t.Parallel()

	// set Ocelot and route
	o := New()
	o.Register("get", "/test", func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintf(w, "test")
		return nil
	})

	// test router
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	o.ServeHTTP(rec, req)

	// check result
	if rec.Code != http.StatusOK {
		t.Error(rec.Code)
	}
	if rec.Body.String() != "test" {
		t.Error(rec.Body.String())
	}
}
