package models

// func HomePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Homepage endpoint hit")
// }
// func AllArticles(w http.ResponseWriter, r *http.Request) {
// 	articles := Articles{
// 		Article{Title: "Test title", Desc: "Test description", Content: "Hello world!"},
// 		Article{Title: "Test title2", Desc: "Test description2", Content: "Hello world2!"},
// 	}
// 	fmt.Println("Endpoint hit: All articles")
// 	json.NewEncoder(w).Encode(articles)
// }
// func TestPostArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Test post endpoint hit")

// }
// func GetAllScenes(w http.ResponseWriter, r *http.Request) {
// 	cursor, err := coll.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// var scenes []bson.M /*[]bson.D  */
// 	// if err = cursor.All(context.TODO(), &scenes); err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// for _, scn := range scenes {
// 	// 	fmt.Println(scn["scene"])
// 	// }
// 	defer cursor.Close(context.TODO())
// 	for cursor.Next(context.TODO()) {
// 		var scene bson.D
// 		if err = cursor.Decode(&scene); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(scene)
// 	}
// }
