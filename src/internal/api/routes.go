package api

import (
	"fmt"
	"net/http"

	"workshop/src/internal/model"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	routes := gin.Default()

	routes.GET("/ping", ping)

	routes.Run(fmt.Sprintf(":%d", model.Env.AppPort))
}

func ping(context *gin.Context) {
	context.String(http.StatusOK, "OK")
}
