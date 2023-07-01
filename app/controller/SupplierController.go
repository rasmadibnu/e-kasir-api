package controller

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/service"
	"kasir-cepat-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupplierController struct {
	service service.SupplierService
}

func NewSupplierController(s service.SupplierService) SupplierController {
	return SupplierController{
		service: s,
	}
}

// @Summary Get Supplier
// @Description REST API Supplier
// @Author RasmadIbnu
// @Success 200 {object} entity.Supplier
// @Failure 404 {object} nil
// @method [GET]
// @Router /suppliers
func (controller SupplierController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	suppliers, err := controller.service.List(m)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Supplier not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Supplier Found", http.StatusOK, suppliers)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Supplier
// @Description REST API Supplier
// @Author RasmadIbnu
// @Success 200 {object} entity.Supplier
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /suppliers
func (controller SupplierController) Store(ctx *gin.Context) {
	var req entity.Supplier

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to crate Supplier", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	supplier, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create Supplier", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Supplier", http.StatusOK, supplier)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Supplier
// @Description REST API Supplier
// @Author RasmadIbnu
// @Success 200 {object} entity.Supplier
// @Failure 404 {object} nil
// @method [GET]
// @Router /suppliers/:id
func (controller SupplierController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	supplier, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Supplier not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Supplier Found", http.StatusOK, supplier)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Supplier
// @Description REST API Supplier
// @Author RasmadIbnu
// @Success 200 {object} entity.Supplier
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /suppliers/:id
func (controller SupplierController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Supplier not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.Supplier

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Supplier", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	supplier, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Supplier", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Supplier", http.StatusOK, supplier)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Supplier
// @Description REST API Supplier
// @Author RasmadIbnu
// @Success 200 {object} entity.Supplier
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /suppliers/:id
func (controller SupplierController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Supplier not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	supplier, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete Supplier", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Supplier", http.StatusOK, supplier)

	ctx.JSON(http.StatusOK, resp)
}
