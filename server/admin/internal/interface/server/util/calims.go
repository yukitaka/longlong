package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	serverjwt "github.com/yukitaka/longlong/server/admin/internal/interface/server/jwt"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/usecase"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
)

func UserData(c echo.Context) (individualId, organizationId int) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*serverjwt.CustomClaims)

	return claims.IndividualId, claims.OrganizationId
}

func Operator(c echo.Context) *entity.OrganizationMember {
	individualId, organizationId := UserData(c)
	con := c.Get("datastore").(*datastore.Connection)
	member, err := usecase.NewOrganizationMemberFinder(con).FindById(organizationId, individualId)
	if err != nil {
		panic("Can't Find a operator")
	}

	return member
}
