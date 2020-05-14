package gun

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router map[string] http.HandlerFunc
}

func New()*Engine{
	return &Engine{router: map[string]http.HandlerFunc{}}
}

func (engine *Engine)ServeHTTP(w http.ResponseWriter,r *http.Request){
	  key := r.Method+"-"+r.URL.Path
	  if handle,ok:=engine.router[key];ok{
	  	 handle(w,r)
	  }else {
		  fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL.Path)
	  }
}

func(engine *Engine)addRoute(method,path string, handler http.HandlerFunc){
	key := method+"-"+path
	engine.router[key]= handler
}

func(engine *Engine)Get(path string, handler http.HandlerFunc){
	engine.addRoute("GET",path,handler)
}
func(engine *Engine)Post(path string, handler http.HandlerFunc){
	engine.addRoute("POST",path,handler)
}

func (engine *Engine)Run(addr string)error{
	return http.ListenAndServe(addr,engine)
}

