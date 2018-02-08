package ocelot

import "net/http"

type (
	Ocelot struct {
		routes map[string]Route
	}

	Context struct {
		r *http.Request
		w http.ResponseWriter
	}

	Route struct {
		method  string
		path    string
		handler HandlerFunc
	}

	HandlerFunc func(*Context) error
)

func New() (o *Ocelot) {
	o = &Ocelot{
		routes: map[string]Route{},
	}
	return
}

func (o *Ocelot) Get(path string, h HandlerFunc) {
	o.routes["GET"+path] = Route{
		method:  "GET",
		path:    path,
		handler: h,
	}
}

func (o *Ocelot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := Context{
		r: r,
		w: w,
	}
	for _, route := range o.routes {
		if r.Method == route.method && r.URL.Path == route.path {
			route.handler(&c)
		}
	}
}

func (o *Ocelot) Start(address string) error {
	http.HandleFunc("/", o.ServeHTTP)
	return http.ListenAndServe(address, nil)
}
