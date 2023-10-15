package requests

type GetUserAdventureRequest struct {
	Data struct {
		Type			 string `json:"type"`
		Attributes struct {
			User_id 	 string `json:"user_id"`
		} `json:"attributes"`
	} `json:"data"`
}