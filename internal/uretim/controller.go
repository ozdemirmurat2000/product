package uretim

import (
	"fmt"
	"net/http"
	"productApp/pkg/jwt"
	"productApp/pkg/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	defo "productApp/internal/defo_tanim"
)

type IUretimController interface {
	GetUretimList(ctx *gin.Context)
	DeleteUretim(ctx *gin.Context)
	AddUretim(ctx *gin.Context)
}

type UretimControllerImpl struct {
	service          IUretimService
	defoTanimService defo.IDefoTanimService
}

func NewUretimControllerImpl(service IUretimService, defoTanimService defo.IDefoTanimService) IUretimController {
	return &UretimControllerImpl{service: service, defoTanimService: defoTanimService}
}

//	Uretim
//
// @Summary      This point user log in request
// @Tags         uretim
// @Accept       json
// @Produce      json
// @Param        siparisNo query string true "siparisNo"
// @Success      200  {object} response.SuccessResponseModel(data={[]UretimResponse})
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /uretim [get]
func (c *UretimControllerImpl) GetUretimList(ctx *gin.Context) {
	siparisNo := ctx.Query("siparisNo")

	if siparisNo == "" {

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("siparisNo zorunlu"))
		return

	}

	request := UretimRequest{
		SiparisNo: siparisNo,
	}

	orderList, err := c.service.GetUretimList(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}
	if orderList == nil {
		ctx.JSON(http.StatusOK, response.ErrorResponse("uretim bulunamadi"))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("uretim list", orderList))
}

//	Uretim
//
// @Summary      This point user log in request
// @Tags         uretim
// @Accept       json
// @Produce      json
// @Param        id query int true "id"
// @Success      200  {object} response.SuccessResponseModel(data={[]UretimResponse})
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /uretim [delete]
func (c *UretimControllerImpl) DeleteUretim(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("id zorunlu"))
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}

	if err := c.service.DeleteUretim(idInt); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("uretim silindi", nil))
}

// @Summary      Upload images and order data
// @Tags         uretim
// @Accept       multipart/form-data
// @Produce      json
// @Param        siparisNo formData string true "Siparis no"
// @Param        uretimDurum formData string true "Uretim durum"
// @Param        uretimYeri formData string true "Uretim yeri"
// @Param        miktari formData int true "Miktari"
// @Param        images formData file false "Resimler" multiple
// @Success      200  {object} response.SuccessResponseModel
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /uretim [post]
func (c *UretimControllerImpl) AddUretim(ctx *gin.Context) {
	var request UretimAddRequest

	claims, _ := ctx.MustGet("claims").(*jwt.Claims)

	fmt.Println(claims)

	err := ctx.Request.ParseMultipartForm(20 << 20)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}

	request.SiparisNo = ctx.Request.FormValue("siparisNo")
	request.UretimDurum = ctx.Request.FormValue("uretimDurum")
	request.UretimYeri = ctx.Request.FormValue("uretimYeri")
	request.Miktar, err = strconv.Atoi(ctx.Request.FormValue("miktari"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}

	defo, err := c.defoTanimService.GetDefoByName(ctx.Request.FormValue("uretimDurum"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}

	if defo.ResimZorunlumu {

		request.File = ctx.Request.MultipartForm.File["images"]

		if len(request.File) == 0 {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse("resim zorunlu"))
			return
		}
	}

	request.UretimTarihi = time.Now()
	request.Kullanici = claims.Username

	fmt.Println(request.Kullanici)

	if request.SiparisNo == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("siparisNo zorunlu"))
		return
	}
	if request.UretimDurum == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("uretimDurum zorunlu"))
		return
	}
	if request.UretimYeri == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("uretimYeri zorunlu"))
		return
	}

	if request.Kullanici == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("kullanici zorunlu"))
		return
	}

	if err := c.service.AddUretim(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("uretim eklendi", nil))
}
