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

func TestGetAdventuresForUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var requestBody requests.GetUserAdventureRequest
	userID := "65299d4ceb708107b33729c6"
	requestBody.Data.Attributes.User_id = userID
	router := gin.Default()


	url := "/user/adventures/"

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	response := httptest.NewRecorder()

	router.POST(url, controllers.GetAdventuresForUser())
	body, _ := json.Marshal(requestBody)
	req, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	// var responseBody map[string]interface{}
	// err := json.Unmarshal(response.Body.Bytes(), &responseBody)
	// assert.Nil(t, err)
	// assert.NotNil(t, responseBody["data"])
}