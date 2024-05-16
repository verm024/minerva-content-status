package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cont *Controller) GetAllUsers(c echo.Context) error {
	fmt.Println("Executing Get All Users")
	return c.String(http.StatusOK, "Executing Get All Users")
}

func (cont *Controller) RegisterNewUser(c echo.Context) error {
	cont.uc.RegisterNewUser()
	return c.String(http.StatusOK, "Executing Get All Users")
}
