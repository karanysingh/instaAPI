
package models
type (
	User struct {
		UserId int 			
		Name   string       
		Email string        
		Password   string           
	}
)
type (
	Post struct {
		PostId int 
		UserId int			
		Caption   string       
		Imageurl string        
		Timestamp   string         
	}
)