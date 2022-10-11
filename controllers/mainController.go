package controllers

import (
	"FirstGolangProject/repositories"
	"encoding/json"
	"io"
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
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	modelsSlice := h.connection.All()
	encoder.Encode(modelsSlice)
}

func (h *helloWorld) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 0, 32)
	model := h.connection.FindById(int(id))
	encoder.Encode(model)
}

func (h *helloWorld) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	model := h.connection.Model.NewInstance()
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &model)
	h.connection.Insert(model)
	json.NewEncoder(w).Encode(model)
}
