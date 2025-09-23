package order

import (
	"fmt"
	"mime/multipart"
	"net/http"
	appErrors "productApp/pkg/errors"
	"productApp/pkg/image_storage"
	"productApp/pkg/models"
	"time"

	"gorm.io/gorm"
)

type IOrderService interface {
	GetOrderBySiparisID(siparisID string) (*models.OrderResponse, *appErrors.Error)
	GetOrderSummaryList(islemAdi string, musteriKodu string) ([]models.OrderSummaryResponse, *appErrors.Error)
	GetOrderUretimBilgileriBySiparisID(siparisID string, uretimYeri string) (*models.SiparisUretimResponse, *appErrors.Error)
	AddNewUretim(request models.UretimAddRequest) *appErrors.Error
	DeleteUretim(id int) *appErrors.Error
	GetCustomerOrdersByIslemAdi(islemAdi string) ([]models.CustomerOrdersResponse, *appErrors.Error)
	AddNewModelResim(file *multipart.FileHeader, kodu string) (string, *appErrors.Error)
	// UpdateUretim(uretim UretimUpdateRequest) *appErrors.Error
}

type OrderServiceImpl struct {
	db   *gorm.DB
	repo OrderRepository
}

func NewOrderServiceImpl(db *gorm.DB, repo OrderRepository) IOrderService {
	return &OrderServiceImpl{repo: repo, db: db}
}

func (s *OrderServiceImpl) GetOrderBySiparisID(siparisID string) (*models.OrderResponse, *appErrors.Error) {

	orderList, err := s.repo.GetOrderBySiparisID(siparisID)
	if err != nil {
		return nil, err
	}
	return orderList, nil
}

func (s *OrderServiceImpl) GetOrderSummaryList(islemAdi string, musteriKodu string) ([]models.OrderSummaryResponse, *appErrors.Error) {
	orderSummary, err := s.repo.GetOrderSummary(islemAdi, musteriKodu)
	if err != nil {
		return nil, err
	}
	return orderSummary, nil
}

func (s *OrderServiceImpl) GetOrderUretimBilgileriBySiparisID(siparisID string, uretimYeri string) (*models.SiparisUretimResponse, *appErrors.Error) {
	orderSummary, err := s.repo.GetOrderUretimBilgileriBySiparisID(siparisID, uretimYeri)
	if err != nil {
		return nil, err
	}
	return orderSummary, nil
}

func (s *OrderServiceImpl) AddNewUretim(request models.UretimAddRequest) *appErrors.Error {

	_, errSiparis := s.repo.GetOrderBySiparisID(request.SiparisNo)
	if errSiparis != nil {
		return errSiparis
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {

		now := time.Now()
		uretimMiktari := float64(request.Miktar)

		if request.UretimYeri == "kalite" {
			request.UretimYeri = "Kalite Kontrol"
		}
		if request.UretimYeri == "yikama" {
			request.UretimYeri = "Yıkama"
		}
		if request.UretimYeri == "dokuma" {
			request.UretimYeri = "Dokuma"
		}
		if request.UretimYeri == "paketleme" {
			request.UretimYeri = "Paketleme"
		}
		if request.UretimYeri == "sevkiyat" {
			request.UretimYeri = "Sevkiyat"
		}

		uretim := models.UretimModel{
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

			url, err := image_storage.UploadImage(image, "uretim")
			if err != nil {
				fmt.Println("buraya geldi")
				return &appErrors.Error{
					Code:    http.StatusInternalServerError,
					Message: appErrors.ERR_UNKNOWN,
				}
			}
			fmt.Println("resim kayit edildi")

			errRepo := s.repo.AddUretimUploads(tx, uretim.KeyNumber, url)
			if errRepo != nil {
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

func (s *OrderServiceImpl) GetCustomerOrdersByIslemAdi(islemAdi string) ([]models.CustomerOrdersResponse, *appErrors.Error) {
	return s.repo.GetCustomerOrderByIslemAdi(islemAdi)
}

func (s *OrderServiceImpl) AddNewModelResim(file *multipart.FileHeader, kodu string) (string, *appErrors.Error) {

	url, err := image_storage.UploadImage(file, "iplik")
	if err != nil {
		return "", err
	}

	_, errRepo := s.repo.AddModelResim(url, kodu)
	if errRepo != nil {
		image_storage.DeleteImage(url)
		return "", errRepo
	}

	return url, nil
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
