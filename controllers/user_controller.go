package controllers

import (
    "context"
    "gin-mongo-api/configs"
    "gin-mongo-api/models"
    "gin-mongo-api/responses"
    "gin-mongo-api/requests"
    "gin-mongo-api/serializers"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func hashAndSalt(pwd []byte) string {
    // Use GenerateFromPassword to hash & salt pwd.
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
    // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        return false
    }
    
    return true
}
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
        password := []byte(requestBody.Data.Attributes.Password)
        salted_pw := hashAndSalt(password)

        newUser := serializers.SerializeCreateUserRequest(requestBody)
        newUser.Password = salted_pw
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
            defer cancel()

            var requestBody requests.GetUserRequest
            var userResponse responses.UserResponse
            var user models.User
            // Binds http request to requestBody
            if err := c.ShouldBindJSON(&requestBody); err != nil {
                userResponse.Data.Error = "Invalid Request"
                userResponse.Data.Type = "user"
                c.JSON(http.StatusBadRequest, userResponse)
                return
            }
    
			defer cancel()
            email := requestBody.Data.Attributes.Email
            filter := bson.D{{Key:"email", Value: email}}
			err := userCollection.FindOne(ctx, filter).Decode(&user)
			if err != nil {
                userResponse.Data.Error = "Invalid Email / Password"
                userResponse.Data.Type = "user"
                userResponse.Data.Attributes = map[string]interface{}{"email": requestBody.Data.Attributes.Email, "password": requestBody.Data.Attributes.Password}
                c.JSON(http.StatusUnauthorized, userResponse)
                return
			}
            password := []byte(requestBody.Data.Attributes.Password)
            if comparePasswords(user.Password, password) == false{
                userResponse.Data.Error = "Invalid Email / Password"
                userResponse.Data.Type = "user"
                userResponse.Data.Attributes = map[string]interface{}{"email": requestBody.Data.Attributes.Email, "password": requestBody.Data.Attributes.Password}
			    c.JSON(http.StatusUnauthorized, userResponse)
                return
            }
            userResponse.Data.Type = "user"
            userResponse.Data.Attributes = map[string]interface{}{"name":user.Name, "user_id": user.User_id}
			c.JSON(http.StatusOK, userResponse)
	}
}
	