package controllers

import (
	"minerva-content-status/dto"
	helper_response "minerva-content-status/helper"
	usecase "minerva-content-status/use-case"
	"minerva-content-status/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContentManagementController struct {
	uc *usecase.ContentManagementUseCase
}

func InitializeContentManagementController(uc *usecase.ContentManagementUseCase) *ContentManagementController {
	contentManagementController := ContentManagementController{uc}

	return &contentManagementController
}

func (cont *ContentManagementController) GetContentManagementDashboard(c echo.Context) error {
	reqQuery := dto.GetContentManagementDashboardRequestDTO{}
	reqQueryErr := c.Bind(&reqQuery)

	if reqQueryErr != nil {
		return helper_response.ErrorResponseHandler(c, reqQueryErr, http.StatusBadRequest)
	}

	validationErr := validators.ValidateRequest(reqQuery)

	if validationErr != nil {
		return helper_response.ErrorResponseHandler(c, validationErr, http.StatusBadRequest)
	}

	results, resultErr := cont.uc.GetContentManagementDashboard(&dto.GetContentManagementDashboardDTO{Status: reqQuery.Status, SortBy: reqQuery.SortBy, Search: reqQuery.Search})

	if resultErr != nil {
		return helper_response.ErrorResponseHandler(c, resultErr, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, results)
}

func (cont *ContentManagementController) CreateContent(c echo.Context) error {

	reqBody := dto.CreateContentRequestDTO{}
	reqBindErr := c.Bind(&reqBody)

	if reqBindErr != nil {
		return helper_response.ErrorResponseHandler(c, reqBindErr, http.StatusBadRequest)
	}

	validateError := validators.ValidateRequest(reqBody)

	if validateError != nil {
		return helper_response.ErrorResponseHandler(c, validateError, http.StatusBadRequest)
	}

	err := cont.uc.CreateContent(&dto.CreateContentDTO{Title: reqBody.Title, Description: reqBody.Description})

	if err != nil {
		return helper_response.ErrorResponseHandler(c, err, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, map[string]interface{}{})
}

func (cont *ContentManagementController) UpdateContent(c echo.Context) error {
	req := dto.UpdateContentRequestDTO{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return helper_response.ErrorResponseHandler(c, bindErr, http.StatusBadRequest)
	}

	validationErr := validators.ValidateRequest(req)

	if validationErr != nil {
		return helper_response.ErrorResponseHandler(c, validationErr, http.StatusBadRequest)
	}

	err := cont.uc.UpdateContent(&dto.UpdateContentDTO{ContentManagementId: req.ContentManagementId, Title: req.Title, Description: req.Description})

	if err != nil {
		return helper_response.ErrorResponseHandler(c, err, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, map[string]interface{}{})
}
