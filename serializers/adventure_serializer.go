package serializers


import (
	"gin-mongo-api/models"
	"gin-mongo-api/requests"
)

func SerializeCreateAdventureRequest(req requests.AdventureRequest) models.Adventure {
	return models.Adventure{
		User_id:                    req.Data.Attributes.User_id,
		Activity:                   req.Data.Attributes.Activity,
		Date:                       req.Data.Attributes.Date,
		Image_url:                  req.Data.Attributes.Image_url,
		Stress_level:               req.Data.Attributes.Stress_level,
		Hours_slept:                req.Data.Attributes.Hours_slept,
		Sleep_stress_notes:         req.Data.Attributes.Sleep_stress_notes,
		Hydration:                  req.Data.Attributes.Hydration,
		Diet:                       req.Data.Attributes.Diet,
		Diet_hydration_notes:       req.Data.Attributes.Diet_hydration_notes,
		Beta_notes:                 req.Data.Attributes.Beta_notes,
	}
}
