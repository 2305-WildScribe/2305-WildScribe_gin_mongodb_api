package main_test

import (
	"bytes"
	"encoding/json"
	"gin-mongo-api/configs"
	"gin-mongo-api/controllers"
	"gin-mongo-api/requests"
	"gin-mongo-api/responses"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

var adventureCollection *mongo.Collection = configs.GetCollection(configs.DB, "adventures")
var adventure_id string
var user_id string
func TestGetAUser(t *testing.T) {
	email := "me@gmail.com"
	password := "hi"

	var requestBody requests.GetUserRequest

	requestBody.Data.Type = "user"
	requestBody.Data.Attributes.Email = email
	requestBody.Data.Attributes.Password = password

	router := gin.Default()

    router.POST("/api/v0/user", controllers.GetAUser())

    body, _ := json.Marshal(requestBody)

    req, _ := http.NewRequest(http.MethodPost, "/api/v0/user", bytes.NewBuffer(body))
    response := httptest.NewRecorder()

    router.ServeHTTP(response, req)

	var responseBody responses.AdventureResponse
	err := json.Unmarshal(response.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	// Access the adventure ID
	name := responseBody.Data.Attributes["name"].(string)
	user_id = responseBody.Data.Attributes["user_id"].(string)
    // Assert that the response code is HTTP 201 (Created)
    assert.Equal(t, http.StatusOK, response.Code)
    assert.Equal(t, "Ian", name)
    assert.Equal(t, "652edaa67a75034ea37c6652", user_id)

}
func TestCreateAdventure(t *testing.T) {
    // Create a mock adventure collection

    var requestBody requests.AdventureRequest
    requestBody.Data.Type = "adventure"
    requestBody.Data.Attributes.User_id = user_id
    requestBody.Data.Attributes.Activity = "Test Activity"

    router := gin.Default()

    router.POST("/api/v0/adventure", controllers.CreateAdventure())

    body, _ := json.Marshal(requestBody)

    req, _ := http.NewRequest(http.MethodPost, "/api/v0/adventure", bytes.NewBuffer(body))
    response := httptest.NewRecorder()

    router.ServeHTTP(response, req)

	var responseBody responses.AdventureResponse
	err := json.Unmarshal(response.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	// Access the adventure ID
	adventure_id = responseBody.Data.Attributes["adventure_id"].(string)
    // Assert that the response code is HTTP 201 (Created)
    assert.Equal(t, http.StatusCreated, response.Code)

    // Additional assertions if needed
}
func TestGetAnAdventure(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Create a request JSON body
	requestBody := `{
		"data": {
			"type": "adventure",
			"attributes": {
				"adventure_id": "` + adventure_id + `"
			}
		}
	}`
	router.POST("/api/v0/user/adventure", controllers.GetAnAdventure())

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodPost, "/api/v0/user/adventure", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(recorder, req)

	// Assert that the response code is HTTP 200 (OK)
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Perform additional assertions on the response body or other aspects as needed
}
func TestDeleteAdventure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var requestBody requests.DeleteAdventureRequest
	
	requestBody.Data.Type = "adventure"
	requestBody.Data.Attributes.Adventure_id = adventure_id

	router := gin.Default()
	
	router.DELETE("/api/v0/adventure", controllers.DeleteAdventure())
	
	body, _ := json.Marshal(requestBody)
	
	req, _ := http.NewRequest(http.MethodDelete, "/api/v0/adventure", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}


func TestGetAdventuresForUser(t *testing.T) {
    // Create a new Gin router
    router := gin.Default()

    // Create a request JSON body
    requestBody := `{
        "data": {
            "type": "adventures",
            "attributes": {
                "user_id": "` + user_id + `"
            }
        }
    }`
	router.POST("/api/v0/user/adventures", controllers.GetAdventuresForUser())

    // Create an HTTP request
    req, err := http.NewRequest(http.MethodPost, "/api/v0/user/adventures", strings.NewReader(requestBody))
    if err != nil {
        t.Fatal(err)
    }

    // Create a response recorder to capture the response
    recorder := httptest.NewRecorder()

    // Serve the request using the router
    router.ServeHTTP(recorder, req)

    // Assert that the response code is HTTP 200 (OK)
    assert.Equal(t, http.StatusOK, recorder.Code)

    // Perform additional assertions on the response body or other aspects as needed
}