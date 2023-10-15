package controllers

import (
    "context"
    "gin-mongo-api/configs"
    "gin-mongo-api/models"
    "gin-mongo-api/responses"
    "gin-mongo-api/requests"
    "gin-mongo-api/serializers"
    "net/http"
    "time"
    // "fmt"
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
        defer cancel()
        // Sets a request struct
        var requestBody requests.CreateAdventureRequest
        // Binds response to request struct
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        // //validate the request body
        // if err := c.BindJSON(&requestBody); err != nil {
        //     c.JSON(http.StatusBadRequest, responses.AdventureResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
        //     return
        // }
        // fmt.Printf("requestBody: %+v\n", requestBody)
        //use the validator library to validate required fields
        if validationErr := validateAdventure.Struct(&requestBody.Data.Attributes); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.AdventureResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }
        // Serializes the request and assings it
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
	