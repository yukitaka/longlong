package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	serverutil "github.com/yukitaka/longlong/server/admin/internal/interface/server/util"
	"github.com/yukitaka/longlong/server/core/pkg/domain/usecase"
	"github.com/yukitaka/longlong/server/core/pkg/interface/authentication"
	"net/http"
)

type addRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

func Members(c echo.Context) error {
	org, err := serverutil.OrganizationFromContext(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	itr := usecase.NewOrganizationManager(org, serverutil.DatastoreConnection(c))
	members, err := itr.Members()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, members)
}

func AddMembers(c echo.Context) error {
	org, err := serverutil.OrganizationFromContext(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	m := new(addRequest)
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	itr := usecase.NewAuthentication(serverutil.DatastoreConnection(c))

	pass, err := authentication.Encrypt(m.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err = itr.Store(org.Id, m.Id, pass)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, fmt.Sprintf("%s is created on %s", m.Id, org.Name))
}
