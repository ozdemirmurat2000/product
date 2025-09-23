package order

import (
	"net/http"
	appErrors "productApp/pkg/errors"
)

var (
	err_customer_order_not_found    = appErrors.New(http.StatusNotFound, "müşteriye ait sipariş bulunamadı")
	err_fetch_customer_order_failed = appErrors.New(http.StatusInternalServerError, "müşteriye ait sipariş getirilirken hata oluştu")
)
