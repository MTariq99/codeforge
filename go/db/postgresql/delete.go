package postgresql

import (
	"codeforge/models"

	"github.com/gin-gonic/gin"
)

func (db *CodeForgeDBImpl) DeleteUserDB(c *gin.Context, userid *int64) error {
	tx := db.SqlDB.MustBegin()
	_, err := tx.NamedExec(
		`
	delete
	from
		codeforge.users
	where
		id = :id`, map[string]interface{}{"id": userid})
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (db *CodeForgeDBImpl) DeleteBlogDB(c *gin.Context, delBlog *models.DeleteBlogReq) error {
	tx := db.SqlDB.MustBegin()
	_, err := tx.NamedExec(
		`
		DELETE 
		FROM
			codeforge.blogs
		WHERE
		id=:id
		`, map[string]interface{}{"id": delBlog.BlogId})
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
