package router

import (
	"fmt"
	helper_response "minerva-content-status/helper"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/unrolled/secure"
)

func (r *router) appMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		const BEARER_PREFIX string = "Bearer "

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, BEARER_PREFIX) {
			return helper_response.ResponseHandler(c, map[string]interface{}{}, []error{echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid bearer token")})
		}

		claims := jwt.MapClaims{}
		tokenString := strings.TrimPrefix(authHeader, BEARER_PREFIX)
		_, parseJwtError := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if parseJwtError != nil {
			return helper_response.ResponseHandler(c, map[string]interface{}{}, []error{parseJwtError})
		}

		for key, val := range claims {
			fmt.Printf("Key: %v, value: %v\n", key, val)
		}
		return next(c)
	}
}

func (r *router) initializePreMiddleware() {
	r.e.Pre(middleware.RemoveTrailingSlash())

	secureMiddleware := secure.New(secure.Options{FrameDeny: true, ContentTypeNosniff: true, BrowserXssFilter: true, STSIncludeSubdomains: true, STSSeconds: 31536000, STSPreload: true})
	r.e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	r.e.Use(middleware.Logger())
	r.e.Use(middleware.Recover())
}
