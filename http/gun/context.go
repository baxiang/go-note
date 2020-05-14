package gun

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req *http.Request
	Path string
	Method string
	StatusCode int
}

func NewContext(w http.ResponseWriter,req *http.Request)*Context{
	return &Context{
		Writer:     w,
		Req:        req,
		Path:       req.URL.Path,
		Method:     req.Method,
	}
}


func(c *Context)SetHeader(key,value string){
	c.Writer.Header().Set(key,value)
}
func (c *Context)SetStatusCode(statusCode int){
	c.StatusCode = statusCode
	c.Writer.WriteHeader(statusCode)
}


func(c *Context)JSON(statusCode int,obj interface{}){
	c.SetHeader("Content-Type","application/json")
	c.SetStatusCode(statusCode)
	encoder := json.NewEncoder(c.Writer)
	if err :=encoder.Encode(obj);err!=nil{
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) String(statusCode int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatusCode(statusCode)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) HTML(SetStatusCode int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatusCode(SetStatusCode)
	c.Writer.Write([]byte(html))
}