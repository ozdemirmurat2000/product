package order

import (
	"net/http"
	appErrors "productApp/pkg/errors"
)

var (
	err_customer_order_not_found    = appErrors.New(http.StatusNotFound, "müşteriye ait sipariş bulunamadı")
	err_fetch_customer_order_failed = appErrors.New(http.StatusInternalServerError, "müşteriye ait sipariş getirilirken hata oluştu")
	err_update_etiket_image_url     = appErrors.New(http.StatusInternalServerError, "sipariş etiket resim eklendikten sonra hata oluştu")
	err_update_paketleme_image_url  = appErrors.New(http.StatusInternalServerError, "sipariş paketleme resim eklendikten sonra hata oluştu")
	err_update_koli_image_url       = appErrors.New(http.StatusInternalServerError, "sipariş koli resim eklendikten sonra hata oluştu")
	err_update_renk_image_url       = appErrors.New(http.StatusInternalServerError, "renk resim eklendikten sonra hata oluştu")
)
