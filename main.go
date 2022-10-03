package main

import (
	"FirstGolangProject/database"
	"FirstGolangProject/repositories"
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
	db := database.Connect()
	repository := repositories.NewConnection(db)
	helloWorld := controllers.NewHelloWorld(repository)
	http.HandleFunc("/", helloWorld.Index)
	http.HandleFunc("/find", helloWorld.FindById)
	http.HandleFunc("/create", helloWorld.Create)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
