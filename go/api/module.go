package api

import (
	"codeforge/db/postgresql"
	"codeforge/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type CodeForgeAPI interface {
	CreateUserAPI(c *gin.Context, user *models.User) error
	LoginUserAPI(c *gin.Context, userLogin *models.LoginUserReq) (string, error)
	UpdateUserAPI(c *gin.Context, updateUser *models.UpdateUserReq) error
	DeleteUser(c *gin.Context, userId *string) error
	CreateBlog(c *gin.Context, img *multipart.FileHeader, blog *models.BlogsReq) error
	UpdateBlogContentAPI(c *gin.Context, updateContent *models.UpdateBlogContentReq) error
	UpdateBlogTitleAPI(c *gin.Context, updateTitle *models.UpdateBlogTitleReq) error
	UpdateBlogImgAPI(c *gin.Context, img *multipart.FileHeader) error
	DeleteBlogAPI(c *gin.Context, delBlog *models.DeleteBlogReq) error
	CreateQuestionAPI(c *gin.Context, file *multipart.FileHeader, question *models.Question) error
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
