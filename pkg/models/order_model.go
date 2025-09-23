package models

import (
	"encoding/base64"
	"mime/multipart"
	"productApp/pkg/utils"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type OrderModel struct {
	KeyNumber            int        `gorm:"column:KEYNUMBER;primaryKey;autoIncrement" json:"keyNumber"`
	SiparisNo            *string    `gorm:"column:SIPARIS_NO" json:"siparisNo"`
	MusteriKodu          *string    `gorm:"column:MUSTERI_KODU" json:"musteriKodu"`
	MusteriAdi           *string    `gorm:"column:MUSTERI_ADI" json:"musteriAdi"`
	SiparisTarihi        *time.Time `gorm:"column:SIPARIS_TARIHI" json:"siparisTarihi"`
	Aciklama             *string    `gorm:"column:ACIKLAMA" json:"aciklama"`
	DDesenKodu           *string    `gorm:"column:D_DESEN_KODU" json:"dDesenKodu"`
	DDesenAciklama       *string    `gorm:"column:D_DESEN_ACIKLAMA" json:"dDesenAciklama"`
	DModelKodu           *string    `gorm:"column:D_MODEL_KODU" json:"dModelKodu"`
	DSiklik1             *float64   `gorm:"column:D_SIKLIK_1" json:"dSiklik1"`
	DSiklik2             *float64   `gorm:"column:D_SIKLIK_2" json:"dSiklik2"`
	DSacak               *string    `gorm:"column:D_SACAK" json:"dSacak"`
	DMamulEbat           *string    `gorm:"column:D_MAMUL_EBAT" json:"dMamulEbat"`
	DDokumaCesidi        *string    `gorm:"column:D_DOKUMA_CESIDI" json:"dDokumaCesidi"`
	DGramaj              *float64   `gorm:"column:D_GRAMAJ" json:"dGramaj"`
	CTarakEni            *float64   `gorm:"column:C_TARAK_ENI" json:"cTarakEni"`
	CNumara              *string    `gorm:"column:C_NUMARA" json:"cNumara"`
	CRenk                *string    `gorm:"column:C_RENK" json:"cRenk"`
	CSiklik              *float64   `gorm:"column:C_SIKLIK" json:"cSiklik"`
	CSacakDahilBoy       *float64   `gorm:"column:C_SACAK_DAHIL_BOY" json:"cSacakDahilBoy"`
	CSarf                *float64   `gorm:"column:C_SARF" json:"cSarf"`
	AAtkiIplikNo         *string    `gorm:"column:A_ATKI_IPLIK_NO" json:"aAtkiIplikNo"`
	ATex                 *float64   `gorm:"column:A_TEX" json:"aTex"`
	AAtkiKolu            *int       `gorm:"column:A_ATKI_KOLU" json:"aAtkiKolu"`
	ASiklik              *float64   `gorm:"column:A_SIKLIK" json:"aSiklik"`
	AHamBoy              *float64   `gorm:"column:A_HAM_BOY" json:"aHamBoy"`
	Renk1                *string    `gorm:"column:RENK_1" json:"renk1"`
	Renk2                *string    `gorm:"column:RENK_2" json:"renk2"`
	Renk3                *string    `gorm:"column:RENK_3" json:"renk3"`
	Renk4                *string    `gorm:"column:RENK_4" json:"renk4"`
	Renk5                *string    `gorm:"column:RENK_5" json:"renk5"`
	Renk6                *string    `gorm:"column:RENK_6" json:"renk6"`
	Renk7                *string    `gorm:"column:RENK_7" json:"renk7"`
	Renk8                *string    `gorm:"column:RENK_8" json:"renk8"`
	RenkAdi1             *string    `gorm:"column:RENK_ADI_1" json:"renkAdi1"`
	RenkAdi2             *string    `gorm:"column:RENK_ADI_2" json:"renkAdi2"`
	RenkAdi3             *string    `gorm:"column:RENK_ADI_3" json:"renkAdi3"`
	RenkAdi4             *string    `gorm:"column:RENK_ADI_4" json:"renkAdi4"`
	RenkAdi5             *string    `gorm:"column:RENK_ADI_5" json:"renkAdi5"`
	RenkAdi6             *string    `gorm:"column:RENK_ADI_6" json:"renkAdi6"`
	RenkAdi7             *string    `gorm:"column:RENK_ADI_7" json:"renkAdi7"`
	RenkAdi8             *string    `gorm:"column:RENK_ADI_8" json:"renkAdi8"`
	ARenkCM1             *float64   `gorm:"column:A_RENK_CM_1" json:"aRenkCM1"`
	ARenkCM2             *float64   `gorm:"column:A_RENK_CM_2" json:"aRenkCM2"`
	ARenkCM3             *float64   `gorm:"column:A_RENK_CM_3" json:"aRenkCM3"`
	ARenkCM4             *float64   `gorm:"column:A_RENK_CM_4" json:"aRenkCM4"`
	ARenkCM5             *float64   `gorm:"column:A_RENK_CM_5" json:"aRenkCM5"`
	ARenkCM6             *float64   `gorm:"column:A_RENK_CM_6" json:"aRenkCM6"`
	ARenkCM7             *float64   `gorm:"column:A_RENK_CM_7" json:"aRenkCM7"`
	ARenkCM8             *float64   `gorm:"column:A_RENK_CM_8" json:"aRenkCM8"`
	ARenkYZ1             *float64   `gorm:"column:A_RENK_YZ_1" json:"aRenkYZ1"`
	ARenkYZ2             *float64   `gorm:"column:A_RENK_YZ_2" json:"aRenkYZ2"`
	ARenkYZ3             *float64   `gorm:"column:A_RENK_YZ_3" json:"aRenkYZ3"`
	ARenkYZ4             *float64   `gorm:"column:A_RENK_YZ_4" json:"aRenkYZ4"`
	ARenkYZ5             *float64   `gorm:"column:A_RENK_YZ_5" json:"aRenkYZ5"`
	ARenkYZ6             *float64   `gorm:"column:A_RENK_YZ_6" json:"aRenkYZ6"`
	ARenkYZ7             *float64   `gorm:"column:A_RENK_YZ_7" json:"aRenkYZ7"`
	ARenkYZ8             *float64   `gorm:"column:A_RENK_YZ_8" json:"aRenkYZ8"`
	ARenkSarf1           *float64   `gorm:"column:A_RENK_SARF_1" json:"aRenkSarf1"`
	ARenkSarf2           *float64   `gorm:"column:A_RENK_SARF_2" json:"aRenkSarf2"`
	ARenkSarf3           *float64   `gorm:"column:A_RENK_SARF_3" json:"aRenkSarf3"`
	ARenkSarf4           *float64   `gorm:"column:A_RENK_SARF_4" json:"aRenkSarf4"`
	ARenkSarf5           *float64   `gorm:"column:A_RENK_SARF_5" json:"aRenkSarf5"`
	ARenkSarf6           *float64   `gorm:"column:A_RENK_SARF_6" json:"aRenkSarf6"`
	ARenkSarf7           *float64   `gorm:"column:A_RENK_SARF_7" json:"aRenkSarf7"`
	ARenkSarf8           *float64   `gorm:"column:A_RENK_SARF_8" json:"aRenkSarf8"`
	SiparisMiktari       *float64   `gorm:"column:SIPARIS_MIKTARI" json:"siparisMiktari"`
	SacakTipi            *string    `gorm:"column:SACAK_TIPI" json:"sacakTipi"`
	Fiyat                *float64   `gorm:"column:FIYAT" json:"fiyat"`
	FiyatTipi            *string    `gorm:"column:FIYAT_TIPI" json:"fiyatTipi"`
	TerminTarihi         *time.Time `gorm:"column:TERMIN_TARIHI" json:"terminTarihi"`
	SiparisKalemAciklama *string    `gorm:"column:SIPARIS_KALEM_ACIKLAMA" json:"siparisKalemAciklama"`
	YikamaVarMi          *bool      `gorm:"column:YIKAMA_VARMI" json:"yikamaVarMi"`
	YikamaTipi           *string    `gorm:"column:YIKAMA_TIPI" json:"yikamaTipi"`
	YikamaAciklama       *string    `gorm:"column:YIKAMA_ACIKLAMA" json:"yikamaAciklama"`
	EtiketVarMi          *bool      `gorm:"column:ETIKET_VARMI" json:"etiketVarMi"`
	EtiketAciklama       *string    `gorm:"column:ETIKET_ACIKLAMA" json:"etiketAciklama"`
	EtiketResim          []byte     `gorm:"column:ETIKET_RESIM" json:"etiketResim"`
	PaketVarMi           *bool      `gorm:"column:PAKET_VARMI" json:"paketVarMi"`
	PaketAciklama        *string    `gorm:"column:PAKET_ACIKLAMA" json:"paketAciklama"`
	PaketResim           []byte     `gorm:"column:PAKET_RESIM" json:"paketResim"`
	KoliVarMi            *bool      `gorm:"column:KOLI_VARMI" json:"koliVarMi"`
	KoliAciklama         *string    `gorm:"column:KOLI_ACIKLAMA" json:"koliAciklama"`
	KoliResim            []byte     `gorm:"column:KOLI_RESIM" json:"koliResim"`
	Dokuma               int        `gorm:"column:DOKUMA;default:0" json:"dokuma"`
	Yikama               int        `gorm:"column:YIKAMA;default:0" json:"yikama"`
	KaliteKontrol        int        `gorm:"column:KALITEKONTROL;default:0" json:"kaliteKontrol"`
	Paketleme            int        `gorm:"column:PAKETLEME;default:0" json:"paketleme"`
	Sevkiyat             int        `gorm:"column:SEVKIYAT;default:0" json:"sevkiyat"`
}

func (*OrderModel) TableName() string {
	return "TBL_SIPARIS_LISTE"
}

type OrderListRequest struct {
	Process string `json:"process"`
}

func (m *OrderModel) ToOrderResponse(siparisResim string) *OrderResponse {
	return &OrderResponse{
		SiparisBilgileriResponse: SiparisBilgileriResponse{
			ID:             strconv.Itoa(m.KeyNumber),
			SiparisNo:      utils.StringValue(m.SiparisNo),
			MusteriKodu:    utils.StringValue(m.MusteriKodu),
			MusteriAdi:     utils.StringValue(m.MusteriAdi),
			SiparisTarihi:  utils.TimeValue(m.SiparisTarihi),
			Aciklama:       utils.StringValue(m.Aciklama),
			SacakTipi:      utils.StringValue(m.SacakTipi),
			TerminTarihi:   utils.TimeValue(m.TerminTarihi),
			SiparisMiktari: utils.Float64Value(m.SiparisMiktari),
			Dokuma:         m.Dokuma,
			Yikama:         m.Yikama,
			KaliteKontrol:  m.KaliteKontrol,
			Paketleme:      m.Paketleme,
			Sevkiyat:       m.Sevkiyat,
			SiparisResim:   siparisResim,
		},
		DesenBilgileriResponse: DesenBilgileriResponse{
			DDesenKodu:     utils.StringValue(m.DDesenKodu),
			DDesenAciklama: utils.StringValue(m.DDesenAciklama),
			DModelKodu:     utils.StringValue(m.DModelKodu),
			DSiklik1:       utils.Float64Value(m.DSiklik1),
			DSiklik2:       utils.Float64Value(m.DSiklik2),
			DSacak:         utils.StringValue(m.DSacak),
			DMamulEbat:     utils.StringValue(m.DMamulEbat),
			DDokumaCesidi:  utils.StringValue(m.DDokumaCesidi),
			DGramaj:        utils.Float64Value(m.DGramaj),
		},
		CozguBilgileriResponse: CozguBilgileriResponse{
			TarakEni:      utils.Float64Value(m.CTarakEni),
			Numara:        utils.StringValue(m.CNumara),
			Renk:          utils.StringValue(m.CRenk),
			Siklik:        utils.Float64Value(m.CSiklik),
			SacakDahilBoy: utils.Float64Value(m.CSacakDahilBoy),
			Sarf:          utils.Float64Value(m.CSarf),
			Tex:           utils.Float64Value(m.ATex),
		},
		AtkiBilgileriResponse: AtkiBilgileriResponse{
			AtkiIplikNumarasi: utils.StringValue(m.AAtkiIplikNo),
			Tex:               utils.Float64Value(m.ATex),
			AtkiKolu:          utils.IntValue(m.AAtkiKolu),
			Siklik:            utils.Float64Value(m.ASiklik),
			HamBoy:            utils.Float64Value(m.AHamBoy),
		},
		RenkBilgileriResponse: []RenkBilgileriResponse{
			{
				RenkKodu:   utils.StringValue(m.Renk1),
				RenkAdi:    utils.StringValue(m.RenkAdi1),
				RenkGramaj: utils.Float64Value(m.ARenkSarf1),
				RenkCm:     utils.Float64Value(m.ARenkCM1),
			},
			{
				RenkKodu:   utils.StringValue(m.Renk2),
				RenkAdi:    utils.StringValue(m.RenkAdi2),
				RenkGramaj: utils.Float64Value(m.ARenkSarf2),
				RenkCm:     utils.Float64Value(m.ARenkCM2),
			},
			{
				RenkKodu:   utils.StringValue(m.Renk3),
				RenkAdi:    utils.StringValue(m.RenkAdi3),
				RenkGramaj: utils.Float64Value(m.ARenkSarf3),
				RenkCm:     utils.Float64Value(m.ARenkCM3),
			},
			{
				RenkKodu:   utils.StringValue(m.Renk4),
				RenkAdi:    utils.StringValue(m.RenkAdi4),
				RenkGramaj: utils.Float64Value(m.ARenkSarf4),
				RenkCm:     utils.Float64Value(m.ARenkCM4),
			},
			{
				RenkKodu:   utils.StringValue(m.Renk5),
				RenkAdi:    utils.StringValue(m.RenkAdi5),
				RenkGramaj: utils.Float64Value(m.ARenkSarf5),
				RenkCm:     utils.Float64Value(m.ARenkCM5),
			},
			{
				RenkKodu:   utils.StringValue(m.Renk6),
				RenkAdi:    utils.StringValue(m.RenkAdi6),
				RenkGramaj: utils.Float64Value(m.ARenkSarf6),
				RenkCm:     utils.Float64Value(m.ARenkCM6),
			},
			{
				RenkKodu:   utils.StringValue(m.Renk7),
				RenkAdi:    utils.StringValue(m.RenkAdi7),
				RenkGramaj: utils.Float64Value(m.ARenkSarf7),
				RenkCm:     utils.Float64Value(m.ARenkCM7),
			},
			{
				RenkKodu:   utils.StringValue(m.Renk8),
				RenkAdi:    utils.StringValue(m.RenkAdi8),
				RenkGramaj: utils.Float64Value(m.ARenkSarf8),
				RenkCm:     utils.Float64Value(m.ARenkCM8),
			},
		},
		IslemBilgileriResponse: []IslemBilgileriResponse{
			{
				IslemAdi:      "Yıkama",
				IslemAktifMi:  utils.BoolValue(m.YikamaVarMi),
				IslemAciklama: utils.StringValue(m.YikamaAciklama),
				IslemResim:    nil,
			},
			{
				IslemAdi:      "Etiket",
				IslemAktifMi:  utils.BoolValue(m.EtiketVarMi),
				IslemAciklama: utils.StringValue(m.EtiketAciklama),
				IslemResim:    strPtr(base64.StdEncoding.EncodeToString(m.EtiketResim)),
			},
			{
				IslemAdi:      "Paketleme",
				IslemAktifMi:  utils.BoolValue(m.PaketVarMi),
				IslemAciklama: utils.StringValue(m.PaketAciklama),
				IslemResim:    strPtr(base64.StdEncoding.EncodeToString(m.PaketResim)),
			},
			{
				IslemAdi:      "Koli",
				IslemAktifMi:  utils.BoolValue(m.KoliVarMi),
				IslemAciklama: utils.StringValue(m.KoliAciklama),
				IslemResim:    strPtr(base64.StdEncoding.EncodeToString(m.KoliResim)),
			},
		},
	}

}

type SiparisBilgileriResponse struct {
	ID             string    `json:"id"`
	SiparisNo      string    `json:"siparisNo"`
	MusteriKodu    string    `json:"musteriKodu"`
	MusteriAdi     string    `json:"musteriAdi"`
	SiparisTarihi  time.Time `json:"siparisTarihi"`
	Aciklama       string    `json:"aciklama"`
	SacakTipi      string    `json:"sacakTipi"`
	TerminTarihi   time.Time `json:"terminTarihi"`
	SiparisMiktari float64   `json:"siparisMiktari"`
	Dokuma         int       `json:"dokuma"`
	Yikama         int       `json:"yikama"`
	KaliteKontrol  int       `json:"kaliteKontrol"`
	Paketleme      int       `json:"paketleme"`
	Sevkiyat       int       `json:"sevkiyat"`
	SiparisResim   string    `json:"siparisResim"`
}

type DesenBilgileriResponse struct {
	DDesenKodu     string  `json:"dDesenKodu"`
	DDesenAciklama string  `json:"dDesenAciklama"`
	DModelKodu     string  `json:"dModelKodu"`
	DSiklik1       float64 `json:"dSiklik1"`
	DSiklik2       float64 `json:"dSiklik2"`
	DSacak         string  `json:"dSacak"`
	DMamulEbat     string  `json:"dMamulEbat"`
	DDokumaCesidi  string  `json:"dDokumaCesidi"`
	DGramaj        float64 `json:"dGramaj"`
}

type CozguBilgileriResponse struct {
	TarakEni      float64 `json:"tarakEni"`
	Numara        string  `json:"numara"`
	Renk          string  `json:"renk"`
	Siklik        float64 `json:"siklik"`
	SacakDahilBoy float64 `json:"sacakDahilBoy"`
	Sarf          float64 `json:"gramaj"`
	Tex           float64 `json:"tex"`
}

type AtkiBilgileriResponse struct {
	AtkiIplikNumarasi string  `json:"atkiIplikNumarasi"`
	Tex               float64 `json:"tex"`
	AtkiKolu          int     `json:"atkiKolu"`
	Siklik            float64 `json:"siklik"`
	HamBoy            float64 `json:"hamBoy"`
}

type RenkBilgileriResponse struct {
	RenkKodu   string  `json:"renkKodu"`
	RenkAdi    string  `json:"renkAdi"`
	RenkGramaj float64 `json:"renkGramaj"`
	RenkCm     float64 `json:"renkCm"`
}

type IslemBilgileriResponse struct {
	IslemAdi      string  `json:"islemAdi"`
	IslemAktifMi  bool    `json:"islemAktifMi"`
	IslemAciklama string  `json:"islemAciklama"`
	IslemResim    *string `json:"islemResim"`
}

type OrderResponse struct {
	SiparisBilgileriResponse SiparisBilgileriResponse `json:"siparisBilgileriResponse"`
	DesenBilgileriResponse   DesenBilgileriResponse   `json:"desenBilgileriResponse"`
	CozguBilgileriResponse   CozguBilgileriResponse   `json:"cozguBilgileriResponse"`
	AtkiBilgileriResponse    AtkiBilgileriResponse    `json:"atkiBilgileriResponse"`
	RenkBilgileriResponse    []RenkBilgileriResponse  `json:"renkBilgileriResponse"`
	IslemBilgileriResponse   []IslemBilgileriResponse `json:"islemBilgileriResponse"`
}

func strPtr(s string) *string {
	return &s
}

type OrderSummaryResponse struct {
	SiparisNo       string  `json:"siparisNo"`
	MusteriAdi      string  `json:"musteriAdi"`
	DesenKodu       string  `json:"desenKodu"`
	DesenAciklamasi string  `json:"desenAciklamasi"`
	ModelKodu       string  `json:"modelKodu"`
	SiparisMiktari  float64 `json:"siparisMiktari"`
	DokumaMiktar    int     `json:"dokumaMiktar"`
	YikamaMiktar    int     `json:"yikamaMiktar"`
	KaliteMiktar    int     `json:"kaliteMiktar"`
	PaketlemeMiktar int     `json:"paketlemeMiktar"`
	SevkiyatMiktar  int     `json:"sevkiyatMiktar"`
}

func (o *OrderModel) ToOrderSummaryResponse(dokumaMiktari int, yikamaMiktari int, kaliteMiktari int, paketlemeMiktari int, sevkiyatMiktari int) OrderSummaryResponse {
	return OrderSummaryResponse{
		SiparisNo:       utils.StringValue(o.SiparisNo),
		MusteriAdi:      utils.StringValue(o.MusteriAdi),
		DesenKodu:       utils.StringValue(o.DDesenKodu),
		DesenAciklamasi: utils.StringValue(o.DDesenAciklama),
		ModelKodu:       utils.StringValue(o.DModelKodu),
		DokumaMiktar:    dokumaMiktari,
		YikamaMiktar:    yikamaMiktari,
		KaliteMiktar:    kaliteMiktari,
		PaketlemeMiktar: paketlemeMiktari,
		SevkiyatMiktar:  sevkiyatMiktari,
	}
}

type UretimModel struct {
	KeyNumber       int        `gorm:"column:KEYNUMBER;primaryKey;autoIncrement"`
	SiparisNo       *string    `gorm:"column:SIPARIS_NO"`
	UretimDurum     *string    `gorm:"column:URETIM_DURUM"`
	UretimYeri      *string    `gorm:"column:URETIM_YERI"`
	Miktar          *float64   `gorm:"column:MIKTAR"`
	Kullanici       *string    `gorm:"column:KULLANICI"`
	UretimTarihSaat *time.Time `gorm:"column:URETIM_TARIH_SAAT"`
}

type UretimUpdateRequest struct {
	KeyNumber       int      `json:"keyNumber" binding:"required"`
	UretimDurum     *string  `json:"uretimDurum"`
	UretimYeri      *string  `json:"uretimYeri"`
	Miktar          *float64 `json:"miktar"`
	Kullanici       *string
	UretimTarihSaat *time.Time `json:"uretimTarihSaat"`
}

func (u *UretimModel) TableName() string {
	return "TBL_URETIM"
}

type UretimResponse struct {
	KeyNumber       int       `json:"keyNumber"`
	SiparisNo       string    `json:"siparisNo"`
	UretimDurum     string    `json:"uretimDurum"`
	UretimYeri      string    `json:"uretimYeri"`
	Miktar          float64   `json:"miktar"`
	Kullanici       string    `json:"kullanici"`
	UretimTarihSaat time.Time `json:"uretimTarihSaat"`
	ImageURL        []string  `json:"imageURL"`
}

type SiparisUretimResponse struct {
	SiparisNo           string              `json:"siparisNo"`
	MusteriAdi          string              `json:"musteriAdi"`
	DesenKodu           string              `json:"desenKodu"`
	ModelKodu           string              `json:"modelKodu"`
	SacakTipi           string              `json:"sacakTipi"`
	Aciklama            string              `json:"aciklama"`
	UretimTarihi        time.Time           `json:"uretimTarihi"`
	TerminTarihi        time.Time           `json:"terminTarihi"`
	DefoTanim           []DefoTanimResponse `json:"defoTanim"`
	UretimResponse      []UretimResponse    `json:"uretimResponse"`
	SiparisMiktari      float64             `json:"siparisMiktari"`
	ToplamMiktar        float64             `json:"toplamMiktar"`
	SaglamMiktar        float64             `json:"saglamMiktar"`
	DefoMiktar          float64             `json:"defoMiktar"`
	UretimChartResponse UretimChartResponse `json:"uretimChartResponse"`
	ChartResponse       []ChartResponse     `json:"chartResponse"`
}

type UretimChartResponse struct {
	DokumaMiktar          float64 `json:"dokumaMiktar"`
	DokumaColorHexCode    string  `json:"dokumaColorHexCode"`
	YikamaMiktar          float64 `json:"yikamaMiktar"`
	YikamaColorHexCode    string  `json:"yikamaColorHexCode"`
	KaliteMiktar          float64 `json:"kaliteMiktar"`
	KaliteColorHexCode    string  `json:"kaliteColorHexCode"`
	PaketlemeMiktar       float64 `json:"paketlemeMiktar"`
	PaketlemeColorHexCode string  `json:"paketlemeColorHexCode"`
	SevkiyatMiktar        float64 `json:"sevkiyatMiktar"`
	SevkiyatColorHexCode  string  `json:"sevkiyatColorHexCode"`
}

type ChartResponse struct {
	ColorHexCode string  `json:"colorHexCode"`
	Percent      float64 `json:"percent"`
	Name         string  `json:"name"`
}

func (u *UretimModel) ToUretimResponse() UretimResponse {
	return UretimResponse{
		KeyNumber:       u.KeyNumber,
		SiparisNo:       utils.StringValue(u.SiparisNo),
		UretimDurum:     utils.StringValue(u.UretimDurum),
		UretimYeri:      utils.StringValue(u.UretimYeri),
		Miktar:          utils.Float64Value(u.Miktar),
		Kullanici:       utils.StringValue(u.Kullanici),
		UretimTarihSaat: utils.TimeValue(u.UretimTarihSaat),
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
	ID       string `gorm:"type:uniqueidentifier;default:newsequentialid();primaryKey"`
	UretimID int    `gorm:"column:uretim_id"`
	Url      string `gorm:"column:image_url"`
}

func (u *UretimUploads) TableName() string {
	return "TBL_URETIM_UPLOADS"
}

type DefoTanimModel struct {
	ID             int    `gorm:"column:KEYNUMBER;primaryKey"`
	UretimYeri     string `gorm:"column:URETIM_YERI"`
	DefoIsmi       string `gorm:"column:DEFO_ISMI"`
	Varsayilan     bool   `gorm:"column:VARSAYIILAN"`
	ResimZorunlumu bool   `gorm:"column:RESIM_ZORUNLUMU"`
}

func (m *DefoTanimModel) TableName() string {
	return "TBL_DEFOTANIM"
}

type DefoTanimResponse struct {
	ID             int    `json:"id"`
	UretimYeri     string `json:"uretimYeri"`
	DefoIsmi       string `json:"defoIsmi"`
	Varsayilan     bool   `json:"varsayilan"`
	ResimZorunlumu bool   `json:"resimZorunlumu"`
}

func (m *DefoTanimModel) ToDefoTanimResponse() DefoTanimResponse {
	return DefoTanimResponse{
		ID:             m.ID,
		UretimYeri:     m.UretimYeri,
		DefoIsmi:       m.DefoIsmi,
		Varsayilan:     m.Varsayilan,
		ResimZorunlumu: m.ResimZorunlumu,
	}
}

type CustomerOrdersResponse struct {
	MusteriKodu        string `json:"musteriKodu"`
	MusteriAdi         string `json:"musteriAdi"`
	AktifSiparisSayisi int    `json:"aktifSiparisSayisi"`
}

type CustomerOrdersModel struct {
	MusteriKodu        string `gorm:"column:MUSTERI_KODU"`
	MusteriAdi         string `gorm:"column:MUSTERI_ADI"`
	AktifSiparisSayisi int    `gorm:"column:SIPARISMIKTARI"`
}

func (m *CustomerOrdersModel) ToCustomerOrdersResponse() CustomerOrdersResponse {
	return CustomerOrdersResponse{
		MusteriKodu:        m.MusteriKodu,
		MusteriAdi:         m.MusteriAdi,
		AktifSiparisSayisi: m.AktifSiparisSayisi,
	}
}

type OrderSummaryModel struct {
	SiparisNo      string  `gorm:"column:SIPARIS_NO"`
	MusteriAdi     string  `gorm:"column:MUSTERI_ADI"`
	DesenKodu      string  `gorm:"column:D_DESEN_KODU"`
	DesenAciklama  string  `gorm:"column:D_DESEN_ACIKLAMA"`
	ModelKodu      string  `gorm:"column:D_MODEL_KODU"`
	SiparisMiktari float64 `gorm:"column:SIPARIS_MIKTARI"`
}

func (m *OrderSummaryModel) ToOrderSummaryResponse(dokumaMiktari int, yikamaMiktari int, kaliteMiktari int, paketlemeMiktari int, sevkiyatMiktari int) OrderSummaryResponse {
	return OrderSummaryResponse{
		SiparisNo:       m.SiparisNo,
		MusteriAdi:      m.MusteriAdi,
		DesenKodu:       m.DesenKodu,
		DesenAciklamasi: m.DesenAciklama,
		ModelKodu:       m.ModelKodu,
		SiparisMiktari:  m.SiparisMiktari,
		DokumaMiktar:    dokumaMiktari,
		YikamaMiktar:    yikamaMiktari,
		KaliteMiktar:    kaliteMiktari,
		PaketlemeMiktar: paketlemeMiktari,
		SevkiyatMiktar:  sevkiyatMiktari,
	}
}

type ModelResimModel struct {
	Kodu      string     `gorm:"column:KODU;not null;primaryKey"`
	ResimURL  string     `gorm:"column:RESIM_URL;not null"`
	CreatedAt *time.Time `gorm:"column:CREATED_AT;not null;default:getdate()"`
	UpdatedAt *time.Time `gorm:"column:UPDATED_AT;not null;default:getdate()"`
}

func (m *ModelResimModel) TableName() string {
	return "TBL_YENI_MODEL_RESIM"
}
func (m *ModelResimModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	m.UpdatedAt = &now
	return nil
}
