package http

import (
	"demo/internal/app"
	"github.com/gin-gonic/gin"
)

type Server struct {
	App    *app.App
	engine *gin.Engine
}

func NewHttpServer(a *app.App) *Server {
	return &Server{
		App: a,
	}
}

func (s *Server) Run(addr string) {
	s.engine.Run(addr)
}
