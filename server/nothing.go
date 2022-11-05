package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) nothing(ctx *gin.Context) {


	ctx.JSON(http.StatusOK, "yes")
}
