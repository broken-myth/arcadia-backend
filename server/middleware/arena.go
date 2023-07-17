package middleware

import (
	"net/http"

	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/gin-gonic/gin"
)

//Checks and authenticates the token in protected routes

func Arena(c *gin.Context) {
	isArenaOpen, err := helper.GetConstant("is_arena_open")

	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, "Unknown Error, Try again later")
		return
	}

	if isArenaOpen == 0 {
		helper.SendError(c, http.StatusForbidden, "Arcaida: Rivals has ended. We hope you had fun!")
		return
	}

	c.Next()
}
