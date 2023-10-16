// mocks/mockAdventureCollection.go
package mocks

import (
    "context"
    "github.com/golang/mock/gomock"
    "go.mongodb.org/mongo-driver/mongo"
)

type MockAdventureCollection struct {
    ctrl     *gomock.Controller
    recorder *MockAdventureCollectionMockRecorder
}

func NewMockAdventureCollection(ctrl *gomock.Controller) *MockAdventureCollection {
    mock := &MockAdventureCollection{ctrl: ctrl}
    mock.recorder = &MockAdventureCollectionMockRecorder{mock}
    return mock
}

func (_m *MockAdventureCollection) EXPECT() *MockAdventureCollectionMockRecorder {
    return _m.recorder
}

func (_m *MockAdventureCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
    ret := _m.ctrl.Call(_m, "InsertOne", ctx, document)
    ret0, _ := ret[0].(*mongo.InsertOneResult)
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

type MockAdventureCollectionMockRecorder struct {
    mock *MockAdventureCollection
}
