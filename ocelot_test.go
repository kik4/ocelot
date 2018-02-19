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
	// regular
	o.Register("get", "/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test")
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

func TestNotFound(t *testing.T) {
	t.Parallel()

	// set Ocelot and route
	o := New()

	// test response
	req := httptest.NewRequest("GET", "/notfound", nil)
	rec := httptest.NewRecorder()
	o.ServeHTTP(rec, req)
	// check result
	if rec.Code != http.StatusNotFound {
		t.Error(rec.Code)
	}

	// notfound handler
	o.NotFound(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "notfound handler registered")
	})

	// test response
	rec2 := httptest.NewRecorder()
	o.ServeHTTP(rec2, req)
	// check result
	if rec.Code != http.StatusNotFound {
		t.Error(rec.Code)
	}
	if rec.Body.String() == rec2.Body.String() || rec2.Body.String() != "notfound handler registered" {
		t.Error(rec.Body.String())
	}
}
