package helper_response

import (
	"encoding/json"
	"minerva-content-status/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResponseHandler(c echo.Context, data interface{}) error {

	mapResponse := map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   data,
	}

	return c.JSON(http.StatusOK, mapResponse)
}

func ErrorResponseHandler(c echo.Context, err dto.CustomErrorInterface, code int) error {
	// * if the error is a custom error, it should has its own code
	if customError, customErrOk := err.(*dto.CustomError); customErrOk {
		err = customError
		// * so set the code param as custom errors code
		code = int(customError.Code)
	}

	// * but if the code is still undefined, then set the code as internal server error (unhandled)
	if code == 0 {
		code = http.StatusInternalServerError
	}

	mapResponse := map[string]interface{}{
		"status": "error",
		"code":   code,
		"data":   map[string]interface{}{},
		"error":  err.Error(),
	}

	return echo.NewHTTPError(code, mapResponse)
}

func DecodeResponseJson(jsonStr string, dataTarget interface{}) (*dto.BaseResponse, error) {
	responseObj := new(dto.BaseResponse)
	err := json.Unmarshal([]byte(jsonStr), &responseObj)

	if err != nil {
		return nil, err
	}

	if responseObj.Data != nil {
		json.Unmarshal(responseObj.Data, &dataTarget)
	}

	return responseObj, nil
}
