package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skullkon/info-app/internal/service"
	v1 "github.com/skullkon/info-app/internal/transport/http/v1"
	"github.com/skullkon/info-app/pkg/logging"
	"github.com/swaggo/swag/example/basic/docs"
	"net/http"
)

type Handler struct {
	services *service.Services
	logger   logging.Logger
}

func NewHandler(services *service.Services, logger logging.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) Init() *gin.Engine {
	// Init gin handler
	router := gin.Default()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", "127.0.0.1", ":8080")

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.logger)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
