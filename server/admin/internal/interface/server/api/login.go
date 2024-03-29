package api

import (
	"github.com/labstack/echo/v4"
	"github.com/yukitaka/longlong/server/admin/internal/interface/server/jwt"
	serverutil "github.com/yukitaka/longlong/server/admin/internal/interface/server/util"
	"github.com/yukitaka/longlong/server/core/pkg/domain/usecase"
	"github.com/yukitaka/longlong/server/core/pkg/util"
	"net/http"
)

type loginRequest struct {
	Id           string `json:"id"`
	Organization string `json:"organization"`
	Password     string `json:"password"`
}

func Login(c echo.Context) error {
	l := new(loginRequest)
	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	itr := usecase.NewAuthentication(serverutil.DatastoreConnection(c))

	individualId, organizationId, err := itr.Auth(l.Organization, l.Id, l.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	secret, err := util.GetEnvironmentValue("JWT_SECRET")
	if err != nil {
		panic(err)
	}
	token, err := jwt.CreateToken(individualId, organizationId, secret)

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
