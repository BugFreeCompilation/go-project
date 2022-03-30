package service

import (
	database "BugFreeCompilation/go-project/database"
	"BugFreeCompilation/go-project/entity"
	err "BugFreeCompilation/go-project/error"
	"encoding/json"
	"errors"
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
	posts, err_get := mydb.GetAll()

	if err_get != nil {
		response.WriteHeader(http.StatusInternalServerError)
		switch {
		case errors.Is(err_get, err.ErrDatabaseOpen):
			json.NewEncoder(response).Encode(err.ServiceError{Message: "Failed to run GetAll: Error opening the database."})
		case errors.Is(err_get, err.ErrDatabaseQuery):
			json.NewEncoder(response).Encode(err.ServiceError{Message: "Failed to run GetAll: Error quering the database."})
		}
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*serviceStruct) Add(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var post entity.Post
	err_decode := json.NewDecoder(request.Body).Decode(&post)
	if err_decode != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err.ServiceError{Message: "Failed to run Add: Error decoding Post."})
		return
	}

	err_add := mydb.Add(&post)
	if err_add != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err.ServiceError{Message: "Failed to run Add: Adding entry failed."})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}

func (*serviceStruct) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var post entity.Post
	err_decode := json.NewDecoder(request.Body).Decode(&post)
	if err_decode != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err.ServiceError{Message: "Failed to run Delete: Error decoding Post."})
		return
	}

	err_delete := mydb.Delete(&post)
	if err_delete != nil {
		response.WriteHeader(http.StatusInternalServerError)
		switch {
		case errors.Is(err_delete, err.ErrDatabaseEntry):
			json.NewEncoder(response).Encode(err.ServiceError{Message: "Failed to run Delete: Entry does not exist."})
		case errors.Is(err_delete, err.ErrDatabaseOpen):
			json.NewEncoder(response).Encode(err.ServiceError{Message: "Failed to run Delete: Database failed to open."})
		case errors.Is(err_delete, err.ErrDatabasePrepare):
			json.NewEncoder(response).Encode(err.ServiceError{Message: "Failed to run Delete: Database failed to prepare."})
		}
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
