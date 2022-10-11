package main

import (
	"FirstGolangProject/database"
	"FirstGolangProject/models"
	"FirstGolangProject/repositories"
	"database/sql"
	"flag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"path/filepath"

	"FirstGolangProject/controllers"
)

var (
	addr = flag.String("addr", ":8500", "http service address")
)

func main() {
	flag.Parse()
	db := database.Connect()
	configInit(".env")
	log.Println(viper.GetString("MODE"))

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

func configInit(configFile string) {
	viper.SetConfigType("env")
	viper.SetConfigFile(filepath.Base(configFile))
	viper.AddConfigPath(filepath.Dir(configFile))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("fatal error config file: ", err)
	}
}
