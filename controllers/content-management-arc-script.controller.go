package controllers

import (
	"minerva-content-status/dto"
	helper_response "minerva-content-status/helper"
	usecase "minerva-content-status/use-case"
	"minerva-content-status/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContentManagementArcScriptController struct {
	uc usecase.ContentManagementArcScriptUseCaseInterface
}

func InitializeContentManagementArcScriptController(uc usecase.ContentManagementArcScriptUseCaseInterface) *ContentManagementArcScriptController {
	cont := ContentManagementArcScriptController{uc}
	return &cont
}

func (cont *ContentManagementArcScriptController) CreateCMAScript(c echo.Context) error {
	req := dto.CreateCMAScriptRequestDTO{}
	bindErr := c.Bind(&req)

	if bindErr != nil {
		return helper_response.ErrorResponseHandler(c, bindErr, http.StatusBadRequest)
	}

	validateErr := validators.ValidateRequest(req)

	if validateErr != nil {
		return helper_response.ErrorResponseHandler(c, validateErr, http.StatusBadRequest)
	}

	err := cont.uc.CreateCMAScript(&dto.CreateCMAScriptUseCaseInputDTO{ContentManagementArcId: req.ContentManagementArcId, ArcScript: req.ArcScript})

	if err != nil {
		return helper_response.ErrorResponseHandler(c, err, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, map[string]interface{}{})
}

func (cont *ContentManagementArcScriptController) UpdateCMAScript(c echo.Context) error {
	req := dto.UpdateCMAScriptRequestDTO{}
	bindErr := c.Bind(&req)

	if bindErr != nil {
		return helper_response.ErrorResponseHandler(c, bindErr, http.StatusBadRequest)
	}

	validateErr := validators.ValidateRequest(req)

	if validateErr != nil {
		return helper_response.ErrorResponseHandler(c, validateErr, http.StatusBadRequest)
	}

	err := cont.uc.UpdateCMAScript(&dto.UpdateCMAScriptUseCaseInputDTO{ContentManagementArcScriptId: req.ContentManagementArcScriptId, ArcScript: req.ArcScript})

	if err != nil {
		return helper_response.ErrorResponseHandler(c, err, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, map[string]interface{}{})
}

func (cont *ContentManagementArcScriptController) DeleteCMAScript(c echo.Context) error {
	req := dto.DeleteCMAScriptRequestDTO{}
	bindErr := c.Bind(&req)

	if bindErr != nil {
		return helper_response.ErrorResponseHandler(c, bindErr, http.StatusBadRequest)
	}

	validateErr := validators.ValidateRequest(req)

	if validateErr != nil {
		return helper_response.ErrorResponseHandler(c, validateErr, http.StatusBadRequest)
	}

	err := cont.uc.DeleteCMAScript(req.ContentManagementArcScriptId)

	if err != nil {
		return helper_response.ErrorResponseHandler(c, err, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, map[string]interface{}{})
}

func (cont *ContentManagementArcScriptController) CMAScriptListByCMAId(c echo.Context) error {
	req := dto.CMAScriptListByCMAIdRequestDTO{}
	bindErr := c.Bind(&req)

	if bindErr != nil {
		return helper_response.ErrorResponseHandler(c, bindErr, http.StatusBadRequest)
	}

	validateErr := validators.ValidateRequest(req)

	if validateErr != nil {
		return helper_response.ErrorResponseHandler(c, validateErr, http.StatusBadRequest)
	}

	scripts, err := cont.uc.CMAScriptListByCMAId(req.ContentManagementArcId)

	if err != nil {
		return helper_response.ErrorResponseHandler(c, err, http.StatusBadRequest)
	}

	return helper_response.ResponseHandler(c, dto.CMAScriptListByCMAIdResponseDTO{Script: scripts.Script})
}
