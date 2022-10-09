package controller

import (
	"github.com/phungvhbui/go-archi/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (c UserController) GetAll(ctx *gin.Context) {
	dto, err := c.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, dto)
}

// func (c UserController) Get(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, c.service.Get())
// }

// func (c UserController) Create(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, c.service.Get())
// }