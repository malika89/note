### day1: http
原生http实现gee 框架(serverhttp接口和handler)
```go
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
router map[string]HandlerFunc
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	handler, ok := engine.router[key]
}
```

### day2:上下文
 + context
  ```go
    type Context struct {
    // origin objects
    Writer http.ResponseWriter
    Req    *http.Request
    // request info
    Path   string
    Method string
    // response info
    StatusCode int }
  ```
 + router => handlerfunc集成
  ``` go
   type router struct {
     handlers map[string]HandlerFunc
   }
   ```

### day3:前缀树路由-动态路由
 + tireTree.go
  ```go
  type node struct {
     pattern  string // 待匹配路由，例如 /p/:lang
     part     string // 路由中的一部分，例如 :lang
     children []*node // 子节点，例如 [doc, tutorial, intro]
     isWild   bool // 是否精确匹配，part 含有 : 或 * 时为true
  }
  ```
 + router
  ```go
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}
  ```

### day4: 路由分组控制(Route Group Control)
 + routeGroup
  ```go
  type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share a Engine instance
  }
  ```
 + engine 