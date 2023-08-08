package util

import (
	"github.com/labstack/echo/v4"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/usecase"
	"github.com/yukitaka/longlong/server/core/pkg/interface/repository"
)

func OrganizationFromContext(c echo.Context) (*entity.Organization, error) {
	organizationId, _ := UserData(c)

	rep := repository.NewOrganizationsRepository(DatastoreConnection(c))
	itr := usecase.NewOrganizationFinder(rep)

	return itr.FindById(organizationId)
}
