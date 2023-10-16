package controllers

import (
    "context"
    "gin-mongo-api/requests"
    // "gin-mongo-api/authentication"
    "gin-mongo-api/configs"
    "gin-mongo-api/models"
    "gin-mongo-api/responses"
    "gin-mongo-api/serializers"
    "net/http"
    "time"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

var adventureCollection *mongo.Collection = configs.GetCollection(configs.DB, "adventures")
var validateAdventure = validator.New()

func userExists(ctx context.Context, userID string) (bool) {
    objId, _ := primitive.ObjectIDFromHex(userID) 
    count, _ := userCollection.CountDocuments(ctx,  bson.M{"_id": objId})
    fmt.Printf("Filter: %v\n", objId)
    return count > 0
}

func CreateAdventure() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        // Sets a request struct
        var requestBody requests.CreateAdventureRequest
        // Binds response to request struct and checks for required fields
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }
        // Check if User_id is a valid user
        userID := requestBody.Data.Attributes.User_id
        userExists := userExists(ctx, userID)
        // Sends error if user dosen't exist
        if userExists == false {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }
        // If all validations pass then serializes the request and assings it
        adventure := serializers.SerializeCreateAdventureRequest(requestBody)
        // Inserts the serialized model into the db
        result, err := adventureCollection.InsertOne(ctx, adventure)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
            return
        }
        //Returns 201 and success when created
        c.JSON(http.StatusCreated, responses.AdventureResponse{Data: map[string]interface{}{"message": "success", "adventure_id": result.InsertedID }})
    }
}

func DeleteAdventure() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        var requestBody requests.DeleteAdventureRequest
        // Binds the request json to requestBody
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
            return
        }
        // Accesses the nested data
        adventureId := requestBody.Data.Attributes.Adventure_id
        // Sets the object id
        objId, err := primitive.ObjectIDFromHex(adventureId);
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Adventure ID"})
            return
        }
        // Sets the filter
        filter := bson.M{"_id": objId}
        // Delete the object from the collection
        result, err := adventureCollection.DeleteOne(ctx, filter)
        if result.DeletedCount == 0 {
            c.JSON(http.StatusOK, gin.H{"message":  "Invalid Adventure ID" })
            return
        }
        c.JSON(http.StatusOK, gin.H{"message":  "Adventure Deleted" })
    }
}

func GetAnAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
        // Set Request Body
		var requestBody requests.GetAdventureRequest
        // Binds request json to requestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
			return
		}
        // Gets Adventure ID and sets it as a var
		adventureId := requestBody.Data.Attributes.Adventure_id
        // Check if valid ID
		objId, err := primitive.ObjectIDFromHex(adventureId)
        // Returns Bad Request if invalid format
		if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid adventure ID"})
			return
		}
        // Set model type for find
        var adventure models.Adventure
        // Find adventure by objid
		err = adventureCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&adventure)
        // Returns 404 if Adventure not found
		if err != nil {
            c.JSON(http.StatusNotFound, responses.AdventureResponse{Data: map[string]interface{}{"type":"adventure" , "attributes": adventure}})
			return
		}
        // Set Response 
        var response responses.GetAdventureResponse
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

        userID := requestBody.Data.Attributes.User_id

        var adventures []models.Adventure

        cursor, _ := adventureCollection.Find(ctx, bson.M{"user_id": userID})
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Data: map[string]interface{}{"error": err.Error()}})
            return
        }
        defer cursor.Close(ctx)

        for cursor.Next(ctx) {
            var adventure models.Adventure
            if err := cursor.Decode(&adventure); err != nil {
                c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Data: map[string]interface{}{"error": err.Error()}})
                return
            }
            adventures = append(adventures, adventure)
        }

        if err := cursor.Err(); err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Data: map[string]interface{}{"error": err.Error()}})
            return
        }

        c.JSON(http.StatusOK, responses.AdventureResponse{Data: map[string]interface{}{"type":"adventures" , "attributes": adventures}})
    }
}