package image_storage

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	appErrors "productApp/pkg/errors"
	"productApp/pkg/logger"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IImageStorage interface {
	UploadImage(file *multipart.FileHeader, folderPath string) (string, error)
	DeleteImage(folderPath string) error
}

type ImageStorageImpl struct {
}

const FolderPath_Base = "C:/uploads"

func NewImageStorageImpl() IImageStorage {
	return &ImageStorageImpl{}
}

func (s *ImageStorageImpl) UploadImage(file *multipart.FileHeader, folderPath string) (string, error) {
	// full path oluştur (disk yolu)
	fullPath := filepath.Join(FolderPath_Base, folderPath)

	// klasör yoksa oluştur
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			logger.Logger.Error("Error creating directory", zap.Error(err))
			return "", &appErrors.Error{
				Code:    http.StatusInternalServerError,
				Message: appErrors.ERR_IMAGE_UPLOAD,
			}
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
		return "", &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_IMAGE_UPLOAD,
		}
	}
	defer outFile.Close()

	// kaynak dosyayı aç
	srcFile, err := file.Open()
	if err != nil {
		logger.Logger.Error("Error opening file", zap.Error(err))
		return "", &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_IMAGE_UPLOAD,
		}
	}
	defer srcFile.Close()

	// dosyayı kopyala
	if _, err := io.Copy(outFile, srcFile); err != nil {
		logger.Logger.Error("Error copying file", zap.Error(err))
		return "", &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_IMAGE_UPLOAD,
		}
	}

	// Kullanıcıya verilecek URL (Nginx servis edecek)
	publicURL := "/uploads/" + string(folderPath) + "/" + fileName

	return publicURL, nil
}

func (s *ImageStorageImpl) DeleteImage(folderPath string) error {

	fullPath := filepath.Join(FolderPath_Base, folderPath)

	if err := os.Remove(fullPath); err != nil {
		logger.Logger.Error("Error removing file", zap.Error(err))
		return &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_IMAGE_DELETE,
		}
	}

	return nil
}
