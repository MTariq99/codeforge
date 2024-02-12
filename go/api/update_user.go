package api

import (
	"codeforge/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (api *CodeForgeAPIImpl) UpdateUserAPI(c *gin.Context, updateUser *models.UpdateUserReq) error {
	uid := c.MustGet("user_id").(int64)
	currentTime := time.Now().UTC().Unix()
	updateUser.Updated_at = &currentTime

	if updateUser.UserId != uid {
		return fmt.Errorf("unauthorized for updating this user")
	}
	if err := api.postgres.UpdateUserDB(c, updateUser); err != nil {
		return err
	}
	return nil
}
