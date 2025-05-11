package bee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func NEW() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(pattern string, method string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
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
	c := newContext(w, r)
	engine.router.handle(c)
}
