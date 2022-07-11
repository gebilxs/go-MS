package goweb

import (
	"log"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type routerGroup struct {
	name          string
	handleFuncMap map[string]HandleFunc
}

func (r *routerGroup) Add(name string, handleFunc HandleFunc) {
	r.handleFuncMap[name] = handleFunc
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
		name:          name,
		handleFuncMap: make(map[string]HandleFunc),
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
func (e *Engine) Run() {
	//路由组循环
	//user key : get value : func
	for _, group := range e.routerGroups {
		for key, value := range group.handleFuncMap {
			http.HandleFunc("/"+group.name+key, value)
		}
	}
	//单纯的路由循环
	//for key, value := range e.handleFuncMap {
	//	http.HandleFunc(key, value)
	//}
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
