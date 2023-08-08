package util

import (
	"github.com/labstack/echo/v4"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
)

func DatastoreMiddleware(con *datastore.Connection) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("datastore", con)
			return next(c)
		}
	}
}

func DatastoreConnection(c echo.Context) *datastore.Connection {
	return c.Get("datastore").(*datastore.Connection)
}
