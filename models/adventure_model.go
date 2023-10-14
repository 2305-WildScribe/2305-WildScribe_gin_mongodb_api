package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Adventure struct {
	User_id 		int 					`json:"user_id"`
	Activity  		string                	`json:"activity"`
    Date 			string                 	`json:"date,omitempty"`
    Notes  			string					`json:"notes,omitempty"`
    Image_url 		string                 	`json:"image_url,omitempty"`
    Stress_level 	string                	`json:"stress_level,omitempty"`
    Hydration 		int                 	`json:"hydration,omitempty"`
    Diet 			string                 	`json:"diet,omitempty"`
}