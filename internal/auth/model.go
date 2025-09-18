package auth

type UserModel struct {
	ID          int    `gorm:"column:KEYNUMBER;primaryKey"`
	UserCode    string `gorm:"column:KULLANICI_KODU"`
	UserName    string `gorm:"column:KULLANICI_ADI"`
	Password    string `gorm:"column:PAROLA"`
	IsAdmin     bool   `gorm:"column:ADMIN"`
	Phone       string `gorm:"column:GSM"`
	Email       string `gorm:"column:MAIL"`
	ShowMenu    bool   `gorm:"column:ACILIS_MENU_GOSTER"`
	LandingPage string `gorm:"column:ACILIS_SAYFASI"`
	CardID1     string `gorm:"column:KART_ID_1"`
	CardID2     string `gorm:"column:KART_ID_2"`
	CardID3     string `gorm:"column:KART_ID_3"`
	Authority   string `gorm:"column:YETKI"`
	Weaving     bool   `gorm:"column:DOKUMA"`
	Wash        bool   `gorm:"column:YIKAMA"`
	Quality     bool   `gorm:"column:KALITE"`
	Packaging   bool   `gorm:"column:PAKETLEME"`
	Shipment    bool   `gorm:"column:SEVKIYAT"`
}

func (u *UserModel) TableName() string {
	return "DG_TNM_KULLANICI"
}

type LoginResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"accessToken"`
	RefreshToken string       `json:"refreshToken"`
}

type LoginRequest struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *UserModel) ToUserResponse() UserResponse {
	return UserResponse{
		ID:          u.ID,
		UserName:    u.UserName,
		IsAdmin:     u.IsAdmin,
		ShowMenu:    u.ShowMenu,
		LandingPage: u.LandingPage,
		CardID1:     u.CardID1,
		CardID2:     u.CardID2,
		CardID3:     u.CardID3,
		Authority:   u.Authority,
		Weaving:     u.Weaving,
		Wash:        u.Wash,
		Quality:     u.Quality,
		Packaging:   u.Packaging,
		Shipment:    u.Shipment,
		Phone:       u.Phone,
		Email:       u.Email,
	}
}

type UserResponse struct {
	ID          int    `json:"id"`
	UserName    string `json:"userName"`
	IsAdmin     bool   `json:"isAdmin"`
	ShowMenu    bool   `json:"showMenu"`
	LandingPage string `json:"landingPage"`
	CardID1     string `json:"cardID1"`
	CardID2     string `json:"cardID2"`
	CardID3     string `json:"cardID3"`
	Authority   string `json:"authority"`
	Weaving     bool   `json:"weaving"`
	Wash        bool   `json:"wash"`
	Quality     bool   `json:"quality"`
	Packaging   bool   `json:"packaging"`
	Shipment    bool   `json:"shipment"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}
