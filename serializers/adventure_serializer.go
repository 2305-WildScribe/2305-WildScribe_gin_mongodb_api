package serializers


import "gin-mongo-api/models"
import "gin-mongo-api/requests"

func SerializeCreateAdventureRequest(req requests.CreateAdventureRequest) models.Adventure {
    return models.Adventure{
        User_id:      req.Data.Attributes.User_id,
        Activity:     req.Data.Attributes.Activity,
        Date:         req.Data.Attributes.Date,
        Notes:        req.Data.Attributes.Notes,
        Image_url:    req.Data.Attributes.Image_url,
        Stress_level: req.Data.Attributes.Stress_level,
        Hydration:    req.Data.Attributes.Hydration,
        Diet:         req.Data.Attributes.Diet,
    }
}
