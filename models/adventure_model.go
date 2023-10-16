package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Adventure struct {
    User_id              string `json:"user_id" binding:"required"`
    Adventure_id         string
    Activity             string `json:"activity" binding:"required"`
    Date                 string `json:"date,omitempty"`
    Image_url            string `json:"image_url,omitempty"`
    Stress_level         string `json:"stress_level,omitempty"`
    Hours_slept          int    `json:"hours_slept,omitempty"`
    Sleep_stress_notes   string `json:"sleep_stress_notes,omitempty"`
    Hydration            string `json:"hydration,omitempty"`
    Diet                 string `json:"diet,omitempty"`
    Diet_hydration_notes string `json:"diet_hydration_notes,omitempty"`
    Beta_notes           string `json:"beta_notes,omitempty"`
}