package server

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yukitaka/longlong/server/core/pkg/domain/usecase"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"github.com/yukitaka/longlong/server/core/pkg/interface/repository"
	"github.com/yukitaka/longlong/server/core/pkg/util"
	"net/http"
	"strconv"
)

type loginRequest struct {
	Id           string `json:"id"`
	Organization string `json:"organization"`
	Password     string `json:"password"`
}

type Server struct {
	*echo.Echo
	*datastore.Connection
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
	r.GET("/organization", organization)

	return &Server{Echo: e}
}

func (s *Server) Run(port int, con *datastore.Connection) {
	s.Echo.Use(datastoreMiddleware(con))

	s.Logger.Fatal(s.Start(":" + strconv.Itoa(port)))
}

func login(c echo.Context) error {
	l := new(loginRequest)
	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	db := c.Get("datastore").(*datastore.Connection).DB

	rep := usecase.NewAuthenticationRepository(
		repository.NewAuthenticationsRepository(db),
		repository.NewOrganizationsRepository(db),
		repository.NewOrganizationMembersRepository(db))
	itr := usecase.NewAuthentication(rep)
	auth, err := itr.Auth(l.Organization, l.Id, l.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, auth)
}

func organization(c echo.Context) error {
	organizationId, _ := userData(c)

	rep := repository.NewOrganizationsRepository(c.Get("datastore").(*datastore.Connection).DB)
	itr := usecase.NewOrganizationFinder(rep)
	org, err := itr.FindById(organizationId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, org)
}

func userData(c echo.Context) (individualId, organizationId int) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	return claims.IndividualId, claims.OrganizationId
}

func datastoreMiddleware(con *datastore.Connection) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("datastore", con)
			return next(c)
		}
	}
}
