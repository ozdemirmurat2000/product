package order

import (
	"net/http"
	"productApp/pkg/jwt"
	"productApp/pkg/models"
	"productApp/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IOrderController interface {
	GetOrderBySiparisID(ctx *gin.Context)
	GetOrderSummaryList(ctx *gin.Context)
	GetOrderUretimBilgileriBySiparisID(ctx *gin.Context)
	AddNewUretim(ctx *gin.Context)
	DeleteUretim(ctx *gin.Context)
	GetCustomerOrdersByIslemAdi(ctx *gin.Context)
	AddNewModelResim(ctx *gin.Context)
}

type OrderControllerImpl struct {
	service IOrderService
}

func NewOrderControllerImpl(service IOrderService) IOrderController {
	return &OrderControllerImpl{service: service}
}

// Get Order By Siparis ID
//
// @Summary      Get order by siparis id
// @Tags         order
// @Accept       json
// @Produce      json
// @Param 		 siparisID query string true "siparis id"
// @Success      200  {object} response.SuccessResponseModel
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /order [get]
func (c *OrderControllerImpl) GetOrderBySiparisID(ctx *gin.Context) {

	siparisID := ctx.Query("siparisID")
	if siparisID == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("siparisID bos gecilemez"))
		return
	}

	orderList, err := c.service.GetOrderBySiparisID(siparisID)
	if err != nil {
		ctx.JSON(err.Code, response.ErrorResponse(err.Message))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("order list", orderList))
}

// OrderSummaryList
//
// @Summary      Get order summary list
// @Tags         order
// @Accept       json
// @Produce      json
// @Param 		 islemAdi query string true "islem adi"
// @Param 		 musteriKodu query string true "musteri kodu"
// @Success      200  {object} response.SuccessResponseModel
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /order/summary [get]
func (c *OrderControllerImpl) GetOrderSummaryList(ctx *gin.Context) {

	islemAdi := ctx.Query("islemAdi")
	musteriKodu := ctx.Query("musteriKodu")

	orderSummaryList, err := c.service.GetOrderSummaryList(islemAdi, musteriKodu)
	if err != nil {
		ctx.JSON(err.Code, response.ErrorResponse(err.Message))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("order summary list", orderSummaryList))
}

// OrderUretimBilgileriBySiparisID
//
// @Summary      Get order uretim bilgileri by siparis id
// @Tags         order
// @Accept       json
// @Produce      json
// @Param			 siparisID query string true "siparis id"
// @Param			 uretimYeri query string true "uretim yeri"
// @Success      200  {object} response.SuccessResponseModel
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /order/summary/uretim [get]
func (c *OrderControllerImpl) GetOrderUretimBilgileriBySiparisID(ctx *gin.Context) {

	siparisID := ctx.Query("siparisID")
	uretimYeri := ctx.Query("uretimYeri")
	if siparisID == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("siparisID bos gecilemez"))
		return
	}
	if uretimYeri == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("uretimYeri bos gecilemez"))
		return
	}

	orderUretimBilgileri, err := c.service.GetOrderUretimBilgileriBySiparisID(siparisID, uretimYeri)
	if err != nil {
		ctx.JSON(err.Code, response.ErrorResponse(err.Message))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("order uretim bilgileri", orderUretimBilgileri))
}

// Add New Uretim
//
// @Summary      Add new uretim
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        siparisNo formData string true "Siparis no"
// @Param        uretimDurum formData string true "Uretim durum"
// @Param        uretimYeri formData string true "Uretim yeri"
// @Param        miktari formData int true "Miktari"
// @Param        images formData file false "Resimler" multiple
// @Success      200  {object} response.SuccessResponseModel
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /order/uretim [post]
func (c *OrderControllerImpl) AddNewUretim(ctx *gin.Context) {

	ctx.Request.ParseMultipartForm(20 << 20)
	var request models.UretimAddRequest

	claims := ctx.MustGet("claims").(*jwt.Claims)

	request.Kullanici = claims.Username

	request.File = ctx.Request.MultipartForm.File["images"]
	request.SiparisNo = ctx.Request.FormValue("siparisNo")
	request.UretimDurum = ctx.Request.FormValue("uretimDurum")
	request.UretimYeri = ctx.Request.FormValue("uretimYeri")
	miktar, err := strconv.Atoi(ctx.Request.FormValue("miktari"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}
	request.Miktar = miktar

	if miktar <= 0 {

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("miktar gecersiz"))
		return
	}

	if request.SiparisNo == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("siparisNo gecersiz"))
		return
	}

	if request.UretimYeri == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("uretimYeri gecersiz"))
		return
	}

	if request.UretimDurum == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("uretimDurum gecersiz"))
		return
	}

	if err := c.service.AddNewUretim(request); err != nil {
		ctx.JSON(err.Code, response.ErrorResponse(err.Message))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("uretim eklendi", nil))
}

// Delete Uretim
//
// @Summary      Delete uretim
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        id query string true "id"
// @Success      200  {object} response.SuccessResponseModel
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /order/uretim [delete]
func (c *OrderControllerImpl) DeleteUretim(ctx *gin.Context) {

	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("id zorunlu"))
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err := c.service.DeleteUretim(idInt); err != nil {
		ctx.JSON(err.Code, response.ErrorResponse(err.Message))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("uretim silindi", nil))
}

// Get Customer Orders By Islem Adi
//
// @Summary      Get customer orders by islem adi
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        islemAdi query string true "islem adi"
// @Success      200  {object} response.SuccessResponseModel
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /order/customerOrders [get]
func (c *OrderControllerImpl) GetCustomerOrdersByIslemAdi(ctx *gin.Context) {

	islemAdi := ctx.Query("islemAdi")
	if islemAdi == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("islemAdi zorunlu"))
		return
	}

	customerOrders, err := c.service.GetCustomerOrdersByIslemAdi(islemAdi)
	if err != nil {
		ctx.JSON(err.Code, response.ErrorResponse(err.Message))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("customer orders", customerOrders))
}

// Add New Model Resim
//
// @Summary      Add new model resim
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        image formData file true "Image"
// @Param        kodu formData string true "Kodu"
// @Success      200  {object} response.SuccessResponseModel
// @Failure      400  {object} response.ErrorResponseModel
// @Security     BearerAuth
// @Router       /order/modelResim [post]
func (c *OrderControllerImpl) AddNewModelResim(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}
	kodu := ctx.Request.FormValue("kodu")

	if file == nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("file zorunlu"))
		return
	}

	if kodu == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("kodu zorunlu"))
		return
	}

	url, appErr := c.service.AddNewModelResim(file, kodu)
	if appErr != nil {
		ctx.JSON(appErr.Code, response.ErrorResponse(appErr.Message))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("model resim eklendi", url))
}
