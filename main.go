package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	loadEnvVariables()
	connectToDB()
	// handleRequests()

}

var coll *mongo.Collection

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil { // Attempt to load env, initialise an err variable and check if there is an error in 1 line.
		log.Println("No .env file found")
	} else {
		log.Println("Env variables loaded")
	}
}
func connectToDB() {
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
	defer func() { // delay execution of this anonymous function until the surrounding connectToDB function returns. Disconnects client instance.
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll = client.Database("VN").Collection("ChapterTest")
	// Declare & initialise a collection variable based on the client variable initialised above, using a database from that instance and a collection from that database
	// title := "Back to the Future"
	// var result bson.M
	// // Declare a result variable with a BSON M type
	// err = coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "title", Value: title}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the title %s\n", title)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)
	fmt.Println("----------------------------------------")
	HandleRequests()
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}
func AllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Desc: "Test description", Content: "Hello world!"},
		Article{Title: "Test title2", Desc: "Test description2", Content: "Hello world2!"},
	}
	fmt.Println("Endpoint hit: All articles")
	json.NewEncoder(w).Encode(articles)
}
func TestPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test post endpoint hit")
}
func GetAllScenes(w http.ResponseWriter, r *http.Request) {
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// var scenes []bson.M /*[]bson.D  */
	// if err = cursor.All(context.TODO(), &scenes); err != nil {
	// 	log.Fatal(err)
	// }
	// for _, scn := range scenes {
	// 	fmt.Println(scn["scene"])
	// }
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var scene bson.D
		if err = cursor.Decode(&scene); err != nil {
			log.Fatal(err)
		}
		fmt.Println(scene)
	}
}
func PostScene(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var s SceneObj
	err = json.Unmarshal(body, &s)
	if err != nil {
		panic(err)
	}
	log.Println(s.Scene)
}
