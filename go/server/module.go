package server

import (
	"codeforge/api"

	"github.com/gin-gonic/gin"
)

type ServerImpl interface {
	CreateUser(c *gin.Context)
}

type Server struct {
	api api.CodeForgeAPI
}

func NewServer() *Server {
	api := api.NewCodeForgeAPIImpl()
	return &Server{
		api: api,
	}
}

func NewServerImpl(route *gin.Engine) *gin.Engine {
	server := NewServer()
	route.POST("/signup", server.CreateUser)

	return route
}

var _ ServerImpl = &Server{}
