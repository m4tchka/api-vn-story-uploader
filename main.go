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
	client := ctl.ConnectToDB()
	fmt.Println("xxxxxxxxxxxx")
	router := rtr.HandleRequests()
	defer func() { // delay execution of this anonymous function until the surrounding connectToDB function returns. Disconnects client instance.
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	log.Fatal(http.ListenAndServe("127.0.0.1:"+os.Getenv("PORT"), router))
}
