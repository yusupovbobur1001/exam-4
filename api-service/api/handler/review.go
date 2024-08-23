package handler

import (
	pb "api_service/genproto/booking"
	pbu "api_service/genproto/user"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Review
// @Description Create Review
// @Tags Review
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param review body booking.CreateReviewRequest true "review"
// @Success 200 {object} booking.CreateReviewResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /review/reviews [post]
func (h *Handler) CreateReview(c *gin.Context) {
	req := pb.CreateReviewRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.AbortWithError(400, err)
		return
	}

	res, err := h.ClientBooking.CreateReview(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Update Review
// @Description Update Review
// @Tags Review
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Param review body booking.UpdateReviewRequest false "review"
// @Success 200 {object} booking.UpdateReviewResponse 
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /review/reviews/{id} [put]
func (h *Handler) UpdateReview(c *gin.Context) {
	id := c.Param("id")
	
	req := pb.UpdateReviewRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(400, err)
		return
	}
	
	t1, err := h.ClientUser.GetByIdProfile(c, &pbu.GetProfileRequest{Id: req.UserId})
	if t1 == nil || err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, err)
		return
	}

	t2, err := h.ClientBooking.GetProviders(c, &pb.GetProvidersRequest{XId: req.ProcviderId})
	if t2 == nil || err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, err)
		return
	}

	req.XId = id

	res, err := h.ClientBooking.UpdateReview(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Delete Review
// @Description Delete Review
// @Tags Review
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} booking.DeleteReviewResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /review/reviews/{id} [delete]
func (h *Handler) DeleteReview(c *gin.Context) {
	id := c.Param("id")

	res, err := h.ClientBooking.DeleteReview(c, &pb.DeleteReviewRequest{XId: id})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, err)
		return
	}

	c.JSON(200, res)
}

// @Summary List Reviews
// @Description List Reviews
// @Tags Review
// @Accept json
// @Produce json
// @Security     ApiKeyAuth
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} booking.ListReviewsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /review/reviews [get]
func (h *Handler) ListReviews(c *gin.Context) {
	l := c.Query("limit")
	limit, err := strconv.Atoi(l)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(400, err)
		return
	}
	fmt.Println(limit+10)

	o := c.Query("offset")
	offset, err := strconv.Atoi(o)
	if err != nil {
		fmt.Println(err)
		h.Logger.Error(err.Error())
		c.JSON(400, err)
		return
	}
	fmt.Println(offset, "00")
	res, err := h.ClientBooking.ListReviews(c, &pb.ListReviewsRequest{Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, err)
		return
	}

	c.JSON(200, res)
}