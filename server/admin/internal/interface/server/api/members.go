package api

import (
	"github.com/labstack/echo/v4"
	serverutil "github.com/yukitaka/longlong/server/admin/internal/interface/server/util"
	"github.com/yukitaka/longlong/server/core/pkg/domain/usecase"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"github.com/yukitaka/longlong/server/core/pkg/interface/repository"
	"net/http"
)

func Members(c echo.Context) error {
	org, err := serverutil.OrganizationFromContext(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	con := c.Get("datastore").(*datastore.Connection)
	rep := usecase.NewOrganizationManagerRepository(
		repository.NewOrganizationsRepository(con),
		repository.NewOrganizationMembersRepository(con),
		repository.NewIndividualsRepository(con),
	)
	itr := usecase.NewOrganizationManager(org, rep)
	members, err := itr.Members()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, members)
}
