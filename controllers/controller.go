package controllers

import (
	"net/http"
)

type HelloWorld interface {
	Index(w http.ResponseWriter, r *http.Request)
	Exist(w http.ResponseWriter, r *http.Request)
}
