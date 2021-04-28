package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prakasitnan/go-jwt/src/controller"
	"github.com/prakasitnan/go-jwt/src/middleware"
	"github.com/prakasitnan/go-jwt/src/service"
	"net/http"
)

func main() {
	var loginService service.LoginService 	= service.StaticLoginService()
	var jwtService service.JWTService 		= service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.Default()

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	v1 := server.Group("/v1")
	v1.Use(middleware.AuthorizeJWT())
	{
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message":"success" })
		})
	}


	server.Run(":8080")

}
