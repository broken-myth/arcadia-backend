package middleware

import (
	"net/http"

	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
)

//Checks and authenticates the token in protected routes

func Auth(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		utils.SendResponse(c, http.StatusUnauthorized, "Authorization header not found")
		return
	}

	userID, err := utils.ValidateToken(authHeader)
	if err != nil {
		utils.SendResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.Set("userID", userID)
	c.Next()
}
