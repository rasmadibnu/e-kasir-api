package controller

import (
	"net/http"

	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/service"
	"kasir-cepat-api/helper"
	"kasir-cepat-api/security"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service service.UserService
	auth    service.AuthService
}

func NewAuthController(s service.UserService, a service.AuthService) AuthController {
	return AuthController{
		service: s,
		auth:    a,
	}
}

// @Summary Login user
// @Description REST API User
// @Author RasmadIbnu
// @Success 200 {object} entity.User
// @Failure 400, 404 {object} err.Error, nil
// @method [POST]
// @Router /auth/login
func (controller AuthController) Login(ctx *gin.Context) {
	var loginReq entity.User

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Login Failed", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	findUser, _ := controller.service.FindByUsername(loginReq.Username)

	if findUser.ID == 0 {
		resp := helper.ErrorJSON(ctx, "The Username or Password is Incorrect", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	loggedIn, err := controller.auth.Login(loginReq)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "The Username or Password is Incorrect", http.StatusBadRequest, nil)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	token, err := security.CreateToken(loggedIn, 0)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "There was an error generating the API token. Please try again", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Login Successfully", http.StatusOK, token)

	ctx.JSON(http.StatusOK, resp)
}

// // @Summary Login user
// // @Description REST API User
// // @Author RasmadIbnu
// // @Success 200 {object} entity.User
// // @Failure 400, 404 {object} err.Error, nil
// // @method [POST]
// // @Router /auth/login
// func (controller AuthController) RefreshToken(ctx *gin.Context) {
// 	var refTokenReq request.RefreshToken

// 	if err := ctx.ShouldBindJSON(&refTokenReq); err != nil {
// 		resp := helper.ErrorJSON(ctx, "Refresh Token Failed", http.StatusBadRequest, err.Error())

// 		ctx.JSON(http.StatusBadRequest, resp)

// 		return
// 	}

// 	user, err := controller.service.FindById(refTokenReq.UserID)

// 	if err != nil {
// 		resp := helper.ErrorJSON(ctx, "User not found", http.StatusBadRequest, err.Error())

// 		ctx.JSON(http.StatusBadRequest, resp)

// 		return
// 	}

// 	newToken, err := security.CreateToken(user, refTokenReq.Expire)

// 	if err != nil {
// 		resp := helper.ErrorJSON(ctx, "There was an error generating the API token. Please try again", http.StatusBadRequest, err.Error())

// 		ctx.JSON(http.StatusBadRequest, resp)

// 		return
// 	}

// 	resp := helper.SuccessJSON(ctx, "Refresh token successfully", http.StatusOK, newToken)

// 	ctx.JSON(http.StatusOK, resp)
// }
