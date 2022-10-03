package main

import (
	"FirstGolangProject/database"
	"FirstGolangProject/models"
	"FirstGolangProject/repositories"
	"database/sql"
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

	handle(db, &models.Hobby{}, "/hobbies")
	handle(db, &models.KeyValue{}, "/key-values")
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handle(db *sql.DB, model models.BaseModel, baseName string) {
	repository := repositories.NewConnection(db, model)
	helloWorld := controllers.NewHelloWorld(repository)
	http.HandleFunc(baseName+"/", helloWorld.Index)
	http.HandleFunc(baseName+"/find", helloWorld.FindById)
	http.HandleFunc(baseName+"/create", helloWorld.Create)
}
