package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	serverjwt "github.com/yukitaka/longlong/server/admin/internal/interface/server/jwt"
)

func UserData(c echo.Context) (individualId, organizationId int) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*serverjwt.CustomClaims)

	return claims.IndividualId, claims.OrganizationId
}
