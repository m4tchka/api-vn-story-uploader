package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	m "vn-story-uploader/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Coll *mongo.Collection

func ConnectToDB() *mongo.Client {
	if err := godotenv.Load(); err != nil { // Attempt to load env, initialise an err variable and check if there is an error in 1 line.
		log.Println("No .env file found")
	} else {
		log.Println("Env variables found")
	}
	uri := os.Getenv("MONGODB_URI")
	// Get the env variable with the key and initialise the uri variable with it.
	if uri == "" {
		// Check that there is indeed a env variable with that name
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	// Create a client instance and connect to it with options defined in the uri (ex. the database name, user name, passwrod etc). Initialise it to a client variable
	if err != nil {
		panic(err)
	}
	/* 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	   	err = client.Connect(ctx)
	   	if err != nil {
	   		panic(err)
	   	} */
	Coll = client.Database("VN").Collection("ChapterTest")
	// Declare & initialise a collection variable based on the client variable initialised above, using a database from that instance and a collection from that database
	return client // Return the connected client instance

}
func HomePage(w http.ResponseWriter, r *http.Request) { // Homepage route that simply prints a string.
	fmt.Fprintf(w, "Homepage endpoint hit")
}
func AllArticles(w http.ResponseWriter, r *http.Request) { // Test endpoint that returns a list of articles
	articles := m.Articles{
		m.Article{Title: "Test title", Desc: "Test description", Content: "Hello world!"},
		m.Article{Title: "Test title2", Desc: "Test description2", Content: "Hello world2!"},
	}
	log.Println("Endpoint hit: All articles")
	w.Header().Set("Content-Type", "application/json")             // Indicates that the request body format is JSON.
	w.Header().Set("Access-Control-Allow-Origin", "*")             // Indicates that code from any origin is allowed to access this endpoint.
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Indicates that any request must have the header of "Content-Type"
	//https://stackoverflow.com/questions/39507065/enable-cors-in-golang
	w.WriteHeader(http.StatusOK)        // Give the response a header of "200 OK"
	json.NewEncoder(w).Encode(articles) // Write to the body of the respnose, the articles defined above.
}

//	func TestPostArticles(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Test post endpoint hit")
//	}
func GetSpecificScene(w http.ResponseWriter, r *http.Request) {
	log.Println("------------- GET SPECIFIC SCENE FUNC TRIGGERED -----------------")
	specificSceneId, err := strconv.Atoi(mux.Vars(r)["id"]) //Take the variable from the parameter of the request and store it in specificSceneId
	if err != nil {
		panic(err)
	}
	var scene bson.M // Declare the variable scene, of type bson.M
	// bson.M is a MongoDB-specific type that is useful when the specific order of keys in an object is NOT important. If it did matter, bson.D would be more suitable.
	if err := Coll.FindOne(context.Background(), bson.M{"id": specificSceneId}).Decode(&scene); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Specific scene:", scene)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(scene)
}
func GetAllScenes(w http.ResponseWriter, r *http.Request) {
	log.Println("------------- GET ALL SCENES FUNC TRIGGERED -----------------")
	cursor, err := Coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var scenes []bson.M // Must not be bson.D
	if err = cursor.All(context.TODO(), &scenes); err != nil {
		log.Fatal(err)
	}
	for _, scn := range scenes {
		fmt.Println("Scene printed >>>", scn)
	}
	defer cursor.Close(context.TODO())
	/*
		for cursor.Next(context.TODO()) {
			var scene bson.M
			if err = cursor.Decode(&scene); err != nil {
				log.Fatal(err)
			}
			fmt.Println(scene)
		}
	*/
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(scenes)

}
func PostScene(w http.ResponseWriter, r *http.Request) {
	log.Println("------------- POST SCENE FUNC TRIGGERED -----------------")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// log.Println(string(body))
	var s m.SceneObj
	err = json.Unmarshal(body, &s)
	if err != nil {
		panic(err)
	}
	log.Println(s.Scene)
	insertAScene(s)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(s)
}
func insertAScene(s m.SceneObj) {
	insOneRes, err := Coll.InsertOne(context.Background(), s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully inserted:", insOneRes)
}
func DeleteScene(w http.ResponseWriter, r *http.Request) {
	log.Println("------------- DELETE SCENE FUNC TRIGGERED -----------------")
	idToBeDeleted, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	fmt.Printf("id: %d of type %T\n", idToBeDeleted, idToBeDeleted)
	deleteAScene(idToBeDeleted)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
func deleteAScene(id int) {
	var scene bson.M
	if err := Coll.FindOne(context.Background(), bson.M{"id": id}).Decode(&scene); err != nil {
		log.Fatal(err)
	}
	fmt.Println("scene:", scene)
}
