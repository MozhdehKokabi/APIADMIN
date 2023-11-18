package admin

import (
	"APIADMIN/auth"
	"APIADMIN/models"
	"APIADMIN/repository"
	"fmt"

	// "log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Createwebsite(c echo.Context) error {
	var req models.ReqWebsite
	err := c.Bind(&req)
	if err != nil {

		return err
	}

	newWebsite := models.ReqWebsite{
		Name:    req.Name,
		Domain:  req.Domain,
		Address: req.Address,
	}

	insertDynStmt := `insert into "website" ("name","address", "domain") values ($1, $2, $3)`

	result, err := repository.Db.Exec(insertDynStmt, newWebsite.Name, newWebsite.Address, newWebsite.Domain)

	fmt.Println(result)
	return c.JSON(http.StatusOK, "Successfully create website")
}

func ReadWebsite(c echo.Context) error {
	var req models.ReqWebsite
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	newWebsite := models.ReqWebsite{
		Name:    req.Name,
		Domain:  req.Domain,
		Address: req.Address,
	}

	token, err := auth.GenerateJWT(req.Domain)
	if err != nil {
		return err
	}
	var jwttoken = fmt.Sprintf("token: " + token)
	// var jwttoken = fmt.Sprintf("Bearer " + token) // Corrected the token format to include "Bearer"

	c.Response().Header().Set("Authorization", jwttoken)
	fmt.Println(jwttoken)
	var name string
	var domain string
	var address string
	err = repository.Db.QueryRow("select name, address from website where domain = $1", req.Domain).Scan(&name, &address)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	newWebsite.Address = address
	newWebsite.Name = name
	fmt.Println(newWebsite)

	if domain == req.Domain {
		return c.JSON(http.StatusOK, "wrong URL")
	}
	return c.JSON(http.StatusOK, newWebsite)
}

func UpdateWebsite(c echo.Context) error {
	var req models.ReqWebsite
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	newWebsite := models.ReqWebsite{
		Name:    req.Name,
		Domain:  req.Domain,
		Address: req.Address,
	}
	var domain string
	err = repository.Db.QueryRow("select domain from website where domain= $1", req.Domain).Scan(&domain)

	if domain == req.Domain {
		_ = repository.Db.QueryRow("UPDATE website SET name= $2, address= $3 where domain = $1  ", req.Domain, req.Name, req.Address)

		return c.JSON(http.StatusOK, newWebsite)
	}
	return c.JSON(http.StatusOK, "Wrong Domain")
}

func DeleteWebsite(c echo.Context) error {
	var req models.ReqWebsite
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	var domain string
	err = repository.Db.QueryRow("select domain from website where domain= $1", req.Domain).Scan(&domain)

	if domain == req.Domain {
		result := repository.Db.QueryRow("delete from website where domain= $1", req.Domain)
		fmt.Println(result)

		return c.JSON(http.StatusOK, "Successfully Delete")
	}
	return c.JSON(http.StatusOK, "Wrong Username")
}
