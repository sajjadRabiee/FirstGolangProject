package main

import (
	"awesomeProject/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/exist", controllers.Exist)
	log.Fatal(http.ListenAndServe(":8500", nil))
}
