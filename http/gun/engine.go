package gun

import (
	"net/http"
)
type HandlerFunc func(*Context)


type Engine struct {
	router *router
}

func New()*Engine{
	return &Engine{router: newRouter()}
}

func (engine *Engine)ServeHTTP(w http.ResponseWriter,r *http.Request){
	  c :=NewContext(w,r)
	  engine.router.handle(c)
}

func(engine *Engine)addRoute(method,path string, handler HandlerFunc){
	engine.router.addRouter(method,path,handler)
}

func(engine *Engine)Get(path string, handler HandlerFunc){
	engine.addRoute("GET",path,handler)
}
func(engine *Engine)Post(path string, handler HandlerFunc){
	engine.addRoute("POST",path,handler)
}

func (engine *Engine)Run(addr string)error{
	return http.ListenAndServe(addr,engine)
}

