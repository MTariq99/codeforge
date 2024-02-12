package api

import (
	"codeforge/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (api *CodeForgeAPIImpl) UpdateBlogContentAPI(c *gin.Context, updateContent *models.UpdateBlogContentReq) error {
	currentTime := time.Now().UTC().Unix()
	uid := c.MustGet("user_id").(int64)

	if updateContent.UserId != uid {
		return fmt.Errorf("unauthorized user for updating the content")
	}
	updateContent.Updated_at = &currentTime
	if err := api.postgres.UpdateBlogContentDB(c, updateContent); err != nil {
		return err
	}

	return nil

}
