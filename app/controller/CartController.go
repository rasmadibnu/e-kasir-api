package controller

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/service"
	"kasir-cepat-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	service service.CartService
}

func NewCartController(s service.CartService) CartController {
	return CartController{
		service: s,
	}
}

// @Summary Get Cart
// @Description REST API Cart
// @Author RasmadIbnu
// @Success 200 {object} entity.Cart
// @Failure 404 {object} nil
// @method [GET]
// @Router /carts
func (controller CartController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	Carts, err := controller.service.List(m)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Cart not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Cart Found", http.StatusOK, Carts)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Cart
// @Description REST API Cart
// @Author RasmadIbnu
// @Success 200 {object} entity.Cart
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /carts
func (controller CartController) Store(ctx *gin.Context) {
	var req entity.Cart
	types := ctx.Param("type")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to crate Cart", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Cart, err := controller.service.Insert(req, types)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create Cart", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Cart", http.StatusOK, Cart)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Cart
// @Description REST API Cart
// @Author RasmadIbnu
// @Success 200 {object} entity.Cart
// @Failure 404 {object} nil
// @method [GET]
// @Router /carts/:id
func (controller CartController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	Cart, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Cart not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Cart Found", http.StatusOK, Cart)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Cart
// @Description REST API Cart
// @Author RasmadIbnu
// @Success 200 {object} entity.Cart
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /carts/:id
func (controller CartController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Cart not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.Cart

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Cart", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Cart, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Cart", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Cart", http.StatusOK, Cart)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Cart
// @Description REST API Cart
// @Author RasmadIbnu
// @Success 200 {object} entity.Cart
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /carts/:id
func (controller CartController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Cart not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	Cart, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete Cart", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Cart", http.StatusOK, Cart)

	ctx.JSON(http.StatusOK, resp)
}
