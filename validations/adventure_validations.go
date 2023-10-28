package validations

type AdventureValidation struct {
    Data struct {
        Type string `json:"type" binding:"required"`
        Attributes struct {
            User_id              string `validate:"required"`
            Activity             string `validate:"required"`
        } 
    } 
}