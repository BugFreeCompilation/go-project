package service

import (
	database "BugFreeCompilation/go-project/database"
	"BugFreeCompilation/go-project/entity"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test will pass, you can make it fail if you rename the variable
const (
	TEXT string = "Golang the Movie"
)

var (
	mydb_test database.Database = database.InitSQLite3()
	myservice Service           = InitService(mydb_test)
)

func TestAdd(t *testing.T) {
	// https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
	var jsonStr = []byte(`{"TITLE": "Golang the Movie"}`)

	// https://stackoverflow.com/questions/40680498/how-to-test-http-request-handlers
	request, _ := http.NewRequest("POST", "/data", bytes.NewBuffer(jsonStr))

	response := httptest.NewRecorder()

	myservice.Add(response, request)

	// the return should equal the input
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	assert.Equal(t, TEXT, post.TITLE)
}
