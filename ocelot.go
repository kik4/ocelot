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

const (
	pathToServerError string = "_ServerError"
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

// ServerError adds handler called when error occured
func (o *Ocelot) ServerError(h HandlerFunc) {
	o.routes[pathToServerError] = route{
		path:    pathToServerError,
		handler: h,
	}
}

// ServeHTTP match registered routes against request path
func (o *Ocelot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range o.routes {
		if r.Method == route.method && r.URL.Path == route.path {
			err := route.handler(w, r)
			if err != nil {
				// respond server error
				if h, ok := o.routes[pathToServerError]; ok {
					w.WriteHeader(http.StatusInternalServerError)
					h.handler(w, r)
				} else {
					http.Error(w, "500 internal server error", http.StatusInternalServerError)
				}
			}
			return
		}
	}

	// respond not found
	http.NotFound(w, r)
}

// Start http server
func (o *Ocelot) Start(address string) error {
	return http.ListenAndServe(address, o)
}
