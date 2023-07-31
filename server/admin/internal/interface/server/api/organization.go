package api

import (
	"github.com/labstack/echo/v4"
	"github.com/yukitaka/longlong/server/admin/internal/interface/server/util"
	"net/http"
)

func Organization(c echo.Context) error {
	org, err := util.OrganizationFromContext(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, org)
}
