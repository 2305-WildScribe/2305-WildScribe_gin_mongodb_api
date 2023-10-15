package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gin-mongo-api/controllers"
	"gin-mongo-api/requests"
	// "github.com/stretchr/testify/mock"
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
		"location": "TestLocation",
		"title":    "TestTitle",
	}
	body, _ := json.Marshal(userData)
	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	// // Expect InsertOne to be called once and return nil as error
	// userCollection.On("InsertOne", mock.Anything, mock.Anything).Return(nil, nil)

	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestDeleteAdventure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var requestBody requests.DeleteAdventureRequest
	
	requestBody.Data.Type = "adventure"
	requestBody.Data.Attributes.Adventure_id = "6529ccc80177d706906e3c72"

	router := gin.Default()
	
	router.DELETE("/adventure", controllers.DeleteAdventure())
	
	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPost, "/adventures", bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

