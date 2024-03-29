package postgresql

import (
	"fmt"
	"log"
	"os"

	"codeforge/config"
	"codeforge/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CodeForgeDB interface {
	CreateUserDB(c *gin.Context, user *models.User) error
	LoginUserDB(c *gin.Context, userEmail *string) (*models.LoginUserRes, error)
	UpdateUserDB(c *gin.Context, update *models.UpdateUserReq) error
	DeleteUserDB(c *gin.Context, userid *int64) error
	CreateBlogDB(c *gin.Context, blog *models.BlogsReq) error
	UpdateBlogContentDB(c *gin.Context, uContent *models.UpdateBlogContentReq) error
	UpdateBlogTitleDB(c *gin.Context, uTitle *models.UpdateBlogTitleReq) error
	UpdateBlogImgDB(c *gin.Context, uimg *models.UpdateBlogImgReq) error
	DeleteBlogDB(c *gin.Context, delBlog *models.DeleteBlogReq) error
	CreateQuestionDB(c *gin.Context, question *models.Question) error
}

type CodeForgeDBImpl struct {
	SqlDB *sqlx.DB
}

func ConnectSqlDB() *CodeForgeDBImpl {
	cfg := config.Cfg

	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", cfg.PgUser, cfg.PgPassword, cfg.PgHost, cfg.PgPort, cfg.PgDBName, cfg.SslMode)
	log.Println(psqlInfo)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to postgreSQL: %v\n", err)
		os.Exit(1)
	}

	log.Println("⛁ Connected to postgreSQL Database!")

	return &CodeForgeDBImpl{
		SqlDB: db,
	}

}
