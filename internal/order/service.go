package order

import (
	"fmt"
	"net/http"
	appErrors "productApp/pkg/errors"
	image_storage "productApp/pkg/image_storage"
	"time"

	"gorm.io/gorm"
)

type IOrderService interface {
	GetOrderBySiparisID(siparisID string) (*OrderResponse, *appErrors.Error)
	GetOrderSummaryList(islemAdi string) ([]OrderSummaryResponse, *appErrors.Error)
	GetOrderUretimBilgileriBySiparisID(siparisID string, uretimYeri string) (*SiparisUretimResponse, *appErrors.Error)
	AddNewUretim(request UretimAddRequest) *appErrors.Error
	DeleteUretim(id int) *appErrors.Error
	// UpdateUretim(uretim UretimUpdateRequest) *appErrors.Error
}

type OrderServiceImpl struct {
	db           *gorm.DB
	repo         OrderRepository
	imageStorage image_storage.IImageStorage
}

func NewOrderServiceImpl(db *gorm.DB, repo OrderRepository, imageStorage image_storage.IImageStorage) IOrderService {
	return &OrderServiceImpl{repo: repo, imageStorage: imageStorage, db: db}
}

func (s *OrderServiceImpl) GetOrderBySiparisID(siparisID string) (*OrderResponse, *appErrors.Error) {

	orderList, err := s.repo.GetOrderBySiparisID(siparisID)
	if err != nil {
		return nil, err
	}
	return orderList, nil
}

func (s *OrderServiceImpl) GetOrderSummaryList(islemAdi string) ([]OrderSummaryResponse, *appErrors.Error) {
	orderSummary, err := s.repo.GetOrderSummary(islemAdi)
	if err != nil {
		return nil, err
	}
	return orderSummary, nil
}

func (s *OrderServiceImpl) GetOrderUretimBilgileriBySiparisID(siparisID string, uretimYeri string) (*SiparisUretimResponse, *appErrors.Error) {
	orderSummary, err := s.repo.GetOrderUretimBilgileriBySiparisID(siparisID, uretimYeri)
	if err != nil {
		return nil, err
	}
	return orderSummary, nil
}

func (s *OrderServiceImpl) AddNewUretim(request UretimAddRequest) *appErrors.Error {

	_, errSiparis := s.repo.GetOrderBySiparisID(request.SiparisNo)
	if errSiparis != nil {
		return errSiparis
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {

		now := time.Now()
		uretimMiktari := float64(request.Miktar)

		uretim := UretimModel{
			SiparisNo:       &request.SiparisNo,
			UretimDurum:     &request.UretimDurum,
			UretimYeri:      &request.UretimYeri,
			Miktar:          &uretimMiktari,
			Kullanici:       &request.Kullanici,
			UretimTarihSaat: &now,
		}

		if err := tx.Create(&uretim).Error; err != nil {
			return &appErrors.Error{
				Code:    http.StatusInternalServerError,
				Message: appErrors.ERR_UNKNOWN,
			}
		}

		if len(request.File) == 0 {
			return nil
		}

		for _, image := range request.File {

			url, err := s.imageStorage.UploadImage(image, "uretim")
			if err != nil {
				fmt.Println("buraya geldi")
				return &appErrors.Error{
					Code:    http.StatusInternalServerError,
					Message: appErrors.ERR_UNKNOWN,
				}
			}
			fmt.Println("resim kayit edildi")

			err = s.repo.AddUretimUploads(tx, uretim.KeyNumber, url)
			if err != nil {
				return &appErrors.Error{
					Code:    http.StatusInternalServerError,
					Message: appErrors.ERR_UNKNOWN,
				}
			}
			fmt.Println("resim bilgileri veri tabanina kayit edildi")

		}

		return nil
	})
	if err != nil {
		return &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: appErrors.ERR_UNKNOWN,
		}
	}
	return nil
}

func (s *OrderServiceImpl) DeleteUretim(id int) *appErrors.Error {

	return s.repo.DeleteUretim(id)

	// err := s.db.Transaction(func(tx *gorm.DB) error {

	// 	// resim yollarini getir

	// 	// paths, err := s.repo.GetUretimUploadsPath(tx, id)
	// 	// if err != nil {
	// 	// 	logger.Logger.Error("resim yollarini getirirken hata olustu", zap.Error(err))
	// 	// 	return err
	// 	// }

	// 	// for _, path := range paths {
	// 	// 	fmt.Println(path)
	// 	// }

	// 	// if len(paths) != 0 {
	// 	// 	for _, path := range paths {
	// 	// 		err := s.imageStorage.DeleteImage(path)
	// 	// 		if err != nil {
	// 	// 			logger.Logger.Error("resim yollarini silerken hata olustu", zap.Error(err))
	// 	// 			return err
	// 	// 		}
	// 	// 	}
	// 	// }

	// 	// veri tabanindan uretimi sil
	// 	err := s.repo.DeleteUretim(tx, id)

	// 	if err != nil {
	// 		logger.Logger.Error("veri tabanindan uretimi silerken hata olustu", zap.Error(err))
	// 		return err
	// 	}

	// 	return err
	// })
	// if err != nil {
	// 	logger.Logger.Error("veri tabanindan uretimi silerken hata olustu", zap.Error(err))
	// 	return &appErrors.Error{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: appErrors.ERR_UNKNOWN,
	// 	}
	// }

	// return nil
}

// func (s *OrderServiceImpl) UpdateUretim(uretim UretimUpdateRequest) *appErrors.Error {

// 	err := s.db.Transaction(func(tx *gorm.DB) error {

// 		err := s.repo.UpdateUretim(tx, uretim)
// 		if err != nil {
// 			return err
// 		}

// 		return err
// 	})
// 	if err != nil {
// 		return &appErrors.Error{
// 			Code:    http.StatusInternalServerError,
// 			Message: appErrors.ERR_UNKNOWN,
// 		}
// 	}

// 	return nil
// }
