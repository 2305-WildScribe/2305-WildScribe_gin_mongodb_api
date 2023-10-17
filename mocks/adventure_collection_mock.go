// mocks/mockAdventureCollection.go
package mocks

import (
    "context"
    "github.com/golang/mock/gomock"
    "go.mongodb.org/mongo-driver/mongo"
	// "gin-mongo-api/collections"
)
// type AdventureCollectionMock collections.AdventureCollection


type MockAdventureCollection struct {
    ctrl *gomock.Controller
}

func NewMockAdventureCollection(ctrl *gomock.Controller) *MockAdventureCollection {
    return &MockAdventureCollection{ctrl: ctrl}
}

func (m *MockAdventureCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
    // Implement this method in your mock
    return nil, nil
}

func (m *MockAdventureCollection) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
    // Implement this method in your mock
    return nil, nil
}

func (m *MockAdventureCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
    // Implement this method in your mock
    return nil
}


