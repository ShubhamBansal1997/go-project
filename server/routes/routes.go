package routes

import (
	"fmt"
	s "go-assignment/server"
	"go-assignment/server/handlers"
	"go-assignment/server/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(server *s.Server, gin *gin.Engine) {
	publicRoute := gin.Group("")
	postHandler := handlers.NewPostHandlers(server)
	authHandler := handlers.NewAuthHandler(server)
	registerHandler := handlers.NewRegisterHandler(server)
	productHandler := handlers.NewProductHandlers(server)
	cartHandler := handlers.NewCartHandlers(server)
	addressHandler := handlers.NewAddressHandlers(server)
	publicRoute.POST("/login", authHandler.Login)
	publicRoute.POST("/register", registerHandler.Register)
	publicRoute.POST("/refresh", authHandler.RefreshToken)
	publicRoute.GET("/products", productHandler.GetProducts)
	// server.Echo.POST("/login", authHandler.Login)
	// server.Echo.POST("/register", registerHandler.Register)
	// server.Echo.POST("/refresh", authHandler.RefreshToken)
	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(server))

	fmt.Println(server.Config.Auth.AccessSecret)

	// r := server.Echo.Group("")
	// // Configure middleware with the custom claims type
	// config := echojwt.Config{
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(token.JwtCustomClaims)
	// 	},
	// 	SigningKey: []byte(server.Config.Auth.AccessSecret),
	// }
	// r.Use(echojwt.WithConfig(config))
	protectedRouter.GET("/cart", cartHandler.GetCart)
	protectedRouter.POST("/cart", cartHandler.UpdateCart)

	protectedRouter.GET("/address", addressHandler.GetAddress)
	protectedRouter.POST("/address", addressHandler.CreateAddress)
	protectedRouter.DELETE("/address/:id", addressHandler.DeleteAddress)
	protectedRouter.PUT("/address/:id", addressHandler.UpdateAddress)

	protectedRouter.GET("/posts", postHandler.GetPosts)
	protectedRouter.POST("/posts", postHandler.CreatePost)
	protectedRouter.DELETE("/posts/:id", postHandler.DeletePost)
	protectedRouter.PUT("/posts/:id", postHandler.UpdatePost)
}
