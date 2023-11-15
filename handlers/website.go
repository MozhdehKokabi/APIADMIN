package handlers

import (
	"APIADMIN/models"
	"APIADMIN/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Createwebsite(c echo.Context) error {
	var req models.ReqWebsite
	err := c.Bind(&req)
	if err != nil {

		return err
	}

	newPerson := models.ReqWebsite{
		Name:    req.Name,
		Domain:  req.Domain,
		Address: req.Address,
	}

	insertDynStmt := `insert into "website" ("name","address", "domain") values ($1, $2, $3)`

	result, err := repository.Db.Exec(insertDynStmt, newPerson.Name, newPerson.Address, newPerson.Domain)

	fmt.Println(result)
	return c.JSON(http.StatusOK, "Successfully create website")
}
