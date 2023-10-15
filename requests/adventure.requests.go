package requests
	

type DeleteAdventureRequest struct {
    Data struct {
        Type       string `json:"type"`
        Attributes struct {
            User_id int `json:"user_id"`
            Adventure_id string `json:"adventure_id"`
        } `json:"attributes"`
    } `json:"data"`
}

type GetAdventureRequest struct {
    Data struct {
        Type       string `json:"type"`
        Attributes struct {
            Adventure_id string `json:"adventure_id"`
        } `json:"attributes"`
    } `json:"data"`
}

type CreateAdventureRequest struct {
    Data struct {
        Type       string `json:"type"`
        Attributes struct {
            User_id 		int 					`json:"user_id"`
            Activity  		string                	`json:"activity"`
            Date 			string                 	`json:"date,omitempty"`
            Notes  			string					`json:"notes,omitempty"`
            Image_url 		string                 	`json:"image_url,omitempty"`
            Stress_level 	string                	`json:"stress_level,omitempty"`
            Hydration 		int                 	`json:"hydration,omitempty"`
            Diet 			string                 	`json:"diet,omitempty"`
        } `json:"attributes"`
    } `json:"data"`
}
