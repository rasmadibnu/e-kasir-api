package controller

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/service"
	"kasir-cepat-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProdukController struct {
	service service.ProdukService
}

func NewProdukController(s service.ProdukService) ProdukController {
	return ProdukController{
		service: s,
	}
}

// @Summary Get Produk
// @Description REST API Produk
// @Author RasmadIbnu
// @Success 200 {object} entity.Produk
// @Failure 404 {object} nil
// @method [GET]
// @Router /produk
func (controller ProdukController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	produk, err := controller.service.List(m)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Produk not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Produk Found", http.StatusOK, produk)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Produk
// @Description REST API Produk
// @Author RasmadIbnu
// @Success 200 {object} entity.Produk
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /produk
func (controller ProdukController) Store(ctx *gin.Context) {
	var req entity.Produk

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to crate Produk", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	produk, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create Produk", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Produk", http.StatusOK, produk)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Produk
// @Description REST API Produk
// @Author RasmadIbnu
// @Success 200 {object} entity.Produk
// @Failure 404 {object} nil
// @method [GET]
// @Router /produk/:id
func (controller ProdukController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	produk, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Produk not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Produk Found", http.StatusOK, produk)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Produk
// @Description REST API Produk
// @Author RasmadIbnu
// @Success 200 {object} entity.Produk
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /produk/:id
func (controller ProdukController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Produk not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.Produk

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Produk", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Produk, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Produk", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Produk", http.StatusOK, Produk)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Produk
// @Description REST API Produk
// @Author RasmadIbnu
// @Success 200 {object} entity.Produk
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /produk/:id
func (controller ProdukController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Produk not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	Produk, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete Produk", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Produk", http.StatusOK, Produk)

	ctx.JSON(http.StatusOK, resp)
}
