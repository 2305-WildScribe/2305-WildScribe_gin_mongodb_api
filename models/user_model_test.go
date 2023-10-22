package models

import (
    "testing"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewUser(t *testing.T) {
    // Create a new User
    user := NewUser("name", "email", "password")

    // Verify the User fields
    if user.Name != "name" {
        t.Errorf("Expected Name to be 'name', but got %s", user.Name)
    }
    if user.Email != "email" {
        t.Errorf("Expected Email to be 'email', but got %s", user.Email)
    }
    if user.Password != "password" {
        t.Errorf("Expected Password to be 'password', but got %s", user.Password)
    }
}

func TestUserFields(t *testing.T) {
    user := NewUser("name", "email", "password")

    if user.User_id != primitive.NilObjectID {
        t.Errorf("Expected User_id to be primitive.NilObjectID, but got %v", user.User_id)
    }
}