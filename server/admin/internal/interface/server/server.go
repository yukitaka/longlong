package server

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yukitaka/longlong/server/admin/internal/interface/server/api"
	serverjwt "github.com/yukitaka/longlong/server/admin/internal/interface/server/jwt"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"github.com/yukitaka/longlong/server/core/pkg/util"
	"strconv"
)

type Server struct {
	*echo.Echo
	*datastore.Connection
}

func NewServer() *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	secret, _ := util.GetEnvironmentValue("JWT_SECRET")

	e.POST("/login", api.Login)

	r := e.Group("/api/v1")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(serverjwt.CustomClaims)
		},
		SigningMethod: jwt.SigningMethodHS256.Name,
		SigningKey:    []byte(secret),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("/organization", api.Organization)
	r.GET("/members", api.Members)

	return &Server{Echo: e}
}

func (s *Server) Run(port int, con *datastore.Connection) {
	s.Echo.Use(datastoreMiddleware(con))

	s.Logger.Fatal(s.Start(":" + strconv.Itoa(port)))
}

func datastoreMiddleware(con *datastore.Connection) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("datastore", con)
			return next(c)
		}
	}
}
