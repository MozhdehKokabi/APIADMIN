package handlers

import (
	"APIADMIN/function"
	"APIADMIN/models"
	"APIADMIN/repository"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUpAdmin(c echo.Context) error {

	var req models.ReqAdmin
	err := c.Bind(&req)
	if err != nil {

		return err
	}

	hashedPassword, err := function.HashPassword(req.Password)
	if err != nil {
		return err
	}

	newPerson := models.ReqAdmin{
		Password: hashedPassword,
		UserName: req.UserName,
		Email:    req.Email,
		Id:       req.Id,
	}

	insertDynStmt := `insert into "admin" ("password", "username", "email", "id", "role") values ($1, $2, $3, $4, $5)`

	result, err := repository.Db.Exec(insertDynStmt, newPerson.Password, newPerson.UserName, newPerson.Email, newPerson.Id, "Admin")

	fmt.Println(result)
	fmt.Println(err)
	return c.JSON(http.StatusOK, "Successfully SignUp")
}

func LogInAdmin(c echo.Context) error {
	var req models.ReqAdmin
	if err := c.Bind(&req); err != nil {
		return err
	}

	token, err := function.GenerateJWT(req.UserName)
	if err != nil {
		return err
	}

	c.Response().Header().Set("Authorization", "Bearer "+token)

	// Check password
	if function.CheckPassword(req.UserName, req.Password) {
		fmt.Println("Password was correct!")
		return c.JSON(http.StatusOK, "Successfully Login")
	}
	return c.JSON(http.StatusNonAuthoritativeInfo, "Wrong username or password")
}
