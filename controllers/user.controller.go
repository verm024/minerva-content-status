package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	fmt.Println("Executing Get All Users")
	return c.String(http.StatusOK, "Executing Get All Users")
}
