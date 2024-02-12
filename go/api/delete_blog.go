package api

import (
	"codeforge/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (api *CodeForgeAPIImpl) DeleteBlogAPI(c *gin.Context, delBlog *models.DeleteBlogReq) error {
	UserId := c.MustGet("user_id").(int64)

	if UserId != delBlog.UserId {
		return fmt.Errorf("unauthorized person for deleting this blog")
	}
	if err := api.postgres.DeleteBlogDB(c, delBlog); err != nil {
		return err
	}
	return nil
}
