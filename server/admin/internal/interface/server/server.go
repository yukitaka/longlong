package server

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/util"
	"net/http"
	"strconv"
)

type loginRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	secret, _ := util.GetEnvironmentValue("JWT_SECRET")

	e.POST("/login", login)

	r := e.Group("/api/v1")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningMethod: jwt.SigningMethodHS256.Name,
		SigningKey:    []byte(secret),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("", v1)
	r.GET("/organization", organization)

	return &Server{e}
}

func (s *Server) Run(port int) {
	s.Logger.Fatal(s.Start(":" + strconv.Itoa(port)))
}

func login(c echo.Context) error {
	l := new(loginRequest)
	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, l)
}

func v1(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	individualId := claims.IndividualId
	organizationId := claims.OrganizationId

	return c.JSON(http.StatusOK, entity.UserIdentify{IndividualId: individualId, OrganizationId: organizationId})
}

func organization(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
