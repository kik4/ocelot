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
	HandlerFunc func(http.ResponseWriter, *http.Request)
)

const (
	pathToNotFound string = "_NotFound"
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

// NotFound adds handler called when page not found
func (o *Ocelot) NotFound(h HandlerFunc) {
	o.routes[pathToNotFound] = route{
		path:    pathToNotFound,
		handler: h,
	}
}

// ServeHTTP match registered routes against request path
func (o *Ocelot) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// get handler
	route, ok := o.routes[r.Method+r.URL.Path]
	if !ok {
		// respond not found
		if route, ok := o.routes[pathToNotFound]; ok {
			w.WriteHeader(http.StatusNotFound)
			route.handler(w, r)
		} else {
			http.NotFound(w, r)
		}
		return
	}

	// call handler
	route.handler(w, r)
}

// Start http server
func (o *Ocelot) Start(address string) error {
	return http.ListenAndServe(address, o)
}
