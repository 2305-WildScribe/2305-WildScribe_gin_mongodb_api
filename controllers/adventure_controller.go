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

func userExists(ctx context.Context, userID string) (bool, error) {
    objId, _ := primitive.ObjectIDFromHex(userID) 
    count, err := userCollection.CountDocuments(ctx,  bson.M{"_id": objId})
    fmt.Printf("Filter: %v\n", objId)
    return count > 0, err
}


func CreateAdventure() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        // Sets a request struct
        var requestBody requests.CreateAdventureRequest
        // Binds response to request struct and checks for required fields
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            var adventureError   responses.AdventureError
            adventureError.Data.Error = "Required Fields Not Filled"
            adventureError.Data.Attributes.User_id = requestBody.Data.Attributes.User_id 
            adventureError.Data.Attributes.Adventure_id = "nil"
            c.JSON(http.StatusBadRequest, adventureError)
            return
        }
        // Check if User_id is a valid user
        userID := requestBody.Data.Attributes.User_id
        userExists, err := userExists(ctx, userID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
        // Sends error if user dosen't exist
        if !userExists {
            var adventureError   responses.AdventureError
            adventureError.Data.Error = "Invalid User ID"
            adventureError.Data.Attributes.User_id = requestBody.Data.Attributes.User_id 
            adventureError.Data.Attributes.Adventure_id = "nil"
            c.JSON(http.StatusBadRequest, adventureError)
            return
        }
        // If all validations pass then serializes the request and assings it
        adventure := serializers.SerializeCreateAdventureRequest(requestBody)
        // Inserts the serialized model into the db
        result, err := adventureCollection.InsertOne(ctx, adventure)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdventureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
        //Returns 201 and success when created
        c.JSON(http.StatusCreated, responses.AdventureResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}

func DeleteAdventure() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        var requestBody requests.DeleteAdventureRequest
        // Binds the request json to requestBody
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        // Accesses the nested data
        adventureId := requestBody.Data.Attributes.Adventure_id
        // Sets the object id
        objId, err := primitive.ObjectIDFromHex(adventureId)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Adventure ID"})
            return
        }
        // Sets the filter
        filter := bson.M{"_id": objId}
        // Delete the object from the collection
        if _, err := adventureCollection.DeleteOne(ctx, filter); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"message": "Adventure Deleted"}})
        
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
        var requestBody requests.GetUserAdventureRequest

        userID := requestBody.Data.Attributes.User_id

        var adventures []models.Adventure

        cursor, _ := adventureCollection.Find(ctx, bson.M{"userId": userID})
        if err := c.ShouldBindJSON(&requestBody); err != nil {
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