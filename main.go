package main

import (
	"APIADMIN/admin"
	"APIADMIN/handlers"
	"APIADMIN/repository"

	"fmt"

	// "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	err := repository.InitDataBase()
	if err != nil {
		fmt.Println(err)
	}
	e := echo.New()

	e.POST("/loginUser", handlers.LogInUser)
	e.DELETE("/", admin.DeleteUser)
	e.PUT("/updateuser", admin.UpdateUser)
	e.GET("/", admin.ReadUser)
	e.POST("/signupUser", handlers.SignUpUser)
	e.POST("/loginAdmin", handlers.LogInAdmin)
	e.POST("/createwebsite", handlers.Createwebsite)

	e.DELETE("/deletewebsite", admin.DeleteWebsite)
	e.PUT("/updatewebsite", admin.UpdateWebsite)
	e.GET("/readwebsite", admin.ReadWebsite)
	e.POST("/signupAdmin", handlers.SignUpAdmin)

	// e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":3000"))

}
