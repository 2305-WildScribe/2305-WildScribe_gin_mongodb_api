package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Adventure struct {
    User_id              string `json:"user_id" binding:"required"`
    Adventure_id         primitive.ObjectID `json:"adventure_id,omitempty" bson:"_id,omitempty"`
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

func NewAdventure(
    user_id string,
    activity string,
    date string,
    image_url string,
    stress_level string,
    hours_slept int,
    sleep_stress_notes string,
    hydration string,
    diet string,
    diet_hydration_notes string,
    beta_notes string,
) *Adventure {
    return &Adventure{
        User_id:              user_id,
        Activity:             activity,
        Date:                 date,
        Image_url:            image_url,
        Stress_level:         stress_level,
        Hours_slept:          hours_slept,
        Sleep_stress_notes:   sleep_stress_notes,
        Hydration:            hydration,
        Diet:                 diet,
        Diet_hydration_notes: diet_hydration_notes,
        Beta_notes:           beta_notes,
    }
}