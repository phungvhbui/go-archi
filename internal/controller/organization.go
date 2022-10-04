package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phungvhbui/go-archi/internal/service"
)

type OrganizationController struct {
	service *service.OrganizationService
}

func NewOrganizationService(service *service.OrganizationService) *OrganizationController {
	return &OrganizationController{
		service: service,
	}
}

func (c OrganizationController) GetAll(ctx *gin.Context) {
	dto, err := c.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, dto)
}

// func (c OrganizationController) Get(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, c.service.Get())
// }

// func (c OrganizationController) Create(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, c.service.Get())
// }
