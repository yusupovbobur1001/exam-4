package handler

import (
	pb "api_service/genproto/booking"
	pbu "api_service/genproto/user"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary Create Providers
// @Description Create Providers
// @Tags Providers
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param providers body booking.CreateProvidersRequest true "providers"
// @Success 200 {object} booking.CreateProvidersResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /provider/providers [post]
func (h *Handler) CreateProviders(c *gin.Context) {
	req := pb.CreateProvidersRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	t, err := h.ClientUser.GetByIdProfile(c, &pbu.GetProfileRequest{Id: req.UserId})
	if t == nil || err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invaled user id",
		})
		return
	}

	res, err := h.ClientBooking.CreateProviders(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Update Providers
// @Description Update Providers
// @Tags Providers
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Param providers body booking.UpdateProvidersRequest false "providers"
// @Success 200 {object} booking.UpdateProvidersResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /provider/providers/{id} [put]
func (h *Handler) UpdateProviders(c *gin.Context) {
	id := c.Param("id")

	req := pb.UpdateProvidersRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(req.Availabilitys, req.UserId, "-----------------------")
	t, err := h.ClientBooking.UpdateProviders(c, &req)
	if t == nil || err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invaled user id",
		})
		return
	}

	req.UserId = id

	res, err := h.ClientBooking.UpdateProviders(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Delete Providers
// @Description Delete Providers
// @Tags Providers
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} booking.DeleteProvidersResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /provider/providers/{id} [delete]
func (h *Handler) DeleteProviders(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	res, err := h.ClientBooking.DeleteProviders(c, &pb.DeleteProvidersRequest{XId: id})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Get Providers
// @Description Get Providers
// @Tags Providers
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} booking.GetProvidersResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /provider/providers/{id} [get]
func (h *Handler) GetProviders(c *gin.Context) {
	id := c.Param("id")

	res, err := h.ClientBooking.GetProviders(c, &pb.GetProvidersRequest{XId: id})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(res.Availabilitys, "----------")
	c.JSON(http.StatusOK, res)
}

// @Summary List Providers
// @Description List Providers
// @Tags Providers
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} booking.ListProvidersResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /provider/providers [get]
func (h *Handler) ListProviders(c *gin.Context) {
	l := c.Query("limit")
	limit, err := strconv.Atoi(l)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	o := c.Query("offset")
	offset, err := strconv.Atoi(o)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}
	
	res, err := h.ClientBooking.ListProviders(c, &pb.ListProvidersRequest{Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Search Providers
// @Description Search Providers
// @Tags Providers
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param user_id query string false "user_id"
// @Param company_name query string false "company_name"
// @Param location query string false "location"
// @Param availability_as query []string false "availability_as"
// @Success 200 {object} booking.SearchProvidersResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /provider/providers/search [get]
func (h *Handler) SearchProviders(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		userID = ""
	} else {
		t, err := h.ClientUser.GetByIdProfile(c, &pbu.GetProfileRequest{Id: userID})
		if t == nil || err != nil {
			fmt.Println(err)
			h.Logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid user id",
			})
			return
		}
	}

	companyName := c.Query("company_name")
	if companyName == "" {
		companyName = ""
	}

	location := c.Query("location")
	if location == "" {
		location = ""
	}

	availabilitys := []*pb.AvailabilityR{}
	availabilityAs := c.QueryArray("availability_as")
	if len(availabilityAs) > 0 {
		for _, availability := range availabilityAs {
			times := strings.Split(availability, "-")
			if len(times) == 2 {
				availabilitys = append(availabilitys, &pb.AvailabilityR{
					StartTime: times[0],
					EndTime:   times[1],
				})
			}
		}
	}

	res, err := h.ClientBooking.SearchProviders(c, &pb.SearchProvidersRequest{
		UserId:        userID,
		CompanyName:   companyName,
		Location:      &pb.Location{City: location},
		Availabilitys: availabilitys,
	})

	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to search providers",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
