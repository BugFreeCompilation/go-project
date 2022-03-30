package service

import (
	"net/http"
)

type Service interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	Add(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
}
