package api

import (
	"codeforge/models"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
)

func (api *CodeForgeAPIImpl) CreateQuestionAPI(c *gin.Context, file *multipart.FileHeader, question *models.Question) error {
	currentTime := time.Now().UTC().Unix()
	encodedImg, err := EncodeImg(file)
	if err != nil {
		return fmt.Errorf("error in encoding image : %v", err)
	}
	question.Picture = encodedImg
	question.Created_at = &currentTime

	if err := api.postgres.CreateQuestionDB(c, question); err != nil {
		return err
	}
	return nil
}
