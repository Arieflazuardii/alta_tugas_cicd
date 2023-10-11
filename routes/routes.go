package routes

import (
	"praktikum/controllers"
	"praktikum/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	//create a new echo
	e := echo.New()
	
	//user routes
	userJWT := e.Group("")
	userJWT.Use(middleware.JwtMiddleware())
	userJWT.GET("/users", controllers.GetUsersController)
	userJWT.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	userJWT.DELETE("/users/:id", controllers.DeleteUserController)
	userJWT.PUT("/users/:id", controllers.UpdateUserController)

	//book routes
	// book := e.Group("/books")
	userJWT.GET("/books", controllers.GetBooksController)
	userJWT.GET("/books/:id", controllers.GetBookController)
	userJWT.POST("/books", controllers.CreateBookController)
	userJWT.DELETE("/books/:id", controllers.DeleteBookController)
	userJWT.PUT("/books/:id", controllers.UpdateBookController)

	return e
}