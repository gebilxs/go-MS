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
	engine.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s welcome to golang world!", "xixi~")
	})
	engine.Run()
}
