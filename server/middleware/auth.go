package middleware

import (
	"net/http"

	"github.com/delta/arcadia-backend/config"
	generalHelper "github.com/delta/arcadia-backend/server/helper/general"
	userHelper "github.com/delta/arcadia-backend/server/helper/user"
	"github.com/gin-gonic/gin"
)

//Checks and authenticates the token in protected routes

func Auth(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" || len(authHeader) < 7 {
		generalHelper.SendError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var isAdmin = false

	config := config.GetConfig()

	if authHeader == config.Auth.AdminHeader {
		isAdmin = true
	}

	userID, err := userHelper.ValidateToken(authHeader)

	if err != nil && !isAdmin {
		generalHelper.SendError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if !isAdmin {
		c.Set("userID", userID)
	} else {
		c.Set("isAdmin", true)
	}

	c.Next()
}
