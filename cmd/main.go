package main

import (
	"log"
	"net/http"
	"testoviy-zadaniya/photo-uploader/cmd/handler"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	h := handler.New()

	r.HandleFunc("/upload", h.Photo.Upload).Methods("POST")

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal(err)
	}
}
