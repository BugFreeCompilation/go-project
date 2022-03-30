package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type routerStruct struct{}

var (
	router = mux.NewRouter()
)

func InitMuxRouter() Router {
	return &routerStruct{}
}

func (*routerStruct) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(uri, f).Methods("GET")
}

func (*routerStruct) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(uri, f).Methods("POST")
}

func (*routerStruct) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(uri, f).Methods("DELETE")
}

func (*routerStruct) SERVE(port string) {
	fmt.Printf("Using Port: %v", port)
	http.ListenAndServe(port, router)
}
