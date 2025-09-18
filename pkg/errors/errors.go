package appErrors

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

const (
	ERR_UNKNOWN      = "bilinmeyen bir hata olustu lutfen daha sonra tekrar deneyin"
	ERR_IMAGE_UPLOAD = "resim yukleme basarisiz"
	ERR_IMAGE_DELETE = "resim silme basarisiz"
)
