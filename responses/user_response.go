package responses

type UserResponse struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Data    map[string]interface{} `json:"data"`
}

type UserErrorResponse struct {
    Data struct {
        Error string `json:"error"`
        Attributes struct {
            User_id    string      `json:"user_id"`
        }   `json:"attributes"`
    }   `json:"data"`
}