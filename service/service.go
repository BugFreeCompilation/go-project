package service

import (
	database "BugFreeCompilation/go-project/database"
	"BugFreeCompilation/go-project/entity"
	"BugFreeCompilation/go-project/error"
	"encoding/json"
	"log"
	"net/http"
)

type serviceStruct struct{}

var (
	mydb database.Database
)

func InitService(db_input database.Database) Service {
	mydb = db_input
	return &serviceStruct{}
}

func (*serviceStruct) GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := mydb.GetAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*serviceStruct) Add(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	mydb.Add(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}

func (*serviceStruct) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	err = mydb.Delete(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(error.ServiceError{Message: "The entry you tried to delete does not exist."})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
