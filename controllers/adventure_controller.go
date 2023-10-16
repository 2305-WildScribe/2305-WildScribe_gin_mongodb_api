package controllers

import (
	"context"
	"gin-mongo-api/collections"
	"gin-mongo-api/requests"

	// "gin-mongo-api/authentication"
	// "gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"gin-mongo-api/serializers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo"
)

var validateAdventure = validator.New()


func SetAdventureCollection(collection collections.AdventureCollection) {
    adventureCollection = collection
}
func userExists(ctx context.Context, userID string) bool {
    // Checks if a valid user id before going to db
    objID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return false
    }
    // counts to see if there is a user with that id
    count, _ := userCollection.CountDocuments(ctx, bson.M{"_id": objID})
    return count > 0
}

func CreateAdventure() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        // Sets a error response struct
        var error_response responses.AdventureErrorResponse
        // Sets a request struct
        var requestBody requests.CreateAdventureRequest
        // Sets a response struct
        var response responses.AdventureResponse
        response.Data.Type = "adventure"
        // Binds http response to request struct and checks for required fields
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            error_response.Data.Error = "Invalid Request"
            c.JSON(http.StatusBadRequest, error_response)
            return
        }
        // Gets User ID
        userID := requestBody.Data.Attributes.User_id
        // Checks if user exist
        if userExists(ctx, userID) == false {
            error_response.Data.Error = "Invalid user ID"
            error_response.Data.Attributes = map[string]interface{}{"user_id": userID }
            c.JSON(http.StatusBadRequest, error_response)
            return
        }
        // If all validations pass then serializes the request
        adventure := serializers.SerializeCreateAdventureRequest(requestBody)
        // Inserts the serialized model into the db
        result, _ := adventureCollection.InsertOne(ctx, adventure)
        // Sets a resposne struct and returns 201 and success if created
        response.Data.Message = "success"
        response.Data.Attributes =  map[string]interface{}{"adventure_id": result.InsertedID }
        c.JSON(http.StatusCreated, response)
    }
}

func DeleteAdventure() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        // Sets error response struct
        var error_response responses.AdventureErrorResponse
        // Sets requestBody struct
        var requestBody requests.DeleteAdventureRequest
        // Binds the request json to requestBody
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            error_response.Data.Error = "Invalid Request"
            c.JSON(http.StatusBadRequest, error_response)
            return
        }
        // Accesses the nested data
        adventureId := requestBody.Data.Attributes.Adventure_id
        // Sets the object id
        objId, err := primitive.ObjectIDFromHex(adventureId);
        if err != nil {
            error_response.Data.Error = "Invalid adventure ID"
            error_response.Data.Attributes = map[string]interface{}{"adventure_id": adventureId }
            c.JSON(http.StatusBadRequest, error_response)
            return
        }
        // Sets the filter
        filter := bson.M{"_id": objId}
        // Delete the object from the collection
        result, err := adventureCollection.DeleteOne(ctx, filter)
        if result.DeletedCount == 0 {
            error_response.Data.Error = "Invalid adventure ID"
            error_response.Data.Attributes = map[string]interface{}{"adventure_id": adventureId }
            c.JSON(http.StatusBadRequest, error_response)
            return
        }
        // If object deleted return success
        var response responses.AdventureResponse
        response.Data.Type = "adventure"
        response.Data.Message = "success"
        c.JSON(http.StatusOK, response)
    }
}

func GetAnAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
        // Set Request Body
		var requestBody requests.GetAdventureRequest
        // Set Response Defaults
        var response responses.GetAdventureResponse
        response.Data.Type = "adventure"
        // Set Response Error
        var error_response responses.AdventureErrorResponse
        // Binds request json to requestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
            error_response.Data.Error = "Invalid Request"
            c.JSON(http.StatusBadRequest, error_response)
			return
		}
        // Gets Adventure ID and sets it as a var
		adventureId := requestBody.Data.Attributes.Adventure_id
        // Check if valid ID
		objId, err := primitive.ObjectIDFromHex(adventureId)
        // Returns Bad Request if invalid format
		if err != nil {
            error_response.Data.Error = "Invalid adventure ID"
            error_response.Data.Attributes = map[string]interface{}{"adventure_id": adventureId }
            c.JSON(http.StatusBadRequest, error_response)
			return
		}
        // Set model type for find
        var adventure models.Adventure
        // Find adventure by objid
        filter := bson.M{"_id": objId}

		result := adventureCollection.FindOne(ctx,filter).Decode(&adventure)
        // Returns 404 if Adventure not found
		if result == nil {   
            error_response.Data.Error = "Invalid adventure ID"
            error_response.Data.Attributes = map[string]interface{}{"adventure_id": adventureId }
            c.JSON(http.StatusNotFound, error_response)
			return
		}
        // Set Response 
        response.Data.Attributes = adventure
        response.Data.Type = "adventure"
        // Send JSON
		c.JSON(http.StatusOK, response)
	}
}

func GetAdventuresForUser() gin.HandlerFunc {
    return func(c *gin.Context) {
    //     ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    //     defer cancel()
    //     var requestBody requests.GetUserAdventureRequest

    //     userID := requestBody.Data.Attributes.User_id

    //     var adventures []models.Adventure

    //     // cursor, _ := adventureCollection.Find(ctx, bson.M{"user_id": userID})
    //     if err := c.ShouldBindJSON(&requestBody); err != nil {
    //         // c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Data: map[string]interface{}{"error": err.Error()}})
    //         return
    //     }
    //     defer cursor.Close(ctx)

    //     for cursor.Next(ctx) {
    //         var adventure models.Adventure
    //         if err := cursor.Decode(&adventure); err != nil {
    //             // c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Data: map[string]interface{}{"error": err.Error()}})
    //             return
    //         }
    //         adventures = append(adventures, adventure)
    //     }

    //     if err := cursor.Err(); err != nil {
    //         // c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Data: map[string]interface{}{"error": err.Error()}})
    //         return
    //     }

    //     // c.JSON(http.StatusOK, responses.AdventureResponse{Data: map[string]interface{}{"type":"adventures" , "attributes": adventures}})
    }
}