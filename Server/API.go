package Server

import (
	"encoding/json"
	"net/http"
)

type Users struct {
	ID    int `json:"id"`
	Text  int `json:"text"`
	Title int `json:"title"`
}

var (
	users []Users
)

func Getdata(response http.ResponseWriter, r http.Request) {

	response.Header().Set("Content-Type", "application/json")
	result6, err := json.Marshal(users)
	if err != nil {

	}
	response.Write(result6)
}
