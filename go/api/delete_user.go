package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *CodeForgeAPIImpl) DeleteUser(c *gin.Context, UserId *string) error {
	uid := c.MustGet("user_id").(int64)
	id, err := strconv.ParseInt(*UserId, 10, 64)
	if err != nil {
		return fmt.Errorf("error in converting string to int64 : %v", err)
	}
	if id != uid {
		return fmt.Errorf("unauthorized for deleting this user")
	}

	if err := api.postgres.DeleteUserDB(c, &id); err != nil {
		return err
	}
	return nil

}
