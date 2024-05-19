package controllers

import (
	"minerva-content-status/dto"
	helper_response "minerva-content-status/helper"
	usecase "minerva-content-status/use-case"
	"minerva-content-status/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContentManagementArcController struct {
	uc usecase.ContentManagementArcUseCaseInterface
}

func InitializeContentManagementArcController(uc usecase.ContentManagementArcUseCaseInterface) *ContentManagementArcController {
	cont := ContentManagementArcController{uc}
	return &cont
}

func (cont *ContentManagementArcController) CreateCMArc(c echo.Context) error {
	req := dto.CreateCMArcRequestDTO{}
	bindErr := c.Bind(&req)

	if bindErr != nil {
		return helper_response.ErrorResponseHandler(c, bindErr, http.StatusBadRequest)
	}

	validateErr := validators.ValidateRequest(req)

	if validateErr != nil {
		return helper_response.ErrorResponseHandler(c, validateErr, http.StatusBadRequest)
	}

	err := cont.uc.CreateCMArc(&dto.CreateCMArcUseCaseInputDTO{Title: req.Title, Description: req.Description, ContentManagementId: req.ContentManagementId})

	if err != nil {
		return helper_response.ErrorResponseHandler(c, err, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, map[string]interface{}{})
}

func (cont *ContentManagementArcController) UpdateCMArc(c echo.Context) error {
	req := dto.UpdateCMArcRequestDTO{}
	bindErr := c.Bind(&req)

	if bindErr != nil {
		return helper_response.ErrorResponseHandler(c, bindErr, http.StatusBadRequest)
	}

	validateErr := validators.ValidateRequest(req)

	if validateErr != nil {
		return helper_response.ErrorResponseHandler(c, validateErr, http.StatusBadRequest)
	}

	err := cont.uc.UpdateCMArc(&dto.UpdateCmArcUseCaseInputDTO{ContentManagementArcId: req.ContentManagementArcId, Title: req.Title, Description: req.Description, IsFinal: req.IsFinal, IsVoiceRecorded: req.IsVoiceRecorded, IsEdited: req.IsEdited})

	if err != nil {
		return helper_response.ErrorResponseHandler(c, err, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, map[string]interface{}{})
}

func (cont *ContentManagementArcController) DeleteCMArc(c echo.Context) error {
	req := dto.DeleteCMArcRequestDTO{}
	bindErr := c.Bind(&req)

	if bindErr != nil {
		return helper_response.ErrorResponseHandler(c, bindErr, http.StatusBadRequest)
	}

	validateErr := validators.ValidateRequest(req)

	if validateErr != nil {
		return helper_response.ErrorResponseHandler(c, validateErr, http.StatusBadRequest)
	}

	err := cont.uc.DeleteCMArc(req.ContentManagementArcId)

	if err != nil {
		return helper_response.ErrorResponseHandler(c, err, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, map[string]interface{}{})
}

func (cont *ContentManagementArcController) CMArcListByCMId(c echo.Context) error {
	req := dto.CMArcListByCMIdRequestDTO{}
	bindErr := c.Bind(&req)

	if bindErr != nil {
		return helper_response.ErrorResponseHandler(c, bindErr, http.StatusBadRequest)
	}

	validateErr := validators.ValidateRequest(req)

	if validateErr != nil {
		return helper_response.ErrorResponseHandler(c, validateErr, http.StatusBadRequest)
	}

	arcList, err := cont.uc.CMArcListByCMId(req.ContentManagementId)

	if err != nil {
		return helper_response.ErrorResponseHandler(c, err, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, dto.CMArcListByCMIdResponseDTO{ArcList: arcList.ArcList})
}
