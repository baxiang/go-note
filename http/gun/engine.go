package gun

import (
	"net/http"
)
type HandlerFunc func(*Context)


type RouterGroup struct {
	prefix string
	middleware []HandlerFunc
	parent *RouterGroup
	engine *Engine
}


type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New()*Engine{
	 engine :=&Engine{router: newRouter()}
	 engine.RouterGroup = &RouterGroup{engine: engine}
	 engine.groups = []*RouterGroup{engine.RouterGroup}
	 return engine
}

func(g *RouterGroup)Group(prefix string)*RouterGroup{
	engine := g.engine
	newGroup := &RouterGroup{
		prefix:      prefix,
		parent:      g,
		engine:      engine,
	}
	engine.groups = append(engine.groups,newGroup)
	return newGroup
}


func (engine *Engine)ServeHTTP(w http.ResponseWriter,r *http.Request){
	  c :=NewContext(w,r)
	  engine.router.handle(c)
}

func(g *RouterGroup)addRoute(method,path string, handler HandlerFunc){
	pattern :=g.prefix+path
	g.engine.router.addRoute(method,pattern,handler)
}

func(g *RouterGroup)Get(path string, handler HandlerFunc){
	g.addRoute("GET",path,handler)
}
func(g *RouterGroup)Post(path string, handler HandlerFunc){
	g.addRoute("POST",path,handler)
}

func (engine *Engine)Run(addr string)error{
	return http.ListenAndServe(addr,engine)
}

