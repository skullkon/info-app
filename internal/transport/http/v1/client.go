package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skullkon/info-app/internal/domain"
	"net/http"
)

func (h *Handler) initClientRoutes(api *gin.RouterGroup) {
	client := api.Group("/client")
	{
		client.POST("/send", func(context *gin.Context) {
			req := domain.ClientInfo{}
			if err := context.BindJSON(&req); err != nil {
				context.AbortWithError(http.StatusBadRequest, err)
				return
			}
			data, err := h.services.Information.SendData(context, req, context.Request.UserAgent())
			if err != nil {
				h.logger.Error(err)
				context.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			fmt.Println(data)
		})
	}
}
