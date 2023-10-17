package controllers

import (
	"context"
	// "gin-mongo-api/collections"
	"gin-mongo-api/requests"

	// "gin-mongo-api/authentication"
	"gin-mongo-api/configs"
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
    // "encoding/json"
)

var validateAdventure = validator.New()
var adventureCollection *mongo.Collection = configs.GetCollection(configs.DB, "adventures")

// func SetAdventureCollection(collection collections.AdventureCollection) {
//     adventureCollection = collection
// }

func userExists(ctx context.Context, userID string) bool {
    // Checks if a valid user id before going to db
    objID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return false
    }
    // Counts to see if there is a user with that id
    count, _ := userCollection.CountDocuments(ctx, bson.D{{Key:"_id",Value:objID}})
    return count > 0
}

func CreateAdventure() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        // Sets a error response struct
        var error_response responses.AdventureErrorResponse
        // Sets a request struct
        var requestBody requests.AdventureRequest
        // Sets a response struct
        var response responses.AdventureResponse
        response.Data.Type = "adventure"
        // Binds http response to request struct and checks for required fields
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            // requestBodyJSON, _ := json.Marshal(requestBody)
            // error_response.Data.Error = string(requestBodyJSON)
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
        filter := bson.D{{Key:"_id",Value:objId}}
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
        var response responses.GetAnAdventureResponse
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
        // Sets Filter adventure by objid
        filter := bson.D{{Key:"_id",Value:objId}}
        // Finds adventure in collection
        result := adventureCollection.FindOne(ctx, filter).Decode(&adventure)
        // Returns 404 if Adventure not found
		if result != nil {   
            error_response.Data.Error = "Invalid adventure ID"
            error_response.Data.Attributes = map[string]interface{}{"adventure_id": filter }
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
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        var requestBody requests.GetUserAdventureRequest

        if err := c.ShouldBindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, responses.AdventureErrorResponse{
                Data: struct {
                    Error      string                 `json:"error"`
                    Attributes map[string]interface{} `json:"attributes,omitempty"`
                }{
                    Error: "Invalid Request",
                },
            })
            return
        }

        userID := requestBody.Data.Attributes.User_id

        // Create a filter to find documents with the specified user_id
        filter := bson.D{{Key:"user_id",Value:userID}}

        // Find documents that match the filter
        cursor, err := adventureCollection.Find(ctx, filter)
        if err != nil {
            // Handle the error and return an error response
            c.JSON(http.StatusInternalServerError, responses.AdventureErrorResponse{
                Data: struct {
                    Error      string                 `json:"error"`
                    Attributes map[string]interface{} `json:"attributes,omitempty"`
                }{
                    Error: "Internal Server Error",
                },
            })
            return
        }
        defer cursor.Close(ctx)

        var adventures []models.Adventure

        if err := cursor.All(ctx, &adventures); err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdventureErrorResponse{
                Data: struct {
                    Error      string                 `json:"error"`
                    Attributes map[string]interface{} `json:"attributes,omitempty"`
                }{
                    Error: "Internal Server Error",
                },
            })
            return
        }

        response := responses.GetAdventureResponse{
            Data: struct {
                Type       string               `json:"type" binding:"required"`
                Attributes []models.Adventure  `json:"attributes" binding:"required"`
            }{
                Type:       "adventures",
                Attributes: adventures,
            },
        }

        c.JSON(http.StatusOK, response)
    }
}

func UpdateAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			var error_response responses.AdventureErrorResponse
			var requestBody requests.AdventureRequest
			if err := c.ShouldBindJSON(&requestBody); err != nil {
					error_response.Data.Error = "Invalid Request"
					c.JSON(http.StatusBadRequest, error_response)
					return
			}

			// Getting the adventureId from the request's path, may have to change if we want ID to be in body of request
			adventureId := requestBody.Data.Attributes.Adventure_id
			objId, err := primitive.ObjectIDFromHex(adventureId)
			if err != nil {
					error_response.Data.Error = "Invalid adventure ID"
					error_response.Data.Attributes = map[string]interface{}{"adventure_id": adventureId}
					c.JSON(http.StatusBadRequest, error_response)
					return
			}

			update := bson.M{
					"$set": bson.M{
							"activity":             requestBody.Data.Attributes.Activity,
							"date":                 requestBody.Data.Attributes.Date,
							"image_url":            requestBody.Data.Attributes.Image_url,
							"stress_level":         requestBody.Data.Attributes.Stress_level,
							"hours_slept":          requestBody.Data.Attributes.Hours_slept,
							"sleep_stress_notes":   requestBody.Data.Attributes.Sleep_stress_notes,
							"hydration":            requestBody.Data.Attributes.Hydration,
							"diet":                 requestBody.Data.Attributes.Diet,
							"diet_hydration_notes": requestBody.Data.Attributes.Diet_hydration_notes,
							"beta_notes":           requestBody.Data.Attributes.Beta_notes,
					},
			}

			result, err := adventureCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
			if err != nil || result.ModifiedCount == 0 {
					error_response.Data.Error = "Adventure not updated"
					error_response.Data.Attributes = map[string]interface{}{"adventure_id": adventureId}
					c.JSON(http.StatusInternalServerError, error_response)
					return
			}

			var response responses.AdventureResponse
			response.Data.Type = "adventure"
			response.Data.Message = "success"
			c.JSON(http.StatusOK, response)
	}
}
