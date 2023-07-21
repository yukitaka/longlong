package server

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"net/http"
	"strconv"
)

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/api/v1")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("", v1)

	return &Server{e}
}

func (s *Server) Run(port int) {
	s.Logger.Fatal(s.Start(":" + strconv.Itoa(port)))
}

func v1(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	individualId := claims.IndividualId
	organizationId := claims.OrganizationId

	return c.JSON(http.StatusOK, entity.UserIdentify{IndividualId: individualId, OrganizationId: organizationId})
}
