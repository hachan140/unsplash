package controller

import (
	"gin_unsplash/pkg/dto"
	"gin_unsplash/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PhotoController interface {
	ListPhotos(c *gin.Context)
	FetchUnsplashPhotos(c *gin.Context)
}

type photoController struct {
	photoService service.PhotoService
}

func NewPhotoController(serviceProvider service.Provider) PhotoController {
	return &photoController{
		photoService: serviceProvider.PhotoService(),
	}
}

func (p *photoController) FetchUnsplashPhotos(c *gin.Context) {
	var req dto.FetchUnsplashPhotosRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadGateway, dto.ErrorResponse{
			Message: "error parsing request",
		})
		return
	}
	resp, err := p.photoService.FetchUnsplashPhotos(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (p *photoController) ListPhotos(c *gin.Context) {
	var req dto.ListPhotosRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "error parsing request",
		})
		return
	}
	resp, err := p.photoService.ListPhotos(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
