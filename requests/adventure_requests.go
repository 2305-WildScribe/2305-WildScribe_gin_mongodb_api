package requests

// import (
//     "gopkg.in/go-playground/validator.v9"
// )
	

type DeleteAdventureRequest struct {
    Data struct {
        Type       string `json:"type" binding:"required"`
        Attributes struct {
            Adventure_id string `json:"adventure_id" binding:"required"`
        } `json:"attributes" binding:"required"`
    } `json:"data" binding:"required"`
}

type GetAdventureRequest struct {
    Data struct {
        Type       string `json:"type" binding:"required"`
        Attributes struct {
            Adventure_id string `json:"adventure_id" binding:"required"`
        } `json:"attributes" binding:"required"`
    } `json:"data" binding:"required"`
}

type AdventureRequest struct {
    Data struct {
        Type string `json:"type" binding:"required"`
        Attributes struct {
            User_id              string `json:"user_id" binding:"required"`
            Adventure_id         string `json:"adventure_id,omitempty"`
            Activity             string `json:"activity" binding:"required"`
            Date                 string `json:"date,omitempty"`
            Image_url            string `json:"image_url,omitempty"`
            Stress_level         string `json:"stress_level,omitempty"`
            Hours_slept          int    `json:"hours_slept,omitempty"`
            Sleep_stress_notes   string `json:"sleep_stress_notes,omitempty"`
            Hydration            string `json:"hydration,omitempty"`
            Diet                 int `json:"diet,omitempty"`
            Diet_hydration_notes string `json:"diet_hydration_notes,omitempty"`
            Beta_notes           string `json:"beta_notes,omitempty"`
            Lat                  float64 `json:"lat,omitempty"`
            Lon                  float64 `json:"lon,omitempty"`
        } `json:"attributes" binding:"required"`
    } `json:"data" binding:"required"`
}
type GetUserAdventureRequest struct {
	Data struct {
		Type			 string `json:"type" binding:"required"`
		Attributes struct {
			User_id 	 string `json:"user_id" binding:"required"`
		} `json:"attributes" binding:"required"`
	} `json:"data" binding:"required"`
}

// var validate = validator.New()

// func validateRequest(request interface{}) error {
// 	if err := validate.Struct(request); err != nil {
// 		return err
// 	}
// 	return nil
// }