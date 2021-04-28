package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakasitnan/go-jwt/src/dto"
	"github.com/prakasitnan/go-jwt/src/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService service.JWTService
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTService) *loginController {
	return &loginController{
		loginService: loginService,
		jwtService: jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credential.Email, true)
	}
	return ""
}
