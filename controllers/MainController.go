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

func NewHelloWorld(hello, word string, count int) HelloWorld {
	return &helloWorld{
		hello: hello,
		world: word,
		count: count,
	}
}
func (h *helloWorld) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
	for i := 0; i <= h.count; i++ {
		fmt.Fprintln(w, h.hello, h.world)
	}
}

func (h *helloWorld) Exist(w http.ResponseWriter, r *http.Request) {
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
