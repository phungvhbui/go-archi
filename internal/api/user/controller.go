package user

import (
	"net/http"

	"github.com/phungvhbui/go-archi/internal/model/dto"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (c UserController) GetAll(ctx *gin.Context) {
	response, err := c.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// func (c UserController) Get(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, c.service.Get())
// }

func (c UserController) Create(ctx *gin.Context) {
	var request dto.UserDTO

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	response, err := c.service.Create(ctx.Request.Context(), request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}
