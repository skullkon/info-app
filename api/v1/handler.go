package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/skullkon/info-app/pkg/logging"
	"net/http"
)

type Deps struct {
	// зависимости сервисов
}

type Handler struct {
	// хендлеры
	logger *logging.Logger
}

func NewHandler(deps Deps) *Handler {
	return nil
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}
