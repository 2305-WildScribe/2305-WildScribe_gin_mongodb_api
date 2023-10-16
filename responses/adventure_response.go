package responses


type AdventureResponse struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Data    map[string]interface{} `json:"data"`
}

type GetAdventureResponse struct {
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

type AdventureError struct {
    Data struct {
        Error string `json:"error"`
        Attributes struct {
            User_id         string      `json:"user_id"`
            Adventure_id    string      `json:"adventure_id"`
        }   `json:"attributes"`
    }   `json:"data"`
}