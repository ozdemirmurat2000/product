package auth

import (
	"errors"
	"net/http"
	appErrors "productApp/pkg/errors"

	"gorm.io/gorm"
)

type AuthRepository interface {
	GetUserByUserName(userName string) (UserModel, error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (d AuthRepositoryImpl) GetUserByUserName(userName string) (UserModel, error) {

	var user UserModel

	result := d.db.Where("KULLANICI_KODU = ?", userName).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return UserModel{}, &appErrors.Error{Code: http.StatusNotFound, Message: "Kullanici bulunamadi"}
		}
		return UserModel{}, &appErrors.Error{Code: http.StatusInternalServerError, Message: appErrors.ERR_UNKNOWN}
	}
	return user, nil
}
