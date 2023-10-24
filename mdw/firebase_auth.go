package mdw

import (
	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

func FirebaseAuth(skipper middleware.Skipper, auth *auth.Client) echo.MiddlewareFunc {
	fbauth := func(next echo.HandlerFunc) echo.HandlerFunc {
		handler := func(c echo.Context) error {
			if skipper(c) {
				return next(c)
			}

			req := c.Request()
			authHeader := strings.TrimSpace(req.Header.Get(echo.HeaderAuthorization))

			if len(authHeader) == 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "An authorization header is required")
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
			}

			ctx := c.Request().Context()
			token, err := auth.VerifyIDTokenAndCheckRevoked(ctx, bearerToken[1])
			if err != nil {
				return echo.ErrUnauthorized
			}

			c.Set(UserIDContextKey, token.UID)

			return next(c)
		}

		return handler
	}
	return fbauth
}
