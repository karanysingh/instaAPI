package handlers

import (
    "fmt"
    "net/http"
    "strings"
	"strconv"
	"encoding/json"
	"reflect"
    "log"
	models "instapi/models"

	L "instapi/helper"
)

var w http.ResponseWriter
var r *http.Request

func CreateUser(wr http.ResponseWriter, req *http.Request) {
	w = wr
	r = req
	// var userdata []string
	u := models.User{}
	if r.Method == "POST" {
		// log.Println(r.Body)
		log.Println("Create a User")

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		log.Println(u)
		L.AddUser(u.UserId,u.Name,u.Email,u.Password)
		log.Println(reflect.TypeOf(u))
		log.Println(u.Name)
		fmt.Fprintf(w,"User Created")
	} else {
		log.Println("Method not implemented")
	}
}
func CreatePost(wr http.ResponseWriter, req *http.Request) {
	w = wr
	r = req
	post := models.Post{}
	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		log.Println(post)
		L.AddPost(post.PostId,post.UserId,post.Caption,post.Imageurl,post.Timestamp)

		fmt.Fprintf(w,"Post Created")
		
	} else {
		log.Println("Method not implemented")
	}
}

func ShowUser(wr http.ResponseWriter, req *http.Request) {
	w = wr
	r = req
	if r.Method == "GET" {
		log.Println("Show User")
		temp := strings.Trim(r.URL.Path, "users/")
		id,_ := strconv.Atoi(temp)

		user := L.GetUser(id)

		w.Header().Set("Content-Type", "text/json")
		userres, _ := json.Marshal(user)

		fmt.Fprintf(w,string(userres))

	} else {
		log.Println("Method not implemented")
	}
}
func ShowPost(wr http.ResponseWriter, req *http.Request) {
	w = wr
	r = req
	if r.Method == "GET" {
		log.Println("Show Post")
		temp := strings.Trim(r.URL.Path, "posts/")
		id,_ := strconv.Atoi(temp)
		
		post := L.GetPost(id)

		w.Header().Set("Content-Type", "text/json")
		postres, _ := json.Marshal(post)
		resString := string(postres)
		fmt.Fprintf(w,resString)

	} else {
		log.Println("Method not implemented")
	}
}
func ShowAllPosts(wr http.ResponseWriter, req *http.Request) {
	w = wr
	r = req
	if r.Method == "GET" {
		log.Println("Show all posts of a User")
		temp := strings.Trim(r.URL.Path, "posts/users/")
		id,_ := strconv.Atoi(temp)
		log.Println(id)
		posts := L.Showall(id)
		w.Header().Set("Content-Type", "text/json")
		postres, _ := json.Marshal(posts)
		resString := string(postres)
		fmt.Fprintf(w,resString)
	} else {
		log.Println("Method not implemented")
	}
}

func SayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // parse arguments, you have to call this by yourself
    fmt.Println(r.Form)  // print form information in server side
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "This is Instapi") // send data to client side
}