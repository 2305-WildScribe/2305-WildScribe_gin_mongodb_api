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
        Type       string `json:"type" binding:"required"`
        Attributes struct {
            User_id 		        string 				`json:"user_id" binding:"required"`
            Activity  		        string              `json:"activity" binding:"required"`
            Date 			        string              `json:"date,omitempty"`
            Image_url 		        string              `json:"image_url,omitempty"`
            Stress_level 	        string              `json:"stress_level,omitempty"`
            Hours_slept             int                 `json:"hours_slept,omitempty"`
            Sleep_stress_notes      string              `json:"sleep_stress_notes,omitemtpy"`
            Hydration               string                 `json:"hydration,omitempty"`
            Diet                    string              `json:"diet,omitemtpy"`
            Diet_Hydration_notes    string              `json:"diet_and_hydration_notes,omitempty"`
            Beta_notes  	        string				`json:"beta_notes,omitempty"`
        } `json:"attributes" binding:"required"`
    } `json:"data" binding:"required"`
}

type GetUserAdventureRequest struct {
	Data struct {
		Type			 string `json:"type"`
		Attributes struct {
			User_id 	 string `json:"user_id"`
		} `json:"attributes"`
	} `json:"data"`
}