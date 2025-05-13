package bee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
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
