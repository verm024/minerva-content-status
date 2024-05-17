package controllers

import (
	"minerva-content-status/dto"
	helper_response "minerva-content-status/helper"
	"minerva-content-status/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cont *Controller) GetContentManagementDashboard(c echo.Context) error {
	reqQuery := dto.GetContentManagementDashboardRequestQuery{}
	reqQueryErr := c.Bind(&reqQuery)

	if reqQueryErr != nil {
		return helper_response.ErrorResponseHandler(c, reqQueryErr, http.StatusBadRequest)
	}

	validationErr := validators.ValidateRequest(reqQuery)

	if validationErr != nil {
		return helper_response.ErrorResponseHandler(c, validationErr, http.StatusBadRequest)
	}

	results, resultErr := cont.uc.GetContentManagementDashboard(&dto.GetContentManagementDashboardUseCaseFilter{Status: reqQuery.Status, SortBy: reqQuery.SortBy, Search: reqQuery.Search})

	if resultErr != nil {
		return helper_response.ErrorResponseHandler(c, resultErr, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, results)
}

func (cont *Controller) CreateContent(c echo.Context) error {

	reqBody := dto.CreateContentRequestDTO{}
	reqBindErr := c.Bind(&reqBody)

	if reqBindErr != nil {
		return helper_response.ErrorResponseHandler(c, reqBindErr, http.StatusBadRequest)
	}

	validateError := validators.ValidateRequest(reqBody)

	if validateError != nil {
		return helper_response.ErrorResponseHandler(c, validateError, http.StatusBadRequest)
	}

	cont.uc.CreateContent(&dto.CreateContentDTO{Title: reqBody.Title, Description: reqBody.Description})

	return helper_response.ResponseHandler(c, map[string]interface{}{})
}
