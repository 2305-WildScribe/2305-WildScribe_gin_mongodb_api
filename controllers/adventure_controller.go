package controllers

import (
    "context"
    "gin-mongo-api/authentication"
    "gin-mongo-api/configs"
    "gin-mongo-api/models"
    "gin-mongo-api/responses"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

var adventureCollection *mongo.Collection = configs.GetCollection(configs.DB, "adventures")
var validateAdventure = validator.New()

func CreateAdventure() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var adventure models.Adventure
        defer cancel()

        //validate the request body
        if err := c.BindJSON(&adventure); err != nil {
            c.JSON(http.StatusBadRequest, responses.AdventureResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := validateAdventure.Struct(&adventure); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.AdventureResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        result, err := adventureCollection.InsertOne(ctx, adventure)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.AdventureResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}


func GetAnAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		adventureId, ok := requestBody["adventureId"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "adventureId is required and should be a string"})
			return
		}

		objId, err := primitive.ObjectIDFromHex(adventureId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid adventure ID format"})
			return
		}

		var adventure models.Adventure
		err = adventureCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&adventure)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdventureResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": adventure}})
	}
}

func GetAdventuresForUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        userID := authentication.GetAuthenticatedUserID(c)

        var adventures []models.Adventure

        cursor, err := adventureCollection.Find(ctx, bson.M{"userId": userID})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
        defer cursor.Close(ctx)

        for cursor.Next(ctx) {
            var adventure models.Adventure
            if err := cursor.Decode(&adventure); err != nil {
                c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
                return
            }
            adventures = append(adventures, adventure)
        }

        if err := cursor.Err(); err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK, responses.AdventureResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": adventures}})
    }
}