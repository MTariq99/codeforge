package api

import (
	"codeforge/models"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
)

func (api *CodeForgeAPIImpl) UpdateBlogImgAPI(c *gin.Context, img *multipart.FileHeader) error {
	var uImg *models.UpdateBlogImgReq

	currentTime := time.Now().UTC().Unix()

	uid := c.MustGet("user_id").(int64)
	if uImg.UserId != uid {
		return fmt.Errorf("unauthorized user to updating this Image")
	}

	base64URL, err := EncodeImg(img)
	if err != nil {
		return err
	}
	uImg.Updated_at = &currentTime
	uImg.Img = base64URL

	if err := api.postgres.UpdateBlogImgDB(c, uImg); err != nil {
		return err
	}

	return nil
}
