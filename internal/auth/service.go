package auth

import (
	"net/http"
	appErrors "productApp/pkg/errors"
	"productApp/pkg/jwt"
	"productApp/pkg/models"
	"time"
)

type AuthService interface {
	Login(userName, password string) (models.LoginResponse, error)
}

type AuthServiceImpl struct {
	r AuthRepository
}

func NewAuthService(r AuthRepository) AuthService {
	return &AuthServiceImpl{r: r}
}

func (s *AuthServiceImpl) Login(userName, password string) (models.LoginResponse, error) {

	userDB, err := s.r.GetUserByUserName(userName)
	if err != nil {
		return models.LoginResponse{}, err
	}

	if userDB.Password != password {
		return models.LoginResponse{}, &appErrors.Error{Code: http.StatusBadRequest, Message: "kullanici adi veya sifre hatali"}
	}

	isAdmin := userDB.IsAdmin

	accessToken, err := jwt.GenerateJWT(userDB.ID, userDB.UserName, isAdmin, time.Now().Add(time.Hour*24))
	if err != nil {
		return models.LoginResponse{}, &appErrors.Error{Code: http.StatusInternalServerError, Message: appErrors.ERR_UNKNOWN}

	}

	refreshToken, err := jwt.GenerateJWT(userDB.ID, userDB.UserName, isAdmin, time.Now().Add(time.Hour*24*7))

	if err != nil {

		return models.LoginResponse{}, &appErrors.Error{Code: http.StatusInternalServerError, Message: appErrors.ERR_UNKNOWN}

	}

	return models.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken, User: userDB.ToUserResponse()}, nil
}
