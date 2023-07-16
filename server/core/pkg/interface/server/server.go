package server

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	e := echo.New()

	return &Server{e}
}

func (s *Server) Run(port int) {
	s.Logger.Fatal(s.Start(":" + strconv.Itoa(port)))
}
