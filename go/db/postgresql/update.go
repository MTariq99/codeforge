package postgresql

import (
	"codeforge/models"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func (db *CodeForgeDBImpl) UpdateUserDB(c *gin.Context, update *models.UpdateUserReq) error {
	tx := db.SqlDB.MustBegin()
	defer tx.Rollback() // Rollback transaction if there's an error
	var fields []string
	if update.User_Name != nil {
		fields = append(fields, "user_name=:user_name")
	}
	if update.User_bio != nil {
		fields = append(fields, "user_bio=:user_bio")
	}
	if update.User_email != nil {
		fields = append(fields, "user_email=:user_email")
	}
	if update.Updated_at != nil {
		fields = append(fields, "updated_at=:updated_at")

	}
	if len(fields) == 0 {
		return errors.New("no fields to update")
	}

	query := "UPDATE codeforge.users SET " + strings.Join(fields, ", ") + " WHERE id=:id"
	_, err := tx.NamedExec(query, update)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (db *CodeForgeDBImpl) UpdateBlogContentDB(c *gin.Context, uContent *models.UpdateBlogContentReq) error {
	tx := db.SqlDB.MustBegin()
	_, err := tx.NamedExec(`UPDATE codeforge.blogs SET content=:content,updated_at=:updated_at WHERE id=:id`, uContent)
	if err != nil {
		return err
	}
	return nil
}

func (db *CodeForgeDBImpl) UpdateBlogTitleDB(c *gin.Context, uTitle *models.UpdateBlogTitleReq) error {
	tx := db.SqlDB.MustBegin()
	_, err := tx.NamedExec(`UPDATE codeforge.blogs SET title=:title,updated_at=:updated_at WHERE id=:id`, uTitle)
	if err != nil {
		return err
	}
	return nil
}

func (db *CodeForgeDBImpl) UpdateBlogImgDB(c *gin.Context, uimg *models.UpdateBlogImgReq) error {
	tx := db.SqlDB.MustBegin()
	_, err := tx.NamedExec(`UPDATE codeforge.blogs SET blog_image=:blog_image,updated_at=:updated_at WHERE id=:id`, uimg)
	if err != nil {
		return err
	}
	return nil
}
