package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error, obj any) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		fieldError := validationErrors[0]
		fieldName := fieldError.StructField()

		objType := reflect.TypeOf(obj)
		if objType.Kind() == reflect.Ptr {
			objType = objType.Elem()
		}

		field, found := objType.FieldByName(fieldName)
		if !found {
			return fmt.Sprintf("%s geçersiz", fieldName)
		}

		label := field.Tag.Get("label")
		tag := fieldError.Tag()

		switch tag {
		case "required":
			return fmt.Sprintf("%s zorunlu alan", label)
		case "email":
			return "Geçersiz e-posta formatı"
		case "min":
			return fmt.Sprintf("%s en az %s karakter olmalıdır", label, fieldError.Param())
		case "max":
			return fmt.Sprintf("%s en fazla %s olmalıdır", label, fieldError.Param())
		case "alphanumunicode":
			return fmt.Sprintf("%s sadece harf içerebilir", label)
		case "alpha":
			return fmt.Sprintf("%s sadece harf içerebilir", label)
		case "alphanum":
			return fmt.Sprintf("%s sadece harf ve sayı içerebilir", label)
		case "numeric":
			return fmt.Sprintf("%s sadece sayı içerebilir", label)
		default:
			return fmt.Sprintf("%s geçersiz", label)
		}
	}

	return strings.TrimSpace(err.Error())
}
