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

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var userResponse responses.UserResponse
        
        defer cancel()
        var requestBody requests.CreateUserRequest
        // Binds http request to requestBody
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        newUser := serializers.SerializeCreateUserRequest(requestBody)

        result, err := userCollection.InsertOne(ctx, newUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError,userResponse)
            return
        }
        userResponse.Data.Message = "success"
        userResponse.Data.Attributes = map[string]interface{}{"user_id": result.InsertedID}
        c.JSON(http.StatusCreated, userResponse)
    }
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			userId := c.Param("userId")
			var user models.User
            var userResponse responses.UserResponse
			defer cancel()

			objId, _ := primitive.ObjectIDFromHex(userId)

			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
			if err != nil {
					c.JSON(http.StatusInternalServerError, userResponse)
					return
			}
			c.JSON(http.StatusOK, userResponse)
	}
}
	