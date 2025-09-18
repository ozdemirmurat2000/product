package defotanim

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

type DefoTanimRequest struct {
	UretimYeri string `json:"uretimYeri" binding:"required"`
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
