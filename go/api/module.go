package api

import (
	"codeforge/db/postgresql"
	"codeforge/models"

	"github.com/gin-gonic/gin"
)

type CodeForgeAPI interface {
	CreateUserAPI(c *gin.Context, user *models.User) error
}

type CodeForgeAPIImpl struct {
	postgres postgresql.CodeForgeDB
}

func NewCodeForgeAPIImpl() *CodeForgeAPIImpl {
	postgres := postgresql.ConnectSqlDB()
	return &CodeForgeAPIImpl{
		postgres: postgres,
	}
}

var _ CodeForgeAPI = &CodeForgeAPIImpl{}
