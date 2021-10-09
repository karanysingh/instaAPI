package handlers

import (
    "fmt"
    "net/http"
    "strings"
	"strconv"
	"encoding/json"
    "log"
	L "instapi/helper"
)

var w http.ResponseWriter
var r *http.Request

func CreateUser(wr http.ResponseWriter, req *http.Request) {
	w = wr
	r = req
	var userdata []string
	if r.Method == "POST" {
		log.Println(r.Body)
		log.Println("Create a User")

		for _, values := range r.URL.Query() {
			userdata = append(userdata,values[0])
		}

		L.AddUser(11,userdata[0],userdata[1],userdata[2])
		fmt.Fprintf(w,"User Created")
	} else {
		log.Println("Method not implemented")
	}
}
func CreatePost(wr http.ResponseWriter, req *http.Request) {
	w = wr
	r = req
	var postdata []string
	if r.Method == "POST" {
		log.Println("Create a Post")
		for _, values := range r.URL.Query() {
			postdata = append(postdata,values[0])
		}
		userId,_ := strconv.Atoi(postdata[0])
		L.AddPost(11,userId,postdata[1],postdata[2],postdata[3])
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
		L.Showall(id)

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