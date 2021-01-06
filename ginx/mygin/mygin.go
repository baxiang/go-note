package mygin

import (
	"fmt"
	"math"
	"net/http"
	"path"
	"sync"
)

const abortIndex int8 = math.MaxInt8 / 2

type HandlerFun func(ctx *Context)

type IRouter interface {
	Use(...HandlerFun) IRouter
	GET(string, ...HandlerFun) IRouter
	Group(string, ...HandlerFun) *RouterGroup
}

type RouterGroup struct {
	Handlers []HandlerFun
	engine   *Engine
	basePath string
}
type Context struct {
	Request       *http.Request
	ResponseWrite http.ResponseWriter
	engine        *Engine
}

type Engine struct {
	router map[string][]HandlerFun
	pool   sync.Pool
	RouterGroup
}

func NewEngine() *Engine {
	e := &Engine{}
	e.router = make(map[string][]HandlerFun)
	e.pool.New = func() interface{} {
		return e.allocateContext()
	}
	e.RouterGroup = RouterGroup{
		basePath: "/",
		Handlers: nil,
		engine:   e,
	}
	return e
}
func (engine *Engine) allocateContext() *Context {
	return &Context{engine: engine}
}
func (engine *Engine) Run(addr string) error {
	fmt.Println("http start ", addr)
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := engine.pool.Get().(*Context)
	c.ResponseWrite = w
	c.Request = r
	engine.handleHTTPRequest(c)
	engine.pool.Put(c)
}

func (engine *Engine) handleHTTPRequest(c *Context) {
	method := c.Request.Method
	path := c.Request.URL.Path
	key := engine.handleRequestKey(method, path)
	if handlers, ok := engine.router[key]; ok {
		for _, h := range handlers {
			h(c)
		}
	}
}
func (e *Engine) handleRequestKey(httpMethod, path string) string {
	return fmt.Sprintf("%s-%s", httpMethod, path)
}
func (engine *Engine) addRouter(httpMethod, absolutePath string, handlers []HandlerFun) {
	key := engine.handleRequestKey(httpMethod, absolutePath)
	engine.router[key] = handlers
}

func (g *RouterGroup) Group(path string, handlers ...HandlerFun) *RouterGroup {
	rg := &RouterGroup{}
	rg.Handlers = g.CombineHandler(handlers)
	rg.basePath = path
	rg.engine = g.engine
	return rg
}

func (group *RouterGroup) Use(handlers ...HandlerFun) IRouter {
	group.Handlers = append(group.Handlers, handlers...)
	return group
}
func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPaths(group.basePath, relativePath)
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}
	return finalPath
}
func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func (group *RouterGroup) CombineHandler(handlers []HandlerFun) []HandlerFun {

	finalSize := len(group.Handlers) + len(handlers)
	if finalSize >= int(abortIndex) {
		panic("too many handlers")
	}
	mergedHandlers := make([]HandlerFun, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers

}
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers []HandlerFun) IRouter {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.CombineHandler(handlers)
	group.engine.addRouter(httpMethod, absolutePath, handlers)
	return group
}

func (g *RouterGroup) GET(path string, handles ...HandlerFun) IRouter {
	g.handle("GET", path, handles)
	return g
}
