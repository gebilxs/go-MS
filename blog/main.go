package main

import (
	"fmt"
	"go-MS/goweb"
)

func main() {
	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "%s welcome to golang world!", "xixi~")
	//})
	//err := http.ListenAndServe(":8111", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	engine := goweb.New()
	g := engine.Group("user")

	//g.Post("/hello", func(ctx *goweb.Context) {
	//	fmt.Fprintf(ctx.W, "%s post welcome to golang world!", "xixi~")
	//})

	//测试中间件
	g.PreHandle(func(next goweb.HandlerFunc) goweb.HandlerFunc {
		return func(ctx *goweb.Context) {
			fmt.Println("hello world!")
			next(ctx)
		}
	})
	g.Get("/hello/get", func(ctx *goweb.Context) {
		fmt.Println("second world!")
		fmt.Fprintf(ctx.W, "%s with /*/get welcome to golang world!", "test")
	})

	g.Post("/xixi", func(ctx *goweb.Context) {
		fmt.Fprintf(ctx.W, "%s welcome to golang world!", "heloo")
	})
	g.Any("/xixi", func(ctx *goweb.Context) {
		fmt.Fprintf(ctx.W, "%s any welcome to golang world!", "heloo")
	})

	//前缀树测试用数据

	g.Get("/get/:id", func(ctx *goweb.Context) {
		fmt.Fprintf(ctx.W, "%s get user ", "1 frset")
	})
	//order := engine.Group("order")
	//order.Add("/get", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "%s 查询订单", "xck")
	//})
	engine.Run()
}
