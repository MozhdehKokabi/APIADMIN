package handlers

import (
	"APIADMIN/function"
	"APIADMIN/models"
	"APIADMIN/repository"

	"fmt"

	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUpUser(c echo.Context) error {

	var req models.ReqUser
	err := c.Bind(&req)
	if err != nil {

		return err
	}

	hashedPassword, err := function.HashPassword(req.Password)
	if err != nil {
		return err
	}

	newPerson := models.ReqUser{
		Password: hashedPassword,
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
	}

	insertDynStmt := `insert into "students" ("password", "username", "email", "phone", "address", "role") values ($1, $2, $3, $4, $5, $6)`

	result, err := repository.Db.Exec(insertDynStmt, newPerson.Password, newPerson.UserName, newPerson.Email, newPerson.Phone, newPerson.Address, "User")

	fmt.Println(result)

	return c.JSON(http.StatusOK, "Successfully SignUp")
}

func LogInUser(c echo.Context) error {
	var req models.ReqUser
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
