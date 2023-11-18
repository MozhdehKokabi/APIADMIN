package admin

import (
	"APIADMIN/auth"
	"APIADMIN/models"
	"APIADMIN/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func ReadAdmin(c echo.Context) error {
	var req models.ReqAdmin
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	newPerson := models.ReqAdmin{
		Password: req.Password,
		UserName: req.UserName,
		Email:    req.Email,
	}

	token, err := auth.GenerateJWT(req.UserName)
	if err != nil {
		return err
	}
	var jwttoken = fmt.Sprintf("token: " + token)
	c.Response().Header().Set("Authorization", jwttoken)

	var password string
	var email string

	err = repository.Db.QueryRow("select password, email,  from admin where username = $1", req.UserName).Scan(&password, &email)
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err == nil {
		newPerson.Email = email

		fmt.Println(newPerson)
		return c.JSON(http.StatusOK, newPerson)

	}

	return c.JSON(http.StatusOK, "Wrong UserName")
}

func UpdateAdmin(c echo.Context) error {
	var req models.ReqAdmin
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	newPerson := models.ReqAdmin{
		Password: req.Password,
		UserName: req.UserName,
		Email:    req.Email,
	}

	token, err := auth.GenerateJWT(req.UserName)
	if err != nil {
		return err
	}
	var jwttoken = fmt.Sprintf("token: " + token)
	c.Response().Header().Set("Authorization", jwttoken)

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))
	newPerson.Password = string(hash)
	_ = repository.Db.QueryRow("UPDATE admin SET password= $2, email= $3where username = $1 ", req.UserName, string(hash), req.Email)

	return c.JSON(http.StatusOK, newPerson)
}

func DeleteAdmin(c echo.Context) error {
	var req models.ReqAdmin
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	var name string
	err = repository.Db.QueryRow("select username from admin where username= $1", req.UserName).Scan(&name)

	if name == req.UserName {
		result := repository.Db.QueryRow("delete from admin where username= $1", req.UserName)
		fmt.Println(result)

		return c.JSON(http.StatusOK, "Successfully Delete")
	}
	return c.JSON(http.StatusOK, "Wrong Username")
}
