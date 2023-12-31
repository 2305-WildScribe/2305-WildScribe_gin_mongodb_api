package responses

type UserResponse struct {
    Data struct {
    Message         string                      `json:"message,omitempty"`
    Error           string                      `json:"error,omitempty"`
    Type            string                      `json:"type,omitempty"`
    Attributes      map[string]interface{}      `json:"attributes,omitempty"`
    }   `json:"data"`
}

type UserErrorResponse struct {
    Data struct {
        Error string `json:"error"`
        Attributes struct {
            User_id    string      `json:"user_id"`
        }   `json:"attributes"`
    }   `json:"data"`
}