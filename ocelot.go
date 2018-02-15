package ocelot

import (
	"net/http"
	"strings"
)

type (
	// Ocelot is the top level framework instance.
	Ocelot struct {
		routes map[string]route
	}

	// route has handler and matching information.
	route struct {
		method  string
		path    string
		handler HandlerFunc
	}

	// HandlerFunc defines a function to server HTTP requests.
	HandlerFunc func(http.ResponseWriter, *http.Request) error
)

// New creates new instance of Ocelot
func New() (o *Ocelot) {
	o = &Ocelot{
		routes: map[string]route{},
	}
	return
}

// Register adds new route path with method and handler
func (o *Ocelot) Register(method string, path string, h HandlerFunc) {
	m := strings.ToUpper(method)

	o.routes[m+path] = route{
		method:  m,
		path:    path,
		handler: h,
	}
}

// ServeHTTP match registered routes against request path
func (o *Ocelot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range o.routes {
		if r.Method == route.method && r.URL.Path == route.path {
			route.handler(w, r)
		}
	}
}

// Start http server
func (o *Ocelot) Start(address string) error {
	return http.ListenAndServe(address, o)
}
