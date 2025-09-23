package image_storage

import (
	"net/http"
	appErrors "productApp/pkg/errors"
)

var (
	err_image_is_too_large        = appErrors.New(http.StatusInternalServerError, "resim boyutu çok büyük")
	err_image_type_is_not_allowed = appErrors.New(http.StatusInternalServerError, "resim tipi desteklenmiyor")
	err_folder_creation_failed    = appErrors.New(http.StatusInternalServerError, "klasör oluşturulamadı")
	err_image_delete_failed       = appErrors.New(http.StatusInternalServerError, "resim silme başarısız")
)
