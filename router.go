package main

import (
	"log"
	"net/http"

	// "vn-story-uploader/models"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/articles", AllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", TestPostArticles).Methods("POST")
	// ----------------------------------------------------------------
	myRouter.HandleFunc("/scenes" /* models. */, GetAllScenes).Methods("GET")
	myRouter.HandleFunc("/scenes", PostScene).Methods("POST")
	myRouter.HandleFunc("/scenes/{id}", DeleteScene).Methods("DELETE")
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", myRouter))
}
