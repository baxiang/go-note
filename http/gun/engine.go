package gun

import (
	"net/http"
	"strings"
)
type HandlerFunc func(*Context)


type RouterGroup struct {
	prefix string
	middleware []HandlerFunc
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

func Default()*Engine{
	engine :=New()
	engine.Use(Logger(),Recovery())
	return engine
}

func(g *RouterGroup)Group(prefix string)*RouterGroup{
	engine := g.engine
	newGroup := &RouterGroup{
		prefix:      prefix,
		engine:      engine,
	}
	engine.groups = append(engine.groups,newGroup)
	return newGroup
}


func (engine *Engine)ServeHTTP(w http.ResponseWriter,r *http.Request){
	var middleware []HandlerFunc
	for _,group :=range engine.groups{
		if strings.HasPrefix(r.URL.Path,group.prefix){
			middleware = append(middleware,group.middleware...)
		}
	}
	  c :=NewContext(w,r)
	  c.handlers = middleware
	  engine.router.handle(c)
}

func(g *RouterGroup)Use(middleware ...HandlerFunc){
	g.middleware = append(g.middleware,middleware...)
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

