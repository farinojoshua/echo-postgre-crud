package controllers

import (
	"echo-user-app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func FetchUsers(c echo.Context) error {
	result, err := models.FetchUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchUser(c echo.Context) error {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FetchUser(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreUsers(c echo.Context) error {
	// name := c.FormValue("name")
	// email := c.FormValue("email")
	// age := c.FormValue("age")

	// Ageconv, err := strconv.Atoi(age)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Unable to convert age")
	// }

	tambahuser := user{}
	c.Bind(&tambahuser)

	result, err := models.StoreUser(tambahuser.Name, tambahuser.Email, tambahuser.Age)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUsers(c echo.Context) error {
	id := c.Param("id")
	// // name := c.FormValue("name")
	// // email := c.FormValue("email")
	// // age := c.FormValue("age")

	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	tambahuser := user{}
	c.Bind(&tambahuser)

	// ageConv, err := strconv.Atoi(age)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	result, err := models.UpdateUser(idConv, tambahuser.Name, tambahuser.Email, tambahuser.Age)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUsers(c echo.Context) error {
	id := c.FormValue("id")

	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
