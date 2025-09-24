package auth

import (
	"errors"
	"net/http"
	appErrors "productApp/pkg/errors"
	"productApp/pkg/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	GetUserByUserName(userName string) (models.UserModel, error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (d AuthRepositoryImpl) GetUserByUserName(userName string) (models.UserModel, error) {

	var user models.UserModel

	result := d.db.Where("KULLANICI_KODU = ?", userName).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.UserModel{}, &appErrors.Error{Code: http.StatusNotFound, Message: "Kullanici bulunamadi"}
		}
		return models.UserModel{}, &appErrors.Error{Code: http.StatusInternalServerError, Message: appErrors.ERR_UNKNOWN}
	}
	return user, nil
}
