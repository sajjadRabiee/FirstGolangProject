package controllers

import (
	"FirstGolangProject/models"
	"FirstGolangProject/repositories"
	"fmt"
	"net/http"
	"strconv"
)

type helloWorld struct {
	connection *repositories.Connection
	model      models.KeyValue
}

func NewHelloWorld(connection *repositories.Connection) HelloWorld {
	return &helloWorld{
		connection: connection,
	}
}

func (h *helloWorld) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
	keyValues := h.connection.All()
	for _, model := range keyValues {
		fmt.Fprintln(w, model)
	}
}

func (h *helloWorld) FindById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 0, 32)
	model := h.connection.FindById(int(id))
	fmt.Fprintln(w, model)
}

func (h *helloWorld) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
	h.connection.Insert()
}
