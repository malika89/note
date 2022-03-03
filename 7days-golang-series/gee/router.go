package gee

import (
	"net/http"
	"strings"
)
//添加动态路由，通过tiretree.支持两种模式:name和*filepath

type router struct {
	roots    map[string]*node //{"Get":{}:}
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots: make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts :=parsePattern(pattern)
	_,ok :=r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern,parts,0) //顶层添加royte

	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) getRoute(method,path string) (*node,map[string]string) {
	if _,ok :=r.roots[method];!ok{
		return nil,nil
	}
	searchParts :=parsePattern(path)
	n :=r.roots[method].search(searchParts,0)
	if n ==nil{
		return nil,nil
	}
	params := make(map[string]string)
	parts :=parsePattern(n.pattern)

	for index, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[index]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[index:], "/")
			break
		}
	}
	return n, params
}

func (r *router) handle(c *Context) {
	n,params :=r.getRoute(c.Method,c.Path)
	if n!=nil {
		c.Params = params //解析出来的路由参数
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
