package router

import (
	"errors"
	helper_response "minerva-content-status/helper"
	"net/http"
	"os"
	"slices"
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
			return helper_response.ErrorResponseHandler(c, errors.New("bearer token not found in header"), http.StatusUnauthorized)
		}

		claims := jwt.MapClaims{}
		tokenString := strings.TrimPrefix(authHeader, BEARER_PREFIX)
		token, parseJwtError := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if parseJwtError != nil {
			return helper_response.ErrorResponseHandler(c, parseJwtError, http.StatusUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {
			c.Set("locals", map[string]interface{}{"user_id": claims["user_id"], "username": claims["username"], "email": claims["email"], "role": claims["role"]})
			return next(c)
		}
		return helper_response.ErrorResponseHandler(c, errors.New("invalid token"), http.StatusUnauthorized)
	}
}

func (router) roleBasedRouteMiddleware(allowedRoles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			locals := c.Get("locals").(map[string]interface{})
			userRole := locals["role"].(string)

			if slices.Contains(allowedRoles, userRole) {
				return next(c)
			}
			return helper_response.ErrorResponseHandler(c, errors.New("not allowed to access this route"), http.StatusForbidden)
		}
	}
}

func (r *router) initializePreMiddleware() {
	r.e.Pre(middleware.RemoveTrailingSlash())

	secureMiddleware := secure.New(secure.Options{FrameDeny: true, ContentTypeNosniff: true, BrowserXssFilter: true, STSIncludeSubdomains: true, STSSeconds: 31536000, STSPreload: true})
	r.e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	r.e.Use(middleware.Logger())
	r.e.Use(middleware.Recover())
}
