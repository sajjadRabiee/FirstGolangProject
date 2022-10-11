package controllers

import (
	"net/http"
)

type HelloWorld interface {
	Index(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}
