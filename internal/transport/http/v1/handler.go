package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/skullkon/info-app/internal/service"
	"github.com/skullkon/info-app/pkg/logging"
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

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initInformationRoutes(v1)
		h.initClientRoutes(v1)
	}
}
