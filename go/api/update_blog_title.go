package api

import (
	"codeforge/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (api *CodeForgeAPIImpl) UpdateBlogTitleAPI(c *gin.Context, updateTitle *models.UpdateBlogTitleReq) error {
	currentTime := time.Now().UTC().Unix()

	uid := c.MustGet("user_id").(int64)

	if updateTitle.UserId != uid {
		return fmt.Errorf("unauthorized user for updating the title")
	}
	updateTitle.Updated_at = &currentTime

	if err := api.postgres.UpdateBlogTitleDB(c, updateTitle); err != nil {
		return err
	}

	return nil

}
