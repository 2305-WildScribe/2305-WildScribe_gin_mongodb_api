package models

import (
    "testing"
    // "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewAdventure(t *testing.T) {
	// Create a new Adventure
	adventure := NewAdventure("user_id", "activity", "date", "image_url", "stress_level", 5, "sleep_stress_notes", "hydration", "diet", "diet_hydration_notes", "beta_notes")

	// Verify the Adventure fields
	if adventure.User_id != "user_id" {
			t.Errorf("Expected User_id to be 'user_id', but got %s", adventure.User_id)
	}
	if adventure.Activity != "activity" {
			t.Errorf("Expected Activity to be 'activity', but got %s", adventure.Activity)
	}
	if adventure.Date != "date" {
			t.Errorf("Expected Date to be 'date', but got %s", adventure.Date)
	}
	if adventure.Image_url != "image_url" {
			t.Errorf("Expected Image_url to be 'image_url', but got %s", adventure.Image_url)
	}
	if adventure.Stress_level != "stress_level" {
			t.Errorf("Expected Stress_level to be 'stress_level', but got %s", adventure.Stress_level)
	}
	if adventure.Hours_slept != 5 {
			t.Errorf("Expected Hours_slept to be 5, but got %d", adventure.Hours_slept)
	}
	if adventure.Sleep_stress_notes != "sleep_stress_notes" {
			t.Errorf("Expected Sleep_stress_notes to be 'sleep_stress_notes', but got %s", adventure.Sleep_stress_notes)
	}
	if adventure.Hydration != "hydration" {
			t.Errorf("Expected Hydration to be 'hydration', but got %s", adventure.Hydration)
	}
	if adventure.Diet != "diet" {
			t.Errorf("Expected Diet to be 'diet', but got %s", adventure.Diet)
	}
	if adventure.Diet_hydration_notes != "diet_hydration_notes" {
			t.Errorf("Expected Diet_hydration_notes to be 'diet_hydration_notes', but got %s", adventure.Diet_hydration_notes)
	}
	if adventure.Beta_notes != "beta_notes" {
			t.Errorf("Expected Beta_notes to be 'beta_notes', but got %s", adventure.Beta_notes)
	}
}