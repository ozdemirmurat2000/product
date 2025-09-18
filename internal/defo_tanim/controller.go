package defotanim

import (
	"net/http"
	"productApp/pkg/response"

	"github.com/gin-gonic/gin"
)

type IDefoTanimController interface {
	GetDefoTanimList(ctx *gin.Context)
}

type DefoTanimControllerImpl struct {
	service IDefoTanimService
}

func NewDefoTanimControllerImpl(service IDefoTanimService) IDefoTanimController {
	return &DefoTanimControllerImpl{service: service}
}

// DefoTanim
//
// @Summary      Get defo tanim list by uretim yeri
// @Tags         defotanim
// @Accept       json
// @Produce      json
// @Param 		 uretimYeri query string true "uretimYeri"
// @Success      200  {object} response.SuccessResponseModel(data={[]DefoTanimResponse})
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /defo [get]
func (c *DefoTanimControllerImpl) GetDefoTanimList(ctx *gin.Context) {

	uretimYeri := ctx.Query("uretimYeri")
	if uretimYeri == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("uretimYeri zorunlu"))
		return
	}

	request := DefoTanimRequest{
		UretimYeri: uretimYeri,
	}

	defoTanimList, err := c.service.GetDefoTanimList(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}
	if defoTanimList == nil {
		ctx.JSON(http.StatusOK, response.ErrorResponse("bu alana ait defo tanim bulunamadi"))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("defo tanim list", defoTanimList))

}
