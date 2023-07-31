package server

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yukitaka/longlong/server/admin/internal/interface/server/api"
	serverjwt "github.com/yukitaka/longlong/server/admin/internal/interface/server/jwt"
	serverutil "github.com/yukitaka/longlong/server/admin/internal/interface/server/util"
	"github.com/yukitaka/longlong/server/core/pkg/domain/usecase"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"github.com/yukitaka/longlong/server/core/pkg/interface/repository"
	"github.com/yukitaka/longlong/server/core/pkg/util"
	"net/http"
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
	r.GET("/members", members)

	return &Server{Echo: e}
}

func (s *Server) Run(port int, con *datastore.Connection) {
	s.Echo.Use(datastoreMiddleware(con))

	s.Logger.Fatal(s.Start(":" + strconv.Itoa(port)))
}

func members(c echo.Context) error {
	org, err := serverutil.OrganizationFromContext(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	db := c.Get("datastore").(*datastore.Connection).DB
	rep := usecase.NewOrganizationManagerRepository(
		repository.NewOrganizationsRepository(db),
		repository.NewOrganizationMembersRepository(db),
		repository.NewIndividualsRepository(db),
	)
	itr := usecase.NewOrganizationManager(org, rep)
	members, err := itr.Members()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, members)
}

func datastoreMiddleware(con *datastore.Connection) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("datastore", con)
			return next(c)
		}
	}
}
