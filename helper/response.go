package helper_response

import (
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

func ErrorResponseHandler(c echo.Context, err error, code int) error {

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
