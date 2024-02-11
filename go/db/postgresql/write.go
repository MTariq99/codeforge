package postgresql

import (
	"codeforge/models"

	"github.com/gin-gonic/gin"
)

func (db *CodeForgeDBImpl) CreateUserDB(c *gin.Context, user *models.User) error {
	tx := db.SqlDB.MustBegin()

	_, err := tx.Exec(
		`INSERT INTO codeforge.users (
            user_name,
            user_bio,
            user_email,
            user_password,
            gender,
            created_at
        ) VALUES ($1, $2, $3, $4, $5, $6)`,
		user.User_Name,
		user.User_bio,
		user.User_email,
		user.User_password,
		user.Gender,
		user.Created_at,
	)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
