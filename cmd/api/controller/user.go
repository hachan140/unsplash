package controller

import (
	"gin_unsplash/pkg/dto"
	"gin_unsplash/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userController struct {
	userService service.UserService
}
type UserController interface {
	CreateUser(c *gin.Context)
}

func NewUserController(serviceProvider service.Provider) UserController {
	return &userController{
		userService: serviceProvider.UserService(),
	}
}

func (u *userController) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "fail when parsing request",
		})
		return
	}
	err := req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	res, err := u.userService.CreateUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
