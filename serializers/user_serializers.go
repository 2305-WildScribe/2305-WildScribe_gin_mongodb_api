package serializers

import "gin-mongo-api/models"
import "gin-mongo-api/requests"

func SerializeCreateUserRequest(req requests.CreateUserRequest) models.User {
	return models.User{
		Name: 		req.Data.Attributes.Name,
		Email:  	req.Data.Attributes.Email,
		Password: 	req.Data.Attributes.Password,
	}
}