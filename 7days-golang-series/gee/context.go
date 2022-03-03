package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//将路由(router)独立出来，方便之后增强。
//设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持

type H map[string]interface{} //header

//将handlerfunc(w,req)==>handlerfunc(c *Context)

type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter
	Path   string
	Method string
	Params map[string]string //数存储到Params
	StatusCode int
}

func newContext(w http.ResponseWriter,req *http.Request)  *Context{
	return &Context{
		Req:    req,
		Writer: w,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm post 参数,get 方法查询
func(c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func(c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status  设置status
func(c *Context) Status(code int)  {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// SetHeader 设置请求header
func(c *Context) SetHeader(key,value string)  {
	c.Req.Header.Set(key,value)
}

//返回body格式转换=>String/Data/JSON/HTML
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, respObj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(respObj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}




