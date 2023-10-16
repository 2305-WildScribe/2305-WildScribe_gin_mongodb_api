package responses

import "gin-mongo-api/models"


type AdventureResponse struct {

    Data struct {
        Type        string                  `json:"type" binding:"required"`
        Message     string                  `json:"message,omitempty"`
        Attributes  map[string]interface{}  `json:"attributes,omitempty"`
    }    `json:"data"`
}

type GetAdventureResponse struct {
    Data struct {
        Type       string `json:"type" binding:"required"`
        Attributes models.Adventure `json:"attributes" binding:"required"`
    } `json:"data" binding:"required"`
}

type AdventureErrorResponse struct {
    Data struct {
        Error string `json:"error"`
        Attributes  map[string]interface{}  `json:"attributes,omitempty"`
    }   `json:"data"`
}