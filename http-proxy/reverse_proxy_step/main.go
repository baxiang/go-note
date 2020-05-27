package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"errors"
)

func main() {
	//http://127.0.0.1:8001/v1/xxxx/xxx/
	addr := flag.String("addr", "http://127.0.0.1:8001/v1", "server addresss")
	flag.Parse()
	urlTarget, err := url.Parse(*addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	proxy := singleHostReverseProxy(urlTarget)
	log.Fatal(http.ListenAndServe(":8080",proxy))
}


func singleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}


	modifyFunc := func(res *http.Response) error {
		if res.StatusCode != 200 {
			return errors.New("error statusCode")
			oldPayload, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}
			newPayLoad := []byte("hello " + string(oldPayload))



			res.Body = ioutil.NopCloser(bytes.NewBuffer(newPayLoad))
			res.ContentLength = int64(len(newPayLoad))
			res.Header.Set("Content-Length", fmt.Sprint(len(newPayLoad)))
		}
		return nil
	}
	errorHandler := func(res http.ResponseWriter, req *http.Request, err error) {
		res.Write([]byte(err.Error()))
	}

	return &httputil.ReverseProxy{Director: director,ModifyResponse:modifyFunc,ErrorHandler:errorHandler}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}