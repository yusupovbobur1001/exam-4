package handler

import (
	pb "api_service/genproto/booking"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swager
// @Summary Create Payment
// @Description Create Payment
// @Tags Payment
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param payment body booking.CreatePaymentRequest true "payment"
// @Success 200 {object} booking.CreatePaymentResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /payment/payments [post]
func (h *Handler) CreatePayment(c *gin.Context) {
	req := pb.CreatePaymentRequest{}
	
	err := c.BindJSON(&req)
	if err != nil {
		fmt.Println("111111111111111111111111111111111111111111111111111111111111111")
		h.Logger.Error(err.Error())
		c.AbortWithError(400, err)
		return
	}

	t, err := h.ClientBooking.GetBooking(c, &pb.GetBookingRequest{XId: req.BookingId})
	if t == nil || err != nil {
		fmt.Println("22222222222222222222222222222222", err)
		h.Logger.Error(err.Error())
		c.AbortWithError(400, err)
		return
	}

	res, err := h.ClientBooking.CreatePayment(c, &req)
	if err != nil {
		fmt.Println("33333333333333333333333333333333", err)
		h.Logger.Error(err.Error())
		c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// swager
// @Summary Get Payment
// @Description Get Payment
// @Tags Payment
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} booking.GetPaymentResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /payment/payments/{id} [get]
func (h *Handler) GetPayment(c *gin.Context) {
	id := c.Param("id")

	res, err := h.ClientBooking.GetPayment(c, &pb.GetPaymentRequest{XId: id})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(400, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// swager
// @Summary List Payments
// @Description List Payments
// @Tags Payment
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} booking.ListPaymentsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /payment/payments [get]
func (h *Handler) ListPayments(c *gin.Context) {
	l := c.Query("limit")
	limit, err := strconv.Atoi(l)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(400, err)
		return
	}

	o := c.Query("offset")
	offset, err := strconv.Atoi(o)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.ClientBooking.ListPayments(c, &pb.ListPaymentsRequest{Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(400, err)
		return
	}

	c.JSON(http.StatusOK, res)
}	