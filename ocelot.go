package ocelot

import (
	"fmt"
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

// RegisterServerError adds new route path with method and handler
func (o *Ocelot) RegisterServerError(h HandlerFunc) {
	p := "ServerError"
	o.routes[p] = route{
		path:    p,
		handler: h,
	}
}

// ServeHTTP match registered routes against request path
func (o *Ocelot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range o.routes {
		if r.Method == route.method && r.URL.Path == route.path {
			err := route.handler(w, r)
			if err != nil {
				w.WriteHeader(500)
				if h, ok := o.routes["ServerError"]; ok {
					h.handler(w, r)
				} else {
					fmt.Fprintf(w, "500 Server Error")
				}
			}
			return
		}
	}
	w.WriteHeader(404)
	fmt.Fprintf(w, "404 Not Found")
}

// Start http server
func (o *Ocelot) Start(address string) error {
	return http.ListenAndServe(address, o)
}
