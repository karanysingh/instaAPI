package main

import (
	handlers "instapi/handlers"
	L "instapi/helper"
	"log"
	"net/http"
	"os"
)

var w http.ResponseWriter
var r *http.Request

func main() {

	// CONNECTING TO MONGODB
	L.ConnectToDB()

	// ROUTES
	http.HandleFunc("/", handlers.SayhelloName)
	http.HandleFunc("/users", handlers.CreateUser)
	http.HandleFunc("/posts", handlers.CreatePost)
	http.HandleFunc("/users/", handlers.ShowUser)
	http.HandleFunc("/posts/", handlers.ShowPost)
	http.HandleFunc("/users/posts/", handlers.ShowAllPosts)

	// LISTENER
	port := os.Getenv("PORT")
	// if err != nil {
	// 	port = "9090"
	// }
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
