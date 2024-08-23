package handler

import (
	pb "api_service/genproto/booking"
	pbu "api_service/genproto/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Service
// @Description Create Service
// @Tags Service
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param service body booking.CreateServiceRequest true "service"
// @Success 200 {object} booking.CreateServiceResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /service/services [post]
func (h *Handler) CreateService(c *gin.Context) {
	req := pb.CreateServiceRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ClientBooking.CreateService(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Update Service
// @Description Update Service
// @Tags Service
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Param service body booking.UpdateServiceRequest false "service"
// @Success 200 {object} booking.UpdateServiceResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /service/services/{id} [put]
func (h *Handler) UpdateService(c *gin.Context) {
	id := c.Param("id")

	req := pb.UpdateServiceRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	t, err := h.ClientUser.GetByIdProfile(c, &pbu.GetProfileRequest{Id: req.UserId})
	if t == nil || err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.XId = id

	res, err := h.ClientBooking.UpdateService(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Delete Service
// @Description Delete Service
// @Tags Service
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} booking.DeleteServiceResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /service/services/{id} [delete]
func (h *Handler) DeleteService(c *gin.Context) {
	id := c.Param("id")

	res, err := h.ClientBooking.DeleteService(c, &pb.DeleteServiceRequest{XId: id})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary List Services
// @Description List Services
// @Tags Service
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} booking.ListServicesResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /service/services [get]
func (h *Handler) ListServices(c *gin.Context) {
	l := c.Query("limit")
	limit, err := strconv.Atoi(l)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	o := c.Query("offset")
	offset, err := strconv.Atoi(o)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.ClientBooking.ListServices(c, &pb.ListServicesRequest{Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Search Services
// @Description Search Services
// @Tags Service
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param user_id query string false "user_id"
// @Param price query int false "price"
// @Param duration query int false "duration"
// @Success 200 {object} booking.SearchServicesResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /service/services/search [get]
func (h *Handler) SearchServices(c *gin.Context) {
	req := pb.SearchServicesRequest{}
	var userID string
	if len(req.UserId) > 0 {
		userID = c.Query("user_id")
		t, err := h.ClientUser.GetByIdProfile(c, &pbu.GetProfileRequest{Id: userID})
		if t == nil || err != nil {
			h.Logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid user id",
			})
			return
		}
	} else {
		userID = ""
	}

	p := c.Query("price")
	if p == "" {
		p = "0"
	}
	price, err := strconv.Atoi(p)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	d := c.Query("distance")
	if len(d) == 0 {
		d = "0"
	}
	duration, err := strconv.Atoi(d)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.UserId = userID
	req.Price = float32(price)
	req.Duration = int32(duration)
	fmt.Println(req.Duration, req.Price)
	res, err := h.ClientBooking.SearchServices(c, &req)
	if err != nil {
		fmt.Println(err)
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
