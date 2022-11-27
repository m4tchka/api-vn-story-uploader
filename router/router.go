package router

import (
	"log"

	ctl "vn-story-uploader/controller"

	"github.com/gorilla/mux"
)

func HandleRequests() *mux.Router {
	log.Println("Server running ...")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", ctl.HomePage)
	//TODO: TEST ROUTES ONLY -----------
	myRouter.HandleFunc("/articles", ctl.AllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", ctl.TestPostArticles).Methods("POST")
	// --------------------------
	myRouter.HandleFunc("/scenes/{id}" /* models. */, ctl.GetSpecificScene).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/scenes" /* models. */, ctl.GetAllScenes).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/scenes", ctl.PostScene).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/scenes/{id}", ctl.DeleteScene).Methods("DELETE", "OPTIONS")
	return myRouter

}
