package middleware

import (
	"net/http"

	"github.com/arifsetiawan/go-common/request"
	"github.com/labstack/echo"
	"gitlab.com/mindtrex/mindpkg/apierror"
)

// Me ...
type Me struct {
}

// GetMe is
func (m *Me) GetMe(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		bearerToken := request.GetAccessToken(req)

		// Token must be set
		if len(bearerToken) == 0 {
			return apierror.NewError(http.StatusForbidden, "Access token is not set")
		}

		return next(c)
	}
}
