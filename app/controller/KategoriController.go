package controller

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/service"
	"kasir-cepat-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KategoriController struct {
	service service.KategoriService
}

func NewKategoriController(s service.KategoriService) KategoriController {
	return KategoriController{
		service: s,
	}
}

// @Summary Get Kategori
// @Description REST API Kategori
// @Author RasmadIbnu
// @Success 200 {object} entity.Kategori
// @Failure 404 {object} nil
// @method [GET]
// @Router /kategori
func (controller KategoriController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	Kategori, err := controller.service.List(m)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Kategori not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Kategori Found", http.StatusOK, Kategori)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Kategori
// @Description REST API Kategori
// @Author RasmadIbnu
// @Success 200 {object} entity.Kategori
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /kategori
func (controller KategoriController) Store(ctx *gin.Context) {
	var req entity.Kategori

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to crate Kategori", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Kategori, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create Kategori", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Kategori", http.StatusOK, Kategori)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Kategori
// @Description REST API Kategori
// @Author RasmadIbnu
// @Success 200 {object} entity.Kategori
// @Failure 404 {object} nil
// @method [GET]
// @Router /kategori/:id
func (controller KategoriController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	Kategori, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Kategori not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Kategori Found", http.StatusOK, Kategori)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Kategori
// @Description REST API Kategori
// @Author RasmadIbnu
// @Success 200 {object} entity.Kategori
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /kategori/:id
func (controller KategoriController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Kategori not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.Kategori

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Kategori", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Kategori, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Kategori", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Kategori", http.StatusOK, Kategori)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Kategori
// @Description REST API Kategori
// @Author RasmadIbnu
// @Success 200 {object} entity.Kategori
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /kategori/:id
func (controller KategoriController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Kategori not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	Kategori, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete Kategori", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Kategori", http.StatusOK, Kategori)

	ctx.JSON(http.StatusOK, resp)
}
