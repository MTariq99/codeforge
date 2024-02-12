package api

import (
	"codeforge/middlewares"
	"codeforge/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (api *CodeForgeAPIImpl) LoginUserAPI(c *gin.Context, userLogin *models.LoginUserReq) (string, error) {
	res, err := api.postgres.LoginUserDB(c, userLogin.User_email)
	if err != nil {
		return "", err
	}
	if err := CheckPassword(userLogin.User_password, res.User_password); err != nil {
		return "", fmt.Errorf("email or password is incorrect : %v", err)
	}

	token, err := middlewares.GenerateToken(res)
	if err != nil {
		return "", fmt.Errorf("error in geneating jwt token : %v", err)
	}

	return token, err
}
