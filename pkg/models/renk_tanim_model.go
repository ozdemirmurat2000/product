package models

type RenkTanimModel struct {
	KeyNumber    int      `gorm:"column:KEYNUMBER;primaryKey;autoIncrement"`
	Kodu         *string  `gorm:"column:KODU"`
	IplikNo      *string  `gorm:"column:IPLIK_NO"`
	Adi          *string  `gorm:"column:ADI"`
	KullanimYeri *string  `gorm:"column:KULLANIM_YERI"`
	Termin       *int     `gorm:"column:TERMIN"`
	KritikMiktar *float64 `gorm:"column:KRITIK_MIKTAR"`
	IngAdi       *string  `gorm:"column:ING_ADI"`
	Aciklama     *string  `gorm:"column:ACIKLAMA"`
	Resim        *[]byte  `gorm:"column:RESIM"`
	ResimURL     *string  `gorm:"column:RESIM_URL"`
}

func (*RenkTanimModel) TableName() string {
	return "TBL_RENK_TANIM"
}
