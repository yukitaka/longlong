package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	serverjwt "github.com/yukitaka/longlong/server/admin/internal/interface/server/jwt"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/usecase"
)

func UserData(c echo.Context) (individualId, organizationId int) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*serverjwt.CustomClaims)

	return claims.IndividualId, claims.OrganizationId
}

func Operator(c echo.Context) *entity.OrganizationMember {
	individualId, organizationId := UserData(c)
	member, err := usecase.NewOrganizationMemberFinder(DatastoreConnection(c)).FindById(organizationId, individualId)
	if err != nil {
		panic("Can't Find a operator")
	}

	return member
}
