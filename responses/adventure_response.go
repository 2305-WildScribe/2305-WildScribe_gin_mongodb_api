package responses

import "gin-mongo-api/models"


type AdventureResponse struct {
    Data    map[string]interface{} `json:"data"`
}

type GetAdventureResponse struct {
    Data struct {
        Type       string `json:"type" binding:"required"`
        Attributes models.Adventure `json:"attributes" binding:"required"`
    } `json:"data" binding:"required"`
}

type AdventureError struct {
    Data struct {
        Error string `json:"error"`
        Attributes struct {
            Adventure_id    string      `json:"adventure_id"`
        }   `json:"attributes"`
    }   `json:"data"`
}