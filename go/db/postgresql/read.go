package postgresql

import (
	"codeforge/models"

	"github.com/gin-gonic/gin"
)

func (db *CodeForgeDBImpl) LoginUserDB(c *gin.Context, userEmail *string) (*models.LoginUserRes, error) {
	user := models.LoginUserRes{}
	if err := db.SqlDB.Get(&user, `SELECT id,user_name, user_email,user_password FROM codeforge.users WHERE user_email = $1`, userEmail); err != nil {
		return nil, err
	}
	return &user, nil
}
