package controllers

import (
	"fmt"
	"net/http"
)

type helloWorld struct {
	hello string
	world string
	count int
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
	salamIran := helloWorld{
		hello: "salam",
		world: "iran",
		count: 5,
	}
	for i := 0; i <= salamIran.count; i++ {
		fmt.Fprintln(w, salamIran.hello, salamIran.world)
	}
}

func Exist(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
	fruitMap := make(map[string]int)
	fruitMap["Banna"] = 5
	fruit, ok := fruitMap["Banna"]
	if !ok {
		fmt.Fprintln(w, "dose not exist", fruit)
	} else {
		fmt.Fprintln(w, "exist", fruit)
	}
}
