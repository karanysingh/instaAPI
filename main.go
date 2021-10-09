package main

import (
    "fmt"
    "net/http"
	// "regexp"
    "strings"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"strconv"
	"encoding/json"

	"context"
    // "go.mongodb.org/mongo-driver/mongo/readpref"
    "log"
)
// var client *mongo.Client
// var ctx context.Context
var userCollection = new(mongo.Collection)
var postCollection = new(mongo.Collection)
// var ctx context.Context
var w http.ResponseWriter
var r *http.Request
 
func sayhelloName(wr http.ResponseWriter, req *http.Request) {
    r.ParseForm()  // parse arguments, you have to call this by yourself
	w = wr
	r = req
    fmt.Println(r.Form)  // print form information in server side
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}
type (
	// User represents the structure of our resource
	User struct {
		UserId int 			`json:"userid" bson:"userid"`
		Name   string        `json:"name" bson:"name"`
		Email string        `json:"email" bson:"email"`
		Password   string           `json:"password" bson:"password"`
	}
)

func getUser(userid int) []string {
	
	// coll := client.Database("instapi").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	user := User{}
	cur := userCollection.FindOne(ctx, bson.D{{"userid",userid}})
	cur.Decode(&user)
	log.Println(user)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	return user
}
func getPost(postid int) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	cur, err := postCollection.Find(ctx, bson.D{{"postid",postid}})
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		// do something with result....
		log.Println(result)
		// return result

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
func CreateUser(wr http.ResponseWriter, req *http.Request) {
	w = wr
	r = req
	if r.Method == "POST" {
		log.Println(r.Body)
		log.Println("Create a User")
		addUser(10,"name","mail@gmail.com","212")
		// log.Println("Create User",res)
		fmt.Fprintf(w,"User Created")
		// fmt.Fprintf(w,res)
	} else {
		log.Println("Method not implemented")
	}
}
func CreatePost(wr http.ResponseWriter, req *http.Request) {
	w = wr
	r = req
	if r.Method == "POST" {
		log.Println("Create a Post")
		addPost(5,"caption","imageurl","timstamp")
		// log.Println("Create Post",res)
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
		// fmt.Println("path", r.URL.Path)
		temp := strings.Trim(r.URL.Path, "users/")
		id,_ := strconv.Atoi(temp)
		log.Println(id)
		users := getUser(id)
		log.Println(users)
		j,err := json.Marshal(users)
		if err!=nil{
			fmt.Printf("Error: %s", err.Error())
		}
		// fmt.Fprintf(w,j)
		fmt.Println(j)

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
		log.Println(id)
		getPost(id)

	} else {
		log.Println("Method not implemented")
	}
}
func ShowAllPosts(wr http.ResponseWriter, req *http.Request) {
	w = wr
	r = req
	if r.Method == "GET" {
		log.Println("Show all posts of User")
		showall()

	} else {
		log.Println("Method not implemented")
	}
}
func addUser(userid int,name string,email string,password string) (*mongo.InsertOneResult){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	user := User{
		UserId:userid,
		Name:name,
		Email:email,
		Password:password,
	}
	result,err := userCollection.InsertOne(ctx,&user)
	if err != nil {
		log.Fatal(err)
		}
	log.Println(result)
	return result
}
func addPost(postid int,caption string,imageurl string,timestamp string) (*mongo.InsertOneResult){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// var result bson.M
	res,err := postCollection.InsertOne(ctx,bson.M{"postid":postid,"caption":caption,"imageurl":imageurl,"timestamp":timestamp})
	if err != nil {
		log.Fatal(err)
		}
	// result = bson.Decode(curr)
	log.Println(res)
	return res
}
func showall(){
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

		cur, err := userCollection.Find(ctx, bson.D{})
		if err != nil { log.Fatal(err) }
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result bson.D
			err := cur.Decode(&result)
			if err != nil { log.Fatal(err) }
			// do something with result....
			log.Println(result)
			// return result
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

}
func ConnectToDB(){
		
	clientOptions := options.Client().ApplyURI("mongodb+srv://karan:29LH3-WFvyV_sn-@cluster0.root9.mongodb.net/instapi?retryWrites=true&w=majority")
	// defer cancel()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	log.Println("Trying Connecting to mongodb \n")
	if err != nil {
	log.Fatal(err)
	}
	userCollection = client.Database("instapi").Collection("users")
	postCollection = client.Database("instapi").Collection("posts")

}

func main() {
	
	// CONNECTING TO MONGODB
	ConnectToDB()

	// ROUTES
	// http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/users", CreateUser)
	http.HandleFunc("/posts", CreatePost)
	http.HandleFunc("/users/", ShowUser)
	http.HandleFunc("/posts/", ShowPost)
	http.HandleFunc("/users/posts/:userid", ShowAllPosts)
	
	// SCHEMAS
	// FOLDER STRUCTURE
	// ERR HANDLING
	// DATABASE ENV VARIABLES SAFE
	// VERIFY USING REGEX
	// PAGINATION
	// TESTS
	// FRONTEND


	// INSERT OPERATIONS
	// addUser(client,ctx,4,"name","mail@gmail.com","212")
	// addPost(client,ctx,5,"caption","imageurl","timstamp")
	// showall(client,ctx)

	// GETTING OPERATION
	// getUser(client,ctx,2)
	// getPost(client,ctx,5)

	// getAllposts
	// LISTENER
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}