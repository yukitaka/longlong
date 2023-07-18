package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
)

type jwtCustomClaims struct {
	entity.UserIdentify
	jwt.RegisteredClaims
}
