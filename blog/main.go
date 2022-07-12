package main

import (
	"fmt"
	"go-MS/goweb"
	"net/http"
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

	g.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s welcome to golang world!", "xixi~")
	})
	g.Post("/xixi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s welcome to golang world!", "heloo")
	})
	//order := engine.Group("order")
	//order.Add("/get", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "%s 查询订单", "xck")
	//})
	engine.Run()
}
