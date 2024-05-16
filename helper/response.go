package helper_response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResponseHandler(c echo.Context, data interface{}, err []error) error {
	if len(err) >= 1 {
		firstErr := err[0]
		// TODO: Handle status code
		mapResponse := map[string]interface{}{
			"status": "error",
			"code":   http.StatusBadRequest,
			"data":   map[string]interface{}{},
			"error":  firstErr.Error(),
		}

		return echo.NewHTTPError(http.StatusBadRequest, mapResponse)
	}

	mapResponse := map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   data,
	}

	return c.JSON(http.StatusOK, mapResponse)
}
