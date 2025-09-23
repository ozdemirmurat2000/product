//go:build wireinject
// +build wireinject

package injection

import (
	"productApp/internal/auth"
	defotanim "productApp/internal/defo_tanim"
	order "productApp/internal/order"
	uretim "productApp/internal/uretim"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeAuthController(db *gorm.DB) auth.AuthController {
	wire.Build(auth.NewAuthRepository, auth.NewAuthService, auth.NewAuthController)
	return &auth.AuthControllerImpl{}

}

func InitializeOrderController(db *gorm.DB) order.IOrderController {
	wire.Build(order.NewOrderRepositoryImpl, order.NewOrderServiceImpl, order.NewOrderControllerImpl)
	return &order.OrderControllerImpl{}
}

func InitializeDefoTanimController(db *gorm.DB) defotanim.IDefoTanimController {
	wire.Build(defotanim.NewDefoTanimRepositoryImpl, defotanim.NewDefoTanimServiceImpl, defotanim.NewDefoTanimControllerImpl)
	return &defotanim.DefoTanimControllerImpl{}
}

func InitializeUretimController(db *gorm.DB) uretim.IUretimController {
	wire.Build(uretim.NewUretimRepositoryImpl, uretim.NewUretimServiceImpl, uretim.NewUretimControllerImpl, defotanim.NewDefoTanimRepositoryImpl, defotanim.NewDefoTanimServiceImpl)
	return &uretim.UretimControllerImpl{}
}
