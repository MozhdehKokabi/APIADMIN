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

func ReadUser(c echo.Context) error {
	var req models.ReqUser
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	newPerson := models.ReqUser{
		Password: req.Password,
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
	}

	token, err := auth.GenerateJWT(req.UserName)
	if err != nil {
		return err
	}
	var jwttoken = fmt.Sprintf("token: " + token)
	c.Response().Header().Set("Authorization", jwttoken)

	var password string
	var email string
	var phone int
	var address string
	err = repository.Db.QueryRow("select password, email, phone, address from students where username = $1", req.UserName).Scan(&password, &email, &phone, &address)
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err == nil {
		newPerson.Address = address
		newPerson.Email = email
		newPerson.Phone = phone

		fmt.Println(newPerson)
		return c.JSON(http.StatusOK, newPerson)

	}

	return c.JSON(http.StatusOK, "Wrong UserName")
}

func UpdateUser(c echo.Context) error {
	var req models.ReqUser
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	newPerson := models.ReqUser{
		Password: req.Password,
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
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
	_ = repository.Db.QueryRow("UPDATE students SET password= $2, email= $3, phone= $4, address= $5 where username = $1 ", req.UserName, string(hash), req.Email, req.Phone, req.Address)

	return c.JSON(http.StatusOK, newPerson)
}

func DeleteUser(c echo.Context) error {
	var req models.ReqUser
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	var name string
	err = repository.Db.QueryRow("select username from students where username= $1", req.UserName).Scan(&name)

	if name == req.UserName {
		result := repository.Db.QueryRow("delete from students where username= $1", req.UserName)
		fmt.Println(result)

		return c.JSON(http.StatusOK, "Successfully Delete")
	}
	return c.JSON(http.StatusOK, "Wrong Username")
}
