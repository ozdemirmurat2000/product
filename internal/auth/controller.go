package auth

import (
	"errors"
	"net/http"
	appErrors "productApp/pkg/errors"
	"productApp/pkg/response"
	"productApp/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type AuthControllerImpl struct {
	serv AuthService
}

func NewAuthController(s AuthService) AuthController {
	return &AuthControllerImpl{serv: s}
}

//	Auth
//
// @Summary      This point user log in request
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        auth body auth.LoginRequest true "login request params"
// @Success      200  {object} response.SuccessResponseModel(data={auth.LoginResponse})
// @Failure      400  {object} response.ErrorResponseModel
// @Router       /auth/login [post]
func (ctr *AuthControllerImpl) Login(ctx *gin.Context) {

	var request LoginRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(utils.FormatValidationError(err, request)))
		return
	}

	result, err := ctr.serv.Login(request.UserName, request.Password)
	if err != nil {

		if e := new(appErrors.Error); errors.As(err, &e) {
			ctx.JSON(e.Code, response.ErrorResponse(e.Message))
			return

		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(appErrors.ERR_UNKNOWN))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("login success", result))

}
