package controller

import (
	"gin_unsplash/pkg/dto"
	"gin_unsplash/pkg/httperror"
	"gin_unsplash/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PhotoController interface {
	ListPhotos(c *gin.Context)
	FetchUnsplashPhotos(c *gin.Context)
	DeletePhotoByID(c *gin.Context)
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
			Message: "httperror parsing request",
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
			Message: "httperror parsing request",
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

func (p *photoController) DeletePhotoByID(c *gin.Context) {
	var req dto.DeletePhotoByIDRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "fail to parsing Delete Photo request"})
		return
	}
	res, err := p.photoService.DeletePhotoByID(c, req)
	if err != nil {
		if err, ok := err.(*httperror.Error); ok {
			c.JSON(err.Status, dto.ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
