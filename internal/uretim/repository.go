package uretim

import (
	"errors"
	"net/http"
	appErrors "productApp/pkg/errors"
	"productApp/pkg/logger"

	mssql "github.com/microsoft/go-mssqldb"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IUretimRepository interface {
	GetUretimList(request UretimRequest) ([]UretimResponse, error)
	DeleteUretim(id int) error
	DeleteUploads(id int) error
	AddUretim(request UretimAddRequest) (int, error)
	AddUretimUploads(uretimID int, url string) error
}

type UretimRepositoryImpl struct {
	db *gorm.DB
}

func NewUretimRepositoryImpl(db *gorm.DB) IUretimRepository {
	return &UretimRepositoryImpl{db: db}
}

func (r *UretimRepositoryImpl) GetUretimList(request UretimRequest) ([]UretimResponse, error) {
	var orderList []UretimModel
	if err := r.db.Where("SIPARIS_NO = ?", request.SiparisNo).Find(&orderList).Error; err != nil {
		logger.Logger.Error("Error fetching order list", zap.Error(err))
		return nil, &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_UNKNOWN,
		}
	}
	var orderListResponse []UretimResponse
	for _, order := range orderList {
		orderListResponse = append(orderListResponse, order.ToUretimResponse())
	}
	return orderListResponse, nil
}

func (r *UretimRepositoryImpl) DeleteUretim(id int) error {
	if err := r.db.Delete(&UretimModel{}, id).Error; err != nil {
		logger.Logger.Error("Error deleting order", zap.Error(err))
		return &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_UNKNOWN,
		}
	}
	return nil
}

func (r *UretimRepositoryImpl) AddUretimUploads(uretimID int, url string) error {
	if err := r.db.Create(&UretimUploads{
		UretimID: uretimID,
		Url:      url,
	}).Error; err != nil {
		var sqlError mssql.Error
		if errors.As(err, &sqlError) {
			if sqlError.Number == 547 {
				return &appErrors.Error{
					Code:    http.StatusNotFound,
					Message: "bu siparis no ile uretim bulunamadi",
				}
			}
		}
		logger.Logger.Error("Error adding order", zap.Error(err))
		return &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_UNKNOWN,
		}
	}
	return nil
}

func (r *UretimRepositoryImpl) AddUretim(request UretimAddRequest) (int, error) {

	uretim := UretimModel{
		SiparisNo:    request.SiparisNo,
		UretimDurum:  request.UretimDurum,
		UretimYeri:   request.UretimYeri,
		Miktari:      float64(request.Miktar),
		Kullanici:    request.Kullanici,
		UretimTarihi: request.UretimTarihi,
	}

	if err := r.db.Create(&uretim).Error; err != nil {
		logger.Logger.Error("Error adding order", zap.Error(err))
		return 0, &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_UNKNOWN,
		}
	}
	return uretim.ID, nil
}

func (r *UretimRepositoryImpl) DeleteUploads(id int) error {
	if err := r.db.Delete(&UretimUploads{}, id).Where("uretim_id = ?", id).Error; err != nil {
		logger.Logger.Error("Error deleting order", zap.Error(err))
		return &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_UNKNOWN,
		}
	}
	return nil
}
