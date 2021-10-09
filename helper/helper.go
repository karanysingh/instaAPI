package helper

import (
    "net/http"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
    "context"
    "log"
    models "instapi/models"
)

var userCollection = new(mongo.Collection)
var postCollection = new(mongo.Collection)


var w http.ResponseWriter
var r *http.Request


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

func GetUser(userid int) models.User {
	
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	user := models.User{}
	cur := userCollection.FindOne(ctx, bson.D{{"userid",userid}})
	
	cur.Decode(&user)
	return user

}

func GetPost(postid int) models.Post {
    ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	post := models.Post{}
	cur := postCollection.FindOne(ctx, bson.D{{"postid",int32(postid)}})
	
	cur.Decode(&post)
	// cur.NeDecode(&post)
	// log.Println(post)
	return post
}
func AddUser(userid int,name string,email string,password string) (*mongo.InsertOneResult){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	user := models.User{
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
func AddPost(postid int,userid int,caption string,imageurl string,timestamp string) (*mongo.InsertOneResult){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// var result bson.M
	post := models.Post{
		PostId:postid,
		UserId:userid,
		Caption:caption,
		Imageurl:imageurl,
		Timestamp:timestamp,
	}
	res,err := postCollection.InsertOne(ctx,&post)
	if err != nil {
		log.Fatal(err)
		}
	// result = bson.Decode(curr)
	log.Println(res)
	return res
}

func Showall(userid int){
	cur, err := postCollection.Find(context.Background(), bson.D{{"postid",11}})
if err != nil { log.Fatal(err) }
defer cur.Close(context.Background())
post := models.Post{}
for cur.Next(context.Background()) {
  // To decode into a struct, use cursor.Decode()
  err := cur.Decode(&post)
  if err != nil { log.Fatal(err) }
  log.Println(post)
}
	// return post
}