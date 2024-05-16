package controllers

import (
	"fmt"
	usecase "minerva-content-status/use-case"
	"minerva-content-status/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cont *Controller) GetAllUsers(c echo.Context) error {
	fmt.Println("Executing Get All Users")
	return c.String(http.StatusOK, "Executing Get All Users")
}

func (cont *Controller) RegisterNewUser(c echo.Context) error {
	var reqBody = validators.RegisterNewUserRequestBody{}
	reqBodyErr := c.Bind(&reqBody)
	if reqBodyErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, reqBodyErr.Error())
	}

	validationErr := validators.ValidateRequest(reqBody)

	if validationErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, validationErr.Error())
	}

	cont.uc.RegisterNewUser(&usecase.RegisterNewUserStruct{Username: reqBody.Username, Email: reqBody.Email, Password: reqBody.Password})
	return c.String(http.StatusOK, "")
}
