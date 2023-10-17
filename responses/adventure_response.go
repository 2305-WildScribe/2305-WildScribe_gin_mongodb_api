package responses

import "gin-mongo-api/models"

// "gin-mongo-api/models"
// "gin-mongo-api/responses"


type AdventureResponse struct {

    Data struct {
        Type        string                  `json:"type" binding:"required"`
        Message     string                  `json:"message,omitempty"`
        Attributes  map[string]interface{}  `json:"attributes,omitempty"`
    }    `json:"data"`
}

type GetAdventureResponse struct {
    Data struct {
        Type       string              `json:"type" binding:"required"`
        Attributes []models.Adventure   `json:"attributes" binding:"required"`
    } `json:"data" binding:"required"`
}

type GetAnAdventureResponse struct {
    Data struct {
        Type       string              `json:"type" binding:"required"`
        Attributes models.Adventure  `json:"attributes" binding:"required"`
    } `json:"data" binding:"required"`
}

// type UserAdventureResponse struct {
//     User_id              string `json:"user_id" binding:"required"`
//     Adventure_id         string `json:"adventure_id" binding:"required"`
//     Activity             string `json:"activity" binding:"required"`
//     Date                 string `json:"date,omitempty"`
//     Image_url            string `json:"image_url,omitempty"`
//     Stress_level         string `json:"stress_level,omitempty"`
//     Hours_slept          int    `json:"hours_slept,omitempty"`
//     Sleep_stress_notes   string `json:"sleep_stress_notes,omitempty"`
//     Hydration            string `json:"hydration,omitempty"`
//     Diet                 string `json:"diet,omitempty"`
//     Diet_hydration_notes string `json:"diet_hydration_notes,omitempty"`
//     Beta_notes           string `json:"beta_notes,omitempty"`
// }

type AdventureErrorResponse struct {
    Data struct {
        Error string `json:"error"`
        Attributes  map[string]interface{}  `json:"attributes,omitempty"`
    }   `json:"data"`
}