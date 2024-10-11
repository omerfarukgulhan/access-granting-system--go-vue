package util

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	MaxUploadSize             = 5 * 1024 * 1024
	UploadDir                 = "uploads/profile-images"
	AllowedFileTypesForImages = ".jpg,.jpeg,.png,.gif"
)

func SaveProfileImage(c *gin.Context, image *multipart.FileHeader) (string, error) {
	if image.Size > MaxUploadSize {
		return "", errors.New("file size exceeds the 5MB limit")
	}

	ext := strings.ToLower(filepath.Ext(image.Filename))
	if !isAllowedFileType(ext) {
		return "", fmt.Errorf("invalid file type: %s. Allowed types are: %s", ext, AllowedFileTypesForImages)
	}

	newFileName := uuid.New().String() + ext
	err := os.MkdirAll(UploadDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	uploadPath := filepath.Join(UploadDir, newFileName)
	err = c.SaveUploadedFile(image, uploadPath)
	if err != nil {
		return "", fmt.Errorf("failed to save image: %v", err)
	}

	return newFileName, nil
}

func DeleteProfileImage(imagePath string) error {
	if imagePath == "default.jpg" || imagePath == "default.png" {
		return nil
	}

	fullPath := filepath.Join(UploadDir, imagePath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil
	}

	err := os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete image: %v", err)
	}

	return nil
}

func isAllowedFileType(ext string) bool {
	allowedTypes := strings.Split(AllowedFileTypesForImages, ",")
	for _, allowed := range allowedTypes {
		if ext == allowed {
			return true
		}
	}
	return false
}
