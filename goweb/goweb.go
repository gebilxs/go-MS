package goweb

import (
	"fmt"
	"log"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type routerGroup struct {
	name          string
	handleFuncMap map[string]HandleFunc
	//设置handlerMethodMap 前面是请求方式，后面是路由
	handlerMethodMap map[string][]string
}

func (r *routerGroup) Add(name string, handleFunc HandleFunc) {
	r.handleFuncMap[name] = handleFunc
}

// 下面各种请求方式在对应的包中都有常量
func (r *routerGroup) Any(name string, handleFunc HandleFunc) {
	r.handleFuncMap[name] = handleFunc
	r.handlerMethodMap["ANY"] = append(r.handlerMethodMap["ANY"], name)
}

//GET

func (r *routerGroup) Get(name string, handleFunc HandleFunc) {
	r.handleFuncMap[name] = handleFunc
	r.handlerMethodMap[http.MethodGet] = append(r.handlerMethodMap[http.MethodGet], name)
}

//Post

func (r *routerGroup) Post(name string, handleFunc HandleFunc) {
	r.handleFuncMap[name] = handleFunc
	r.handlerMethodMap[http.MethodPost] = append(r.handlerMethodMap[http.MethodPost], name)
}

//user get->handle
//goods
//order
type router struct {
	routerGroups  []*routerGroup
	handleFuncMap map[string]HandleFunc
}

//add function group
func (r *router) Group(name string) *routerGroup {
	routerGroup := &routerGroup{
		name:             name,
		handleFuncMap:    make(map[string]HandleFunc),
		handlerMethodMap: make(map[string][]string),
	}
	r.routerGroups = append(r.routerGroups, routerGroup)
	return routerGroup
}

//增加对应的方法-> 取消 这时候由路由组进行转发
//func (r *router) Add(name string, handleFunc HandleFunc) {
//	r.handleFuncMap[name] = handleFunc
//}

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router: router{},
	}
}

//实现ServerHTTP
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, group := range e.routerGroups {
		for name, methodHandle := range group.handleFuncMap {
			url := "/" + group.name + name
			if r.RequestURI == url {
				routers, ok := group.handlerMethodMap["ANY"]
				if ok {
					for _, routerName := range routers {
						if routerName == name {
							methodHandle(w, r)
							return
						}
					}
				}
				//method进行匹配
				routers, ok = group.handlerMethodMap[method]
				if ok {
					for _, routerName := range routers {
						if routerName == name {
							methodHandle(w, r)
							return
						}
					}
				}
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "%s %s not allowed\n", r.RequestURI, method)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "%s not allowed\n", r.RequestURI)
}
func (e *Engine) Run() {
	//路由组循环
	//user key : get value : func

	//注释掉旧的循环方式
	//for _, group := range e.routerGroups {
	//	for key, value := range group.handleFuncMap {
	//		http.HandleFunc("/"+group.name+key, value)
	//	}
	//}

	//单纯的路由循环
	//for key, value := range e.handleFuncMap {
	//	http.HandleFunc(key, value)
	//}
	http.Handle("/", e)
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
