package image_storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"productApp/internal/config"
	appErrors "productApp/pkg/errors"
	"productApp/pkg/logger"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const MaxFileSize = 20 * 1024 * 1024

func UploadImage(file *multipart.FileHeader, folderPath string) (string, *appErrors.Error) {

	if err := sizeIsAllowed(file.Size); err != nil {
		return "", err
	}

	if err := extensionIsAllowed(file.Filename); err != nil {
		return "", err
	}

	// full path oluştur (disk yolu)
	fullPath := filepath.Join(config.Config.UploadFolder, folderPath)

	// klasör yoksa oluştur
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			logger.Logger.Error("Error creating directory", zap.Error(err))
			return "", err_folder_creation_failed
		}
	}

	// UUID ekleyerek dosya adı oluştur
	ext := filepath.Ext(file.Filename) // .jpg, .png vs.

	fileName := uuid.New().String() + ext
	physicalPath := filepath.Join(fullPath, fileName)

	// hedef dosyayı oluştur
	outFile, err := os.Create(physicalPath)
	if err != nil {
		logger.Logger.Error("Error creating file", zap.Error(err))
		return "", err_folder_creation_failed
	}
	defer outFile.Close()

	// kaynak dosyayı aç
	srcFile, err := file.Open()
	if err != nil {
		logger.Logger.Error("Error opening file", zap.Error(err))
		return "", err_folder_creation_failed
	}
	defer srcFile.Close()

	// dosyayı kopyala
	if _, err := io.Copy(outFile, srcFile); err != nil {
		logger.Logger.Error("Error copying file", zap.Error(err))
		return "", err_folder_creation_failed
	}

	// Kullanıcıya verilecek URL (Nginx servis edecek)
	publicURL := "/uploads/" + string(folderPath) + "/" + fileName

	return publicURL, nil
}
func DeleteImage(folderPath string) error {

	fullPath := filepath.Join(config.Config.UploadFolder, folderPath)

	if err := os.Remove(fullPath); err != nil {
		logger.Logger.Error("Error removing file", zap.Error(err))
		return err_image_delete_failed
	}

	return nil
}

func sizeIsAllowed(size int64) *appErrors.Error {
	if size > MaxFileSize {
		return err_image_is_too_large
	}

	return nil
}

func extensionIsAllowed(extension string) *appErrors.Error {
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	valid := false

	for _, ext := range allowedExtensions {
		if strings.HasSuffix(strings.ToLower(extension), ext) {
			valid = true
			break
		}
	}

	if !valid {
		return err_image_type_is_not_allowed
	}

	return nil
}
