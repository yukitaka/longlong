package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"net/http"
	"strconv"
	"time"
)

type jwtCustomClaims struct {
	entity.UserIdentify
	jwt.RegisteredClaims
}

func CreateToken(c echo.Context) error {
	individualId, _ := strconv.Atoi(c.FormValue("individualId"))
	organizationId, _ := strconv.Atoi(c.FormValue("organizationId"))

	claims := &jwtCustomClaims{
		UserIdentify: entity.UserIdentify{
			IndividualId:   individualId,
			OrganizationId: organizationId,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
