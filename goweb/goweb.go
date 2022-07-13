package goweb

import (
	"fmt"
	"log"
	"net/http"
)

const ANY = "ANY"

type HandleFunc func(ctx *Context)

type routerGroup struct {
	name          string
	handleFuncMap map[string]map[string]HandleFunc
	//设置handlerMethodMap 前面是请求方式，后面是路由
	handlerMethodMap map[string][]string
	treeNode         *treeNode
}

//func (r *routerGroup) Add(name string, handleFunc HandleFunc) {
//	r.handleFuncMap[name] = handleFunc
//}

func (r *routerGroup) handle(name string, method string, handleFunc HandleFunc) {

	_, ok := r.handleFuncMap[name]
	if !ok {
		r.handleFuncMap[name] = make(map[string]HandleFunc)
	}
	_, ok = r.handleFuncMap[name][method]
	if ok {
		panic("There are duplicate routes!")
	}
	r.handleFuncMap[name][method] = handleFunc

	r.treeNode.Put(name)
}

// 下面各种请求方式在对应的包中都有常量
func (r *routerGroup) Any(name string, handleFunc HandleFunc) {
	r.handle(name, ANY, handleFunc)
}

//GET

func (r *routerGroup) Get(name string, handleFunc HandleFunc) {
	//r.handleFuncMap[name] = handleFunc
	//r.handlerMethodMap[http.MethodGet] = append(r.handlerMethodMap[http.MethodGet], name)
	r.handle(name, http.MethodGet, handleFunc)
}

//Post

func (r *routerGroup) Post(name string, handleFunc HandleFunc) {
	//r.handleFuncMap[name] = handleFunc
	//r.handlerMethodMap[http.MethodPost] = append(r.handlerMethodMap[http.MethodPost], name)
	r.handle(name, http.MethodPost, handleFunc)
}

//user get->handle
//goods
//order
type router struct {
	routerGroups  []*routerGroup
	handleFuncMap map[string]HandleFunc
}

//add function group
//这里进行初始化
func (r *router) Group(name string) *routerGroup {
	routerGroup := &routerGroup{
		name:             name,
		handleFuncMap:    make(map[string]map[string]HandleFunc), //以方法名作为key，map【key】覆盖啊，应该判断key存在的时候，去新增对应map【string】HandleFunc
		handlerMethodMap: make(map[string][]string),
		treeNode:         &treeNode{name: "/", children: make([]*treeNode, 0)},
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
		routerName := SubStringLast(r.RequestURI, "/"+group.name)
		node := group.treeNode.Get(routerName)
		if node != nil && node.isEnd {
			//匹配
			ctx := &Context{
				W: w,
				R: r,
			}
			handle, ok := group.handleFuncMap[node.routerName][ANY]
			if ok {
				handle(ctx)
				return
			}
			handle, ok = group.handleFuncMap[node.routerName][method]
			if ok {
				handle(ctx)
				return
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "%s %s not allowed\n", r.RequestURI, method)
			return
			//for name, methodHandle := range group.handleFuncMap {
			//	url := "/" + group.name + name
			//	if r.RequestURI == url {
			//		//更新上下文
			//		ctx := &Context{
			//			W: w,
			//			R: r,
			//		}
			//		handle, ok := methodHandle[ANY]
			//		if ok {
			//			handle(ctx)
			//			return
			//		}
			//		handle, ok = methodHandle[method]
			//		if ok {
			//			handle(ctx)
			//			return
			//		}

			//routers, ok := group.handlerMethodMap["ANY"]
			//if ok {
			//	for _, routerName := range routers {
			//		if routerName == name {
			//			methodHandle(w, r)
			//			return
			//		}
			//	}
			//}
			//method进行匹配
			//routers, ok = group.handlerMethodMap[method]
			//if ok {
			//	for _, routerName := range routers {
			//		if routerName == name {
			//			methodHandle(w, r)
			//			return
			//		}
			//	}
			//}
			//Error 405
			//w.WriteHeader(http.StatusMethodNotAllowed)
			//fmt.Fprintf(w, "%s %s not allowed\n", r.RequestURI, method)
			//		//return
			//	}
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
