package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gin-mongo-api/controllers"
	"gin-mongo-api/requests"
	"gin-mongo-api/mocks"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"fmt"
)

// Create a mock for the mongo collection
// type MockCollection struct {
// 	mock.Mock
// }

// func (m *MockCollection) InsertOne(ctx, document interface{}) (interface{}, error) {
// 	args := m.Called(ctx, document)
// 	return args.Get(0), args.Error(1)
// }

func TestCreateUser(t *testing.T) {
	// Switch gin mode to test mode
	gin.SetMode(gin.TestMode)

	// Create a new router
	router := gin.Default()

	// Mock database interactions
	// mockCollection := new(MockCollection)

	// Point userCollection to our mock
	// userCollection := mockCollection

	// Setup the route
	router.POST("/user", controllers.CreateUser())

	// Test successful insertion
	userData := map[string]string{
		"name":     "Test",
		"email": "TestLocation",
		"password":    "TestTitle",
	}
	body, _ := json.Marshal(userData)
	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	// // Expect InsertOne to be called once and return nil as error
	// userCollection.On("InsertOne", mock.Anything, mock.Anything).Return(nil, nil)

	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusCreated, response.Code)
}
func TestCreateAdventure(t *testing.T) {
    // Create a mock adventure collection
    mockAdventureCollection := new(mocks.MockAdventureCollection)
    controllers.SetAdventureCollection(mockAdventureCollection)

    var requestBody requests.CreateAdventureRequest
    requestBody.Data.Type = "adventure"
    requestBody.Data.Attributes.User_id = "652c6cb6ab7c7d4070bc6f3f"
    requestBody.Data.Attributes.Activity = "Test Activity"

    // Create a mock context with a cancel function
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Define the expected result when inserting the adventure
    expectedResult := &mongo.InsertOneResult{InsertedID: "your-inserted-ID"}

    // Mock the InsertOne method on the adventure collection
    mockAdventureCollection.On("InsertOne", ctx, mock.Anything).Return(expectedResult, nil)

    router := gin.Default()

    router.POST("/api/v0/adventure", controllers.CreateAdventure())

    body, _ := json.Marshal(requestBody)

    req, _ := http.NewRequest(http.MethodPost, "/api/v0/adventure", bytes.NewBuffer(body))
    response := httptest.NewRecorder()

    router.ServeHTTP(response, req)

    // Assert that the response code is HTTP 201 (Created)
    assert.Equal(t, http.StatusCreated, response.Code)

    // Additional assertions if needed
}

func TestDeleteAdventure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var requestBody requests.DeleteAdventureRequest
	
	requestBody.Data.Type = "adventure"
	requestBody.Data.Attributes.Adventure_id = "652cf6787bb2c18496d70f14"

	router := gin.Default()
	
	router.DELETE("/api/v0/adventure", controllers.DeleteAdventure())
	
	body, _ := json.Marshal(requestBody)
	
	req, _ := http.NewRequest(http.MethodDelete, "/api/v0/adventure", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	
	router.ServeHTTP(response, req)
	fmt.Printf("requestBody: %+v\n", response)
	assert.Equal(t, http.StatusOK, response.Code)
}

