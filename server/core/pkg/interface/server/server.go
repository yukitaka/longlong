package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &Server{e}
}

func (s *Server) Run(port int) {
	s.Logger.Fatal(s.Start(":" + strconv.Itoa(port)))
}
