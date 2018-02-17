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
	o.Register("get", "/test", func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintf(w, "test")
		return nil
	})
	// respond error
	o.Register("get", "/error", func(w http.ResponseWriter, r *http.Request) error {
		return fmt.Errorf("Server Error")
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

func TestServerError(t *testing.T) {
	t.Parallel()

	// set Ocelot and route
	o := New()
	// respond error
	o.Register("get", "/error", func(w http.ResponseWriter, r *http.Request) error {
		return fmt.Errorf("Server Error")
	})

	// test error response
	req := httptest.NewRequest("GET", "/error", nil)
	rec := httptest.NewRecorder()
	o.ServeHTTP(rec, req)
	// check result
	if rec.Code != http.StatusInternalServerError {
		t.Error(rec.Code)
	}

	// error handler
	o.ServerError(func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintf(w, "error handler registered")
		return nil
	})

	// test error response
	rec2 := httptest.NewRecorder()
	o.ServeHTTP(rec2, req)
	// check result
	if rec.Code != http.StatusInternalServerError {
		t.Error(rec.Code)
	}
	if rec.Body.String() == rec2.Body.String() || rec2.Body.String() != "error handler registered" {
		t.Error(rec.Body.String())
	}
}

func TestNotFound(t *testing.T) {
	t.Parallel()

	// set Ocelot and route
	o := New()

	// test error response
	req := httptest.NewRequest("GET", "/notfound", nil)
	rec := httptest.NewRecorder()
	o.ServeHTTP(rec, req)
	// check result
	if rec.Code != http.StatusNotFound {
		t.Error(rec.Code)
	}
}
