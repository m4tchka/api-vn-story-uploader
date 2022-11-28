package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	ctl "vn-story-uploader/controller"
	rtr "vn-story-uploader/router"
)

func main() {
	client := ctl.ConnectToDB() // Call this function from controller.go and store  the returned, newly connected MongoDB client
	fmt.Println("xxxxxxxxxxxx")
	router := rtr.HandleRequests() // Call this function from router.go and store the newly set up router instance
	defer func() {                 // Delay execution of this anonymous function until the surrounding connectToDB function returns. Disconnects client instance.
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router)) // Route requests to the specifies port through the router instance from above. If the router instance disconnects, ListenAndServe will return an error and log.Fatal will stop the API.
}
