package api

import (
	"codeforge/models"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
)

func (api *CodeForgeAPIImpl) CreateBlog(c *gin.Context, img *multipart.FileHeader, blog *models.BlogsReq) error {
	currentTime := time.Now().UTC().Unix()
	base64URL, err := EncodeImg(img)
	if err != nil {
		return err
	}
	blog.Created_at = &currentTime
	blog.BlogImage = base64URL

	if err := api.postgres.CreateBlogDB(c, blog); err != nil {
		return err
	}

	return nil

}
