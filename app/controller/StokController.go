package controller

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/service"
	"kasir-cepat-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StokController struct {
	service service.StokService
}

func NewStokController(s service.StokService) StokController {
	return StokController{
		service: s,
	}
}

// @Summary Get Stok
// @Description REST API Stok
// @Author RasmadIbnu
// @Success 200 {object} entity.Stok
// @Failure 404 {object} nil
// @method [GET]
// @Router /stoks
func (controller StokController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	Stok, err := controller.service.List(m)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Stok not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Stok Found", http.StatusOK, Stok)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Stok
// @Description REST API Stok
// @Author RasmadIbnu
// @Success 200 {object} entity.Stok
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /stoks
func (controller StokController) Store(ctx *gin.Context) {
	var req []entity.Stok

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to crate Stok", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Stok, err := controller.service.BacthInsert(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create Stok", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Stok", http.StatusOK, Stok)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Stok
// @Description REST API Stok
// @Author RasmadIbnu
// @Success 200 {object} entity.Stok
// @Failure 404 {object} nil
// @method [GET]
// @Router /stoks/:id
func (controller StokController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	Stok, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Stok not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Stok Found", http.StatusOK, Stok)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Stok
// @Description REST API Stok
// @Author RasmadIbnu
// @Success 200 {object} entity.Stok
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /stoks/:id
func (controller StokController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Stok not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.Stok

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Stok", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Stok, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Stok", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Stok", http.StatusOK, Stok)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Stok
// @Description REST API Stok
// @Author RasmadIbnu
// @Success 200 {object} entity.Stok
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /stoks/:id
func (controller StokController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Stok not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	Stok, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete Stok", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Stok", http.StatusOK, Stok)

	ctx.JSON(http.StatusOK, resp)
}
