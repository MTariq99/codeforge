package server

import (
	"codeforge/api"
	"codeforge/middlewares"

	"github.com/gin-gonic/gin"
)

type ServerImpl interface {
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	CreateBlog(c *gin.Context)
	UpdateBlogContent(c *gin.Context)
	UpdateBlogTitle(c *gin.Context)
	DeleteBlog(c *gin.Context)
	CreateQuestion(c *gin.Context)
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
	route.POST("/login", server.LoginUser)
	route.PUT("/update-user", middlewares.Auth(), server.UpdateUser)
	route.DELETE("/delete-user/:id", middlewares.Auth(), server.DeleteUser)

	route.POST("/create-blog", middlewares.Auth(), server.CreateBlog)
	route.PATCH("/update-content", middlewares.Auth(), server.UpdateBlogContent)
	route.PATCH("/update-title", middlewares.Auth(), server.UpdateBlogTitle)
	route.PATCH("/update-img", middlewares.Auth(), server.UpdateBlogImg)
	route.DELETE("/delete-blog", middlewares.Auth(), server.DeleteBlog)

	route.POST("/ask-question", server.CreateQuestion)

	return route
}

var _ ServerImpl = &Server{}
