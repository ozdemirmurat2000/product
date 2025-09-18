package defotanim

import (
	"net/http"
	appErrors "productApp/pkg/errors"
	"productApp/pkg/logger"
	"productApp/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IDefoTanimRepository interface {
	GetDefoTanimList(DefoTanimRequest) ([]DefoTanimResponse, error)
	GetDefoByName(defoIsmi string) (DefoTanimResponse, error)
}

type DefoTanimRepositoryImpl struct {
	db *gorm.DB
}

func NewDefoTanimRepositoryImpl(db *gorm.DB) IDefoTanimRepository {
	return &DefoTanimRepositoryImpl{db: db}
}

func (r *DefoTanimRepositoryImpl) GetDefoTanimList(request DefoTanimRequest) ([]DefoTanimResponse, error) {
	request.UretimYeri = utils.Capitalize(request.UretimYeri)
	var defoTanimList []DefoTanimModel
	if err := r.db.Where("URETIM_YERI = ?", request.UretimYeri).Find(&defoTanimList).Error; err != nil {
		logger.Logger.Error("Error fetching defo tanim list", zap.Error(err))
		return nil, &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_UNKNOWN,
		}
	}
	var defoTanimListResponse []DefoTanimResponse
	for _, defoTanim := range defoTanimList {
		defoTanimListResponse = append(defoTanimListResponse, defoTanim.ToDefoTanimResponse())
	}
	return defoTanimListResponse, nil
}

func (r *DefoTanimRepositoryImpl) GetDefoByName(defoIsmi string) (DefoTanimResponse, error) {
	var defoTanim DefoTanimModel
	if err := r.db.Where("DEFO_ISMI = ?", defoIsmi).Find(&defoTanim).Error; err != nil {
		logger.Logger.Error("Error fetching defo tanim", zap.Error(err))
		return DefoTanimResponse{}, &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_UNKNOWN,
		}
	}
	return defoTanim.ToDefoTanimResponse(), nil
}
