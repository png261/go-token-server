package controllers

import (
	"github.com/bytesfield/golang-gin-auth-service/src/app/middlewares"
	"github.com/bytesfield/golang-gin-auth-service/src/app/responses"
	gin "github.com/gin-gonic/gin"

	docs "github.com/bytesfield/golang-gin-auth-service/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) initializeRoutes() {
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Home Route
	s.Router.GET("/", s.Home)

	api := s.Router.Group("/api/v1")
	{
		// // Auth Route
		api.POST("/login", s.Login)
		api.POST("/token/refresh", middlewares.AuthMiddleware(), s.RefreshToken)
		// //Users routes
		api.POST("/register", s.Register)
		api.GET("/users", middlewares.AuthMiddleware(), s.GetUsers)
		api.GET("/users/:id", s.GetUser)
		api.PUT("/users/:id", middlewares.AuthMiddleware(), s.UpdateUser)
		api.POST("/users/:id", middlewares.AuthMiddleware(), s.DeleteUser)
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	//Not Found Route
	s.Router.NoRoute(func(c *gin.Context) {
		responses.NotFound(c, "Not Found")
	})
}
