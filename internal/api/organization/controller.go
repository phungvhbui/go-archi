package organization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	service OrganizationService
}

func NewOrganizationController(service OrganizationService) *OrganizationController {
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
