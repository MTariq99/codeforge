package api

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func hashPass(userPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(userPassword), 14)
	if err != nil {
		return "", fmt.Errorf("error in Hashing password")
	}
	return string(bytes), nil
}

func CheckPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
func EncodeImg(file *multipart.FileHeader) (*string, error) {
	var base64Encoding string

	uploadedFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer uploadedFile.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, uploadedFile)
	if err != nil && err != io.EOF {
		return nil, err
	}

	fileBytes := buffer.Bytes()

	mimeType := http.DetectContentType(fileBytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/jpg":
		base64Encoding += "data:image/jpg;base64,"
	}

	base64Encoding += toBase64URL(fileBytes)
	return &base64Encoding, nil
}

func toBase64URL(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

func DecodeImg(encodedString string) ([]byte, error) {
	parts := strings.Split(encodedString, ",")
	if len(parts) != 2 {
		return nil, errors.New("invalid base64 encoded string")
	}

	mimeType := parts[0]
	data := parts[1]

	if !strings.HasPrefix(mimeType, "data:image/") {
		return nil, errors.New("unsupported image format")
	}

	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return decodedData, nil
}
