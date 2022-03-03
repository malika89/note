package gee

import (
	"net/http"
)

// HandlerFunc handler定义
//type HandlerFunc func(w http.ResponseWriter,req *http.Request)
type HandlerFunc func(c *Context)

// Engine  handler和router 映射
type Engine struct {
	//router map[string]HandlerFunc
	router *router
}

func NewGee() *Engine {
	return &Engine{
		//router: make(map[string]HandlerFunc),
		router: newRouter(),
	}
}


// 实现路由映射、查找、添加
func(engine *Engine) addRoute(method,uri string,handler HandlerFunc) {
	engine.router.addRoute(method,uri,handler)
}

func(engine *Engine) Get(uri string,handler HandlerFunc) {
	engine.addRoute("GET",uri,handler)
}

func(engine *Engine) Post(uri string,handler HandlerFunc) {
	engine.addRoute("POST",uri,handler)
}

func(engine *Engine) Put(uri string,handler HandlerFunc) {
	engine.addRoute("PUT",uri,handler)
}

func(engine *Engine) Delete(uri string,handler HandlerFunc) {
	engine.addRoute("DELETE",uri,handler)
}

//实现servehttp- with http

//func(engine *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request)  {
//	key :=req.Method +"-" +req.URL.Path
//	if handler,ok :=e.router[key];ok {
//		handler(w, req)
//	}else {
//		fmt.Fprintf(w,"404 NOT found:=%q \n",req.URL.Path)
//	}
//}
func(engine *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request)  {
	c :=newContext(w,req)
	engine.router.handle(c)
}

func(engine *Engine) Run(addr string)  {
	http.ListenAndServe(addr,engine)
}

