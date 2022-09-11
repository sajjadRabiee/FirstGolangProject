package main

import (
	"flag"
	"log"
	"net/http"

	"awesomeProject/controllers"
)

var (
	addr = flag.String("addr", ":8500", "http service address")
)

func main() {
	flag.Parse()
	helloWorld := controllers.NewHelloWorld("Hello", "World", 5)
	http.HandleFunc("/", helloWorld.Index)
	http.HandleFunc("/exist", helloWorld.Exist)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
