package controller

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/service"
	"kasir-cepat-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransaksiController struct {
	service service.TransaksiService
}

func NewTransaksiController(s service.TransaksiService) TransaksiController {
	return TransaksiController{
		service: s,
	}
}

// @Summary Get Transaksi
// @Description REST API Transaksi
// @Author RasmadIbnu
// @Success 200 {object} entity.Transaksi
// @Failure 404 {object} nil
// @method [GET]
// @Router /transaksi
func (controller TransaksiController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	Transaksi, err := controller.service.List(m)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Transaksi not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Transaksi Found", http.StatusOK, Transaksi)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Transaksi
// @Description REST API Transaksi
// @Author RasmadIbnu
// @Success 200 {object} entity.Transaksi
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /transaksi
func (controller TransaksiController) Store(ctx *gin.Context) {
	var req entity.Transaksi

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to crate Transaksi", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Transaksi, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create Transaksi", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Transaksi", http.StatusOK, Transaksi)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Transaksi
// @Description REST API Transaksi
// @Author RasmadIbnu
// @Success 200 {object} entity.Transaksi
// @Failure 404 {object} nil
// @method [GET]
// @Router /transaksi/:id
func (controller TransaksiController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	Transaksi, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Transaksi not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Transaksi Found", http.StatusOK, Transaksi)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Transaksi
// @Description REST API Transaksi
// @Author RasmadIbnu
// @Success 200 {object} entity.Transaksi
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /transaksi/:id
func (controller TransaksiController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Transaksi not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.Transaksi

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Transaksi", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Transaksi, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Transaksi", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Transaksi", http.StatusOK, Transaksi)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Transaksi
// @Description REST API Transaksi
// @Author RasmadIbnu
// @Success 200 {object} entity.Transaksi
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /transaksi/:id
func (controller TransaksiController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Transaksi not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	Transaksi, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete Transaksi", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Transaksi", http.StatusOK, Transaksi)

	ctx.JSON(http.StatusOK, resp)
}
