package infrastructure

import (
	"github.com/H-J-Ainashi/CATech-RESTwithGolang/interfaces/controllers"
	"github.com/labstack/labstack/echo"
	"github.com/labstack/labstack/echo/middleware"
)

func Init() {
	e := echo.New()

	userController := controllers.NewUserController(NewMySqlDb())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(c echo.Context) error { return userController.GetUsers(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.GetUser(c) })
	e.POST("/users", func(c echo.Context) error { return userController.CreateUser(c) })
	e.PUT("/users/:id", func(c echo.Context) error { return userController.UpdateUser(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.DeleteUser(c) })

	e.Logger.Fatal(e.Start(":8080"))
}
