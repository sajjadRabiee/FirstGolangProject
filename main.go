package main

import (
	"FirstGolangProject/database"
	"flag"
	"log"
	"net/http"

	"FirstGolangProject/controllers"
)

var (
	addr = flag.String("addr", ":8500", "http service address")
)

func main() {
	flag.Parse()
	database.Connection()
	helloWorld := controllers.NewHelloWorld("Hello", "World", 5)
	http.HandleFunc("/", helloWorld.Index)
	http.HandleFunc("/exist", helloWorld.Exist)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
