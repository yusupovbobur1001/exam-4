package handler

import (
	pb "api_service/genproto/booking"
	pbu "api_service/genproto/user"
	"api_service/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Booking
// @Description Create Booking
// @Tags Booking
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param booking body booking.CreateBookingRequest true "booking"
// @Success 200 {object} string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /booking/bookings [post]
func (h *Handler) CreateBooking(c *gin.Context) {
	req := pb.CreateBookingRequest{}
	
	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	t, err := h.ClientUser.GetByIdProfile(c, &pbu.GetProfileRequest{Id: req.UserId})
	if t == nil || err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.Kafka.Producermessage("createBooking", body)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "success")
}

// @Summary Get Booking
// @Description Get Booking
// @Tags Booking
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} booking.GetBookingResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /booking/bookings/{id} [get]
func (h *Handler) GetBooking(c *gin.Context) {
	req := pb.GetBookingRequest{}

	id := c.Param("id")
	req.XId = id

	resp, err := h.ClientBooking.GetBooking(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Update Booking
// @Description Update Booking
// @Tags Booking
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Param booking body model.UpdateBookingRequest false "booking"
// @Success 200 {object} string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /booking/bookings/{id} [put]
func (h *Handler) UpdateBooking(c *gin.Context) {
	id := c.Param("id")

	req := model.UpdateBookingRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	a := pb.UpdateBookingRequest{
		UserId:     req.UserId,
		ProviderId: req.ProviderId,
		ServiceId:  req.ServiceId,
		Status:     req.Status,
		TatolPrice: req.TatolPrice,
		XId:        id,
	}

	body, err := json.Marshal(&a)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(req, body)
	err = h.Kafka.Producermessage("updateBooking", body)
	if err != nil {
		fmt.Println(err)
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}

// @Summary Delete Booking
// @Description Delete Booking
// @Tags Booking
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /booking/bookings/{id} [delete]
func (h *Handler) DeleteBooking(c *gin.Context) {
	id := c.Param("id")

	req := pb.CancelBookingRequest{XId: id}

	body, err := json.Marshal(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.Kafka.Producermessage("createBooking", body)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "success")
}

// @Summary List Bookings
// @Description List Bookings
// @Tags Booking
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} booking.ListBookingsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /booking/bookings [get]
func (h *Handler) GetBookingList(c *gin.Context) {
	fmt.Println("1111111111")
	l := c.Query("limit")
	limit, err := strconv.Atoi(l)
	if err != nil {
		h.Logger.Error(err.Error())
		log.Print(err)
		c.JSON(400, err.Error())
		return
	}

	o := c.Query("offset")
	offset, err := strconv.Atoi(o)
	if err != nil {
		h.Logger.Error(err.Error())
		log.Print(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println("222222")
	resp, err := h.ClientBooking.ListBookings(c, &pb.ListBookingsRequest{Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		h.Logger.Error(err.Error())
		log.Print(err)
		c.JSON(400, err)
		return
	}

	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// swager
// @Summary Get Most Frequent Service ID
// @Description Get Most Frequent Service ID
// @Tags Booking
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Success 200 {object} booking.GetMostRequest
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /booking/most-frequent-service-id [get]
func (h *Handler) GetMostFrequentServiceID(c *gin.Context) {
	req := pb.Void{}

	res, err := h.ClientBooking.GetMostFrequentServiceID(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		log.Print(err)
		c.JSON(400, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
