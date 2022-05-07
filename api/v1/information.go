package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/skullkon/info-app/pkg/logging"
)

type InformationService interface {
}

type InformationHandler struct {
	service InformationService
	logger  *logging.Logger
}

func NewInformationHandler(s InformationService) *InformationHandler {
	return nil
}

func (h *InformationHandler) Get(c *gin.Context) {

}
