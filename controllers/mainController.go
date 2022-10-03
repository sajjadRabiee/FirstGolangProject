package controllers

import (
	"FirstGolangProject/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type helloWorld struct {
	connection *repositories.Connection
}

func NewHelloWorld(connection *repositories.Connection) HelloWorld {
	return &helloWorld{
		connection: connection,
	}
}

func (h *helloWorld) Index(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(r.Method)
	modelsSlice := h.connection.All()
	for _, model := range modelsSlice {
		encoder.Encode(model)
	}
}

func (h *helloWorld) FindById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(r.Method)
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 0, 32)
	model := h.connection.FindById(int(id))
	encoder.Encode(model)
}

func (h *helloWorld) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
	h.connection.Insert()
}
