package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initInformationRoutes(api *gin.RouterGroup) {
	info := api.Group("/info")
	{

		info.GET("/get-all", func(context *gin.Context) {
			all, err := h.services.Information.GetAll(context)
			if err != nil {
				h.logger.Error(err)
				return
			}
			context.JSON(200, gin.H{
				"AllDevices": all,
			})
		})
		info.GET("/top-brand", func(context *gin.Context) {
			all, err := h.services.Information.GetRating(context, "brand")
			if err != nil {
				h.logger.Error(err)
				return
			}
			context.JSON(200, gin.H{
				"Brands": all,
			})
		})
		info.GET("/top-browser", func(context *gin.Context) {
			all, err := h.services.Information.GetRating(context, "browser")
			if err != nil {
				h.logger.Error(err)
				return
			}
			context.JSON(200, gin.H{
				"Browser": all,
			})
		})
		info.GET("/top-resolution", func(context *gin.Context) {
			all, err := h.services.Information.GetRating(context, "resolution")
			if err != nil {
				h.logger.Error(err)
				return
			}
			context.JSON(200, gin.H{
				"Resolution": all,
			})
		})
		info.POST("/os-version", func(context *gin.Context) {
			body := struct {
				OsName string `json:"os_name"`
			}{}
			if err := context.BindJSON(&body); err != nil {
				context.AbortWithError(http.StatusBadRequest, err)
				return
			}

			all, err := h.services.Information.GetRatingWithParam(context, "os", body.OsName, "osVersion")
			if err != nil {
				h.logger.Error(err)
				return
			}

			context.JSON(http.StatusOK, gin.H{
				"body": all,
			})
		})
		info.POST("/browser-version", func(context *gin.Context) {
			body := struct {
				BrowserName string `json:"browser_name"`
			}{}
			if err := context.BindJSON(&body); err != nil {
				context.AbortWithError(http.StatusBadRequest, err)
				return
			}

			all, err := h.services.Information.GetRatingWithParam(context, "browser", body.BrowserName, "browserVersion")
			if err != nil {
				h.logger.Error(err)
				return
			}

			context.JSON(http.StatusOK, gin.H{
				"body": all,
			})
		})
		info.POST("/phone-brand", func(context *gin.Context) {
			body := struct {
				PhoneBrand string `json:"phone_brand"`
			}{}
			if err := context.BindJSON(&body); err != nil {
				context.AbortWithError(http.StatusBadRequest, err)
				return
			}

			all, err := h.services.Information.GetRatingWithParam(context, "brand", body.PhoneBrand, "model")
			if err != nil {
				h.logger.Error(err)
				return
			}

			context.JSON(http.StatusOK, gin.H{
				"body": all,
			})
		})
	}
}
