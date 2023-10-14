package authentication

import "github.com/gin-gonic/gin"

func GetAuthenticatedUserID(c *gin.Context) string {
    return "user123"
}