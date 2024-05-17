package controllers

import (
	"minerva-content-status/dto"
	helper_response "minerva-content-status/helper"
	"minerva-content-status/validators"

	"github.com/labstack/echo/v4"
)

func (cont *Controller) GetContentManagementDashboard(c echo.Context) error {
	var errors []error
	reqQuery := dto.GetContentManagementDashboardRequestQuery{}
	reqQueryErr := c.Bind(&reqQuery)

	if reqQueryErr != nil {
		errors = append(errors, reqQueryErr)
	}

	validationErr := validators.ValidateRequest(reqQuery)

	if validationErr != nil {
		errors = append(errors, validationErr)
	}

	results, resultErr := cont.uc.GetContentManagementDashboard(&dto.GetContentManagementDashboardUseCaseFilter{Status: reqQuery.Status, SortBy: reqQuery.SortBy, Search: reqQuery.Search})

	if resultErr != nil {
		errors = append(errors, resultErr)
	}

	return helper_response.ResponseHandler(c, results, errors)
}

func (cont *Controller) CreateContent(c echo.Context) error {
	var errors []error

	reqBody := dto.CreateContentRequestDTO{}
	reqBindErr := c.Bind(&reqBody)

	if reqBindErr != nil {
		errors = append(errors, reqBindErr)
	}

	validateError := validators.ValidateRequest(reqBody)

	if validateError != nil {
		errors = append(errors, validateError)
	}

	cont.uc.CreateContent(&dto.CreateContentDTO{Title: reqBody.Title, Description: reqBody.Description})

	return helper_response.ResponseHandler(c, map[string]interface{}{}, errors)
}
