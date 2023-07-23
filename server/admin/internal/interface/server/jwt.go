package server

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"time"
)

type JwtCustomClaims struct {
	entity.UserIdentify
	jwt.RegisteredClaims
}

func CreateToken(individualId, organizationId int, secret string) (string, error) {
	claims := &JwtCustomClaims{
		UserIdentify: entity.UserIdentify{
			IndividualId:   individualId,
			OrganizationId: organizationId,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
