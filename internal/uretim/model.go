package uretim

import (
	"mime/multipart"
	"time"
)

type UretimModel struct {
	ID           int       `gorm:"column:KEYNUMBER;primaryKey"`
	SiparisNo    string    `gorm:"column:SIPARIS_NO"`
	UretimDurum  string    `gorm:"column:URETIM_DURUM"`
	UretimYeri   string    `gorm:"column:URETIM_YERI"`
	Miktari      float64   `gorm:"column:MIKTAR"`
	Kullanici    string    `gorm:"column:KULLANICI"`
	UretimTarihi time.Time `gorm:"column:URETIM_TARIH_SAAT"`
}

func (m *UretimModel) TableName() string {
	return "TBL_URETIM"
}

type UretimRequest struct {
	SiparisNo string `json:"siparisNo" binding:"required"`
}

type UretimResponse struct {
	ID           int       `json:"id"`
	SiparisNo    string    `json:"siparisNo"`
	UretimDurum  string    `json:"uretimDurum"`
	UretimYeri   string    `json:"uretimYeri"`
	Miktari      float64   `json:"miktari"`
	Kullanici    string    `json:"kullanici"`
	UretimTarihi time.Time `json:"uretimTarihi"`
}

func (m *UretimModel) ToUretimResponse() UretimResponse {
	return UretimResponse{
		ID:           m.ID,
		SiparisNo:    m.SiparisNo,
		UretimDurum:  m.UretimDurum,
		UretimYeri:   m.UretimYeri,
		Miktari:      m.Miktari,
		Kullanici:    m.Kullanici,
		UretimTarihi: m.UretimTarihi,
	}
}

type UretimAddRequest struct {
	File         []*multipart.FileHeader
	SiparisNo    string
	UretimDurum  string
	UretimYeri   string
	Miktar       int
	Kullanici    string
	UretimTarihi time.Time
}

type UretimUploads struct {
	UretimID int    `gorm:"column:uretim_id"`
	Url      string `gorm:"column:image_url"`
}

func (m *UretimUploads) TableName() string {
	return "TBL_URETIM_UPLOADS"
}
