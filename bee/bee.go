package bee

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	route map[string]HandlerFunc
}

func NEW() *Engine {
	return &Engine{}
}

func (engine *Engine) addRoute(pattern string, method string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.route[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute(pattern, "GET", handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute(pattern, "POST", handler)
}

func (engine *Engine) PUT(pattern string, handler HandlerFunc) {
	engine.addRoute(pattern, "PUT", handler)
}

func (engine *Engine) DELETE(pattern string, handler HandlerFunc) {
	engine.addRoute(pattern, "DELETE", handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.route[key]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}
