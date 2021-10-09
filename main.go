package main

import (
    "net/http"
    "log"
	L "instapi/helper"
	handlers "instapi/handlers"
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
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}