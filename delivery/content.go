package delivery

import (
	"net/http"
	"encoding/json"
	"../model"
)

func GetContent(w http.ResponseWriter, r *http.Request) {

	contents := model.Contents{
		model.Content{ID: 1, Title: "Title 1", Url: "http://example.com", Author: "Sergio", Name: "content_1"},
		model.Content{ID: 2, Title: "Title 2", Url: "http://example.com", Author: "Sergio", Name: "content_2"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(contents); err != nil {
		panic(err)
	}

}