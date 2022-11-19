package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Desc: "Test description", Content: "Hello world!"},
	}
	fmt.Println("Endpoint hit: All articles")
	json.NewEncoder(w).Encode(articles)
}
func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test post endpoint hit")

}
