package api

import (
	"codeforge/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (api *CodeForgeAPIImpl) CreateUserAPI(c *gin.Context, user *models.User) error {
	hp, err := hashPass(user.User_password)
	if err != nil {
		return fmt.Errorf("error in hashing password : %v", err)
	}
	user.User_password = hp
	if err := api.postgres.CreateUserDB(c, user); err != nil {
		return err
	}

	return nil
}
