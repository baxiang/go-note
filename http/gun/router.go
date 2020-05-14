package gun

import "net/http"

type router struct {
	handlers map[string] HandlerFunc
}

func newRouter()*router{
	return &router{
		handlers: map[string]HandlerFunc{},
	}
}

func(r *router)addRouter(method,path string,handle HandlerFunc){
	 key :=method+"-"+path
	 r.handlers[key] = handle
}

func(r *router)handle(c *Context){
	key :=c.Method+"-"+c.Path
	if handler,ok :=r.handlers[key];ok{
		handler(c)
	}else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}