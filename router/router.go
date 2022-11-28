package router

import (
	"log"

	ctl "vn-story-uploader/controller"

	"github.com/gorilla/mux"
)

func HandleRequests() *mux.Router {
	log.Println("Server running ...")
	myRouter := mux.NewRouter().StrictSlash(true) // Define a new router from mux (and allow trailing slashes for routes)
	//TODO: TEST ROUTES ONLY -----------
	myRouter.HandleFunc("/", ctl.HomePage)                           // Route to the home page
	myRouter.HandleFunc("/articles", ctl.AllArticles).Methods("GET") // Test route that displays some article structs
	// --------------------------
	myRouter.HandleFunc("/scenes/{id}", ctl.GetSpecificScene).Methods("GET", "OPTIONS") // Route that takes an id parameter and returns the scene with that id
	myRouter.HandleFunc("/scenes", ctl.GetAllScenes).Methods("GET", "OPTIONS")          // Route that returns all scenes within the collection
	myRouter.HandleFunc("/scenes", ctl.PostScene).Methods("POST", "OPTIONS")            // Route that posts a new scene to the collection
	myRouter.HandleFunc("/scenes/{id}", ctl.DeleteScene).Methods("DELETE", "OPTIONS")   // Route that deletes a single scene whose id matches the parameter from the collection
	return myRouter                                                                     // Return this router so the function can be called and the router served from main.go

}
