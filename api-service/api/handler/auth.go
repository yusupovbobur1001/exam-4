package handler

import (
	pb "api_service/genproto/user"
	"api_service/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Register a new user
// @Description Registers a new user with the provided details.
// @Tags Users
// @Accept  json
// @Produce  json
// @Security     ApiKeyAuth
// @Param request body user.RegisterRequest true "User Registration Request"
// @Success 200 {object} user.RegisterResponse "User registration successful"
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	req := pb.RegisterRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ClientUser.Register(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

// @Summary Update user profile
// @Description Updates the user profile with the provided details.
// @Tags Users
// @Accept  json
// @Produce  json
// @Security     ApiKeyAuth
// @Param id path string true "User ID"
// @Param request body model.UpdateProfileRequest true "User Profile Update Request"
// @Success 200 {object} user.UpdateProfileResponse "User profile updated successfully"
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/users/{id} [put]
func (h *Handler) UpdateUserProfile(c *gin.Context) {
	req := model.UpdateProfileRequest{}

	id := c.Param("id")

	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ClientUser.UpdateUserProfile(c, &pb.UpdateProfileRequest{
		NewFirstName:   req.NewFirstName,
		NewPhoneNumber: req.NewPhoneNumber,
		NewRole:        req.NewRole,
		Id:             id,
	})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

// @Summary Delete user profile
// @Description Deletes the user profile specified by the ID.
// @Tags Users
// @Accept  json
// @Produce  json
// @Security     ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} user.DeleteProfileResponse "User profile deleted successfully"
// @Failure 500 {object} error
// @Router /auth/users/{id} [delete]
func (h *Handler) DeleteUserProfile(c *gin.Context) {
	req := pb.DeleteProfileRequest{}

	id := c.Param("id")
	req.Id = id

	res, err := h.ClientUser.DeleteUserProfile(c, &req)
	if err != nil {
		fmt.Println(err)
		h.Logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

// @Summary Get user profile by ID
// @Description Retrieves the user profile specified by the ID.
// @Tags Users
// @Accept  json
// @Produce  json
// @Security     ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} user.GetProfileResponse "User profile retrieved successfully"
// @Failure 500 {object} error
// @Router /auth/users/{id} [get]
func (h *Handler) GetByIdProfile(c *gin.Context) {
	req := pb.GetProfileRequest{}

	id := c.Param("id")
	req.Id = id

	res, err := h.ClientUser.GetByIdProfile(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

// swager
// @Summary Get all user profiles
// @Description Retrieves all user profiles.
// @Tags Users
// @Accept  json
// @Produce  json
// @Security     ApiKeyAuth
// @Param limit query int false "Limit of users to retrieve"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} user.GetProfilesResponse "User profiles retrieved successfully"
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/users [get]
func (h *Handler) GetAllProfile(c *gin.Context) {
	req := pb.GetProfilesRequest{}

	l := c.Query("limit")
	limit, err := strconv.Atoi(l)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(400, err.Error())
		return
	}

	o := c.Query("offset")
	offset, err := strconv.Atoi(o)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(400, err)
		return
	}

	req.Limit = int32(limit)
	req.Offset = int32(offset)

	res, err := h.ClientUser.GetAllProfile(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, res)
}
