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
	
	// SCHEMAS
	// FOLDER STRUCTURE
	// ERR HANDLING
	// DATABASE ENV VARIABLES SAFE
	// VERIFY USING REGEX
	// PAGINATION
	// TESTS
	// FRONTEND

	// LISTENER
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}