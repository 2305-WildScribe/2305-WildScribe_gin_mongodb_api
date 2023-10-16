package requests

type CreateUserRequest struct {
    Data struct {
        Type       string `json:"type"`
        Attributes struct {
            Name 		string 	`json:"name"`
            Email 		string 	`json:"email"`
            Password	string 	`json:"password"`
        } `json:"attributes"`
    } `json:"data"`
}
