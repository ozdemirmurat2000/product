package order

import (
	"encoding/base64"
	"fmt"
	"net/http"
	appErrors "productApp/pkg/errors"
	"productApp/pkg/logger"
	"productApp/pkg/utils"
	"sort"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrderBySiparisID(siparisID string) (*OrderResponse, *appErrors.Error)
	GetOrderSummary(islemAdi string) ([]OrderSummaryResponse, *appErrors.Error)
	GetOrderUretimBilgileriBySiparisID(siparisID string, uretimYeri string) (*SiparisUretimResponse, *appErrors.Error)
	AddUretimUploads(tx *gorm.DB, uretimID int, url string) error
	DeleteUretim(uretimID int) *appErrors.Error
	GetUretimUploadsPath(tx *gorm.DB, uretimID int) ([]string, *appErrors.Error)
	// UpdateUretim(tx *gorm.DB, uretim UretimUpdateRequest) *appErrors.Error
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepositoryImpl(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) GetOrderBySiparisID(siparisID string) (*OrderResponse, *appErrors.Error) {
	var order OrderModel

	var imageBase64 string

	_ = r.db.Raw("SELECT RESIM FROM SIPARIS_RESIM WHERE SIPARIS_NO = ?", siparisID).Scan(&imageBase64).Error

	if imageBase64 == "" {
		imageBase64 = ""
	} else {
		imageBase64 = base64.StdEncoding.EncodeToString([]byte(imageBase64))
	}

	if err := r.db.Where("SIPARIS_NO = ?", siparisID).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &appErrors.Error{
				Code:    http.StatusNotFound,
				Message: "siparis bulunamadi",
			}
		}
		return nil, &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: "siparis getirilirken bir hata olustu",
		}
	}

	return order.ToOrderResponse(imageBase64), nil
}

func (r *OrderRepositoryImpl) GetOrderSummary(islemAdi string) ([]OrderSummaryResponse, *appErrors.Error) {

	var column string

	islemAdi = utils.CapitalizeAllSmall(islemAdi)

	switch islemAdi {
	case "dokuma":
		column = "DOKUMA"
	case "yikama":
		column = "YIKAMA"
	case "kalite":
		column = "KALITEKONTROL"
	case "paketleme":
		column = "PAKETLEME"
	case "sevkiyat":
		column = "SEVKIYAT"
	default:
		return nil, &appErrors.Error{
			Code:    http.StatusBadRequest,
			Message: "islem adi gecersiz",
		}

	}

	// siparisleri getir
	var orderList []OrderModel
	if err := r.db.Where(column+" = ?", 1).Find(&orderList).Error; err != nil {
		return nil, &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: "siparis ozet getirilirken bir hata olustu",
		}
	}

	fmt.Println(len(orderList))

	orderSummary := []OrderSummaryResponse{}
	// siparisin miktarlarini getir

	for _, order := range orderList {
		var uretimler []UretimModel
		r.db.Where("SIPARIS_NO = ?", order.SiparisNo).Find(&uretimler)

		var uretimlerMap = make(map[string]float64)

		for _, uretim := range uretimler {

			if _, ok := uretimlerMap[*uretim.UretimYeri]; !ok {
				uretimlerMap[*uretim.UretimYeri] = 0
			}
			uretimlerMap[*uretim.UretimYeri] += *uretim.Miktar
		}

		orderSummary = append(orderSummary, order.ToOrderSummaryResponse(
			int(uretimlerMap["Dokuma"]),
			int(uretimlerMap["Yıkama"]),
			int(uretimlerMap["Kalite Kontrol"]),
			int(uretimlerMap["Paketleme"]),
			int(uretimlerMap["Sevkiyat"]),
		))

	}

	return orderSummary, nil
}

func (r *OrderRepositoryImpl) GetOrderUretimBilgileriBySiparisID(siparisID, uretimYeri string) (*SiparisUretimResponse, *appErrors.Error) {

	var newUretimYeri string

	if uretimYeri == "kalite" {
		newUretimYeri = "Kalite Kontrol"
	}
	if uretimYeri == "yikama" {
		newUretimYeri = "Yıkama"
	}
	if uretimYeri == "dokuma" {
		newUretimYeri = "Dokuma"
	}
	if uretimYeri == "paketleme" {
		newUretimYeri = "Paketleme"
	}
	if uretimYeri == "sevkiyat" {
		newUretimYeri = "Sevkiyat"
	}

	// uretimleri getir //
	uretimler := []UretimModel{}
	if err := r.db.Where("SIPARIS_NO = ? AND URETIM_YERI = ?", siparisID, newUretimYeri).Find(&uretimler).Error; err != nil {
		return nil, &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: "siparis uretim bilgileri getirilirken bir hata olustu",
		}
	}

	// uretimlerin resimleri varsa getir
	uretimlerResponse := []UretimResponse{}
	for _, uretim := range uretimler {
		uretimlerResponse = append(uretimlerResponse, uretim.ToUretimResponse())
	}

	for index, uretim := range uretimlerResponse {
		uretimResimleri := []UretimUploads{}
		err := r.db.Where("uretim_id = ?", uretim.KeyNumber).Find(&uretimResimleri).Error
		if err != nil {
			return nil, &appErrors.Error{
				Code:    http.StatusInternalServerError,
				Message: "siparis uretim bilgileri getirilirken bir hata olustu",
			}
		}
		uretimlerResponse[index].ImageURL = []string{}
		for _, resim := range uretimResimleri {
			uretimlerResponse[index].ImageURL = append(uretimlerResponse[index].ImageURL, resim.Url)
		}

	}

	// miktarlari hesapla

	var toplamMiktar float64
	var saglamMiktar float64
	var defoMiktar float64
	var dokumaMiktar float64
	var yikamaMiktar float64
	var kaliteMiktar float64
	var paketlemeMiktar float64
	var sevkiyatMiktar float64

	for _, uretim := range uretimler {
		var uretimYeri string

		fmt.Println("uretim yeri", *uretim.UretimYeri)

		if uretim.UretimYeri != nil {
			uretimYeri = *uretim.UretimYeri
		} else {
			uretimYeri = ""
		}

		if uretimYeri == "" {
			continue
		} else {
			uretimYeri = utils.CapitalizeAllSmall(uretimYeri)

			fmt.Println("uretim yeri", uretimYeri)

			switch uretimYeri {
			case "dokuma":
				dokumaMiktar += *uretim.Miktar
			case "yıkama":
				yikamaMiktar += *uretim.Miktar
			case "kalite kontrol":
				kaliteMiktar += *uretim.Miktar
			case "paketleme":
				paketlemeMiktar += *uretim.Miktar
			case "sevkiyat":
				sevkiyatMiktar += *uretim.Miktar
			default:
				continue
			}

		}
		toplamMiktar += *uretim.Miktar
		if *uretim.UretimDurum == "Sağlam" {
			saglamMiktar += *uretim.Miktar
		} else {
			defoMiktar += *uretim.Miktar
		}
	}

	var order OrderModel
	err := r.db.Where("SIPARIS_NO = ?", siparisID).First(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &appErrors.Error{
				Code:    http.StatusNotFound,
				Message: "siparis bulunamadi",
			}
		}
		return nil, &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: "siparis uretim bilgileri getirilirken bir hata olustu",
		}
	}

	// defolari getir

	defoList := []DefoTanimModel{}

	if uretimYeri == "kalite" {
		uretimYeri = "Kalite Kontrol"
	}
	if uretimYeri == "yikama" {
		uretimYeri = "Yıkama"
	}
	if uretimYeri == "dokuma" {
		uretimYeri = "Dokuma"
	}
	if uretimYeri == "paketleme" {
		uretimYeri = "Paketleme"
	}
	if uretimYeri == "sevkiyat" {
		uretimYeri = "Sevkiyat"
	}

	if err := r.db.Where("URETIM_YERI = ?", uretimYeri).Find(&defoList).Error; err != nil {
		return nil, &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: "siparis uretim bilgileri getirilirken bir hata olustu",
		}
	}

	defoListResponse := []DefoTanimResponse{}

	for _, defo := range defoList {
		defoListResponse = append(defoListResponse, defo.ToDefoTanimResponse())
	}

	// chart hersapla

	chartList := []ChartResponse{}

	var uretimlerMap = make(map[string]float64)

	for _, v := range uretimlerResponse {

		if _, ok := uretimlerMap[v.UretimDurum]; !ok {
			uretimlerMap[v.UretimDurum] = 0
		}
		uretimlerMap[v.UretimDurum] += v.Miktar

	}

	for k, v := range uretimlerMap {
		chartList = append(chartList, ChartResponse{
			ColorHexCode: "#FF0000",
			Percent:      (v / toplamMiktar) * 100,
			Name:         k,
		})
	}

	sort.Slice(chartList, func(i, j int) bool {
		return chartList[i].Percent > chartList[j].Percent
	})

	for i := range chartList {
		chartList[i].ColorHexCode = utils.GetColor(i)
	}

	return &SiparisUretimResponse{
		ChartResponse:  chartList,
		UretimResponse: uretimlerResponse,
		SiparisMiktari: utils.Float64Value(order.SiparisMiktari),
		ToplamMiktar:   toplamMiktar,
		SaglamMiktar:   saglamMiktar,
		DefoMiktar:     defoMiktar,
		UretimChartResponse: UretimChartResponse{
			DokumaMiktar:          dokumaMiktar,
			DokumaColorHexCode:    utils.GetColor(0),
			YikamaMiktar:          yikamaMiktar,
			YikamaColorHexCode:    utils.GetColor(1),
			KaliteMiktar:          kaliteMiktar,
			KaliteColorHexCode:    utils.GetColor(2),
			PaketlemeMiktar:       paketlemeMiktar,
			PaketlemeColorHexCode: utils.GetColor(3),
			SevkiyatMiktar:        sevkiyatMiktar,
			SevkiyatColorHexCode:  utils.GetColor(4),
		},
		DefoTanim:    defoListResponse,
		SiparisNo:    utils.StringValue(order.SiparisNo),
		MusteriAdi:   utils.StringValue(order.MusteriAdi),
		DesenKodu:    utils.StringValue(order.DDesenKodu),
		ModelKodu:    utils.StringValue(order.DModelKodu),
		SacakTipi:    utils.StringValue(order.SacakTipi),
		Aciklama:     utils.StringValue(order.Aciklama),
		UretimTarihi: utils.TimeValue(order.SiparisTarihi),
		TerminTarihi: utils.TimeValue(order.TerminTarihi),
	}, nil
}

func (r *OrderRepositoryImpl) AddUretimUploads(tx *gorm.DB, uretimID int, url string) error {

	res := tx.Create(&UretimUploads{
		ID:       uuid.New().String(),
		UretimID: uretimID,
		Url:      url,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *OrderRepositoryImpl) DeleteUretim(id int) *appErrors.Error {

	if err := r.db.Model(&UretimModel{}).Where("KEYNUMBER = ?", id).Delete(&UretimModel{}).Error; err != nil {
		logger.Logger.Error("veri tabanindan uretimi silerken hata olustu", zap.Error(err))
		return &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: "siparis uretim bilgileri silinirken bir hata olustu",
		}
	}

	return nil
}

func (r *OrderRepositoryImpl) GetUretimUploadsPath(tx *gorm.DB, uretimID int) ([]string, *appErrors.Error) {

	var uretimUploads []UretimUploads
	if err := tx.Where("uretim_id = ?", uretimID).Find(&uretimUploads).Error; err != nil {
		logger.Logger.Error("veri tabanindan uretim resimlerini getirirken hata olustu", zap.Error(err))
		return nil, &appErrors.Error{
			Code:    http.StatusInternalServerError,
			Message: "siparis uretim bilgileri getirilirken bir hata olustu",
		}
	}

	var paths []string
	for _, uretimUpload := range uretimUploads {
		paths = append(paths, uretimUpload.Url)
	}

	return paths, nil
}

// func (r *OrderRepositoryImpl) UpdateUretim(tx *gorm.DB, uretim UretimUpdateRequest) *appErrors.Error {

// 	updates := map[string]interface{}{}

// 	if uretim.Kullanici != nil {
// 		updates["KULLANICI"] = *uretim.Kullanici
// 	}
// 	if uretim.Miktar != nil {
// 		updates["MIKTAR"] = *uretim.Miktar
// 	}
// 	if uretim.UretimDurum != nil {
// 		updates["URETIM_DURUM"] = *uretim.UretimDurum
// 	}
// 	if uretim.UretimYeri != nil {
// 		updates["URETIM_YERI"] = *uretim.UretimYeri
// 	}
// 	if uretim.UretimTarihSaat != nil {
// 		updates["URETIM_TARIH_SAAT"] = *uretim.UretimTarihSaat
// 	}

// 	if err := tx.Model(&UretimModel{}).Where("KEYNUMBER = ?", uretim.KeyNumber).Updates(updates).Error; err != nil {
// 		return &appErrors.Error{
// 			Code:    http.StatusInternalServerError,
// 			Message: "siparis uretim bilgileri guncellenirken bir hata olustu",
// 		}
// 	}

// 	return nil
// }
