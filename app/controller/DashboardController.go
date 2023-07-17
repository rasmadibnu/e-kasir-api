package controller

import (
	"kasir-cepat-api/app/service"
	"kasir-cepat-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	service service.DashboardService
}

func NewDashboardController(s service.DashboardService) DashboardController {
	return DashboardController{
		service: s,
	}
}

// @Summary Get Cart
// @Description REST API Cart
// @Author RasmadIbnu
// @Success 200 {object} entity.Cart
// @Failure 404 {object} nil
// @method [GET]
// @Router /dashboard
func (controller DashboardController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	dashboard, err := controller.service.List(m)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Dashboard not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Dashboard Found", http.StatusOK, dashboard)

	ctx.JSON(http.StatusOK, resp)
}
