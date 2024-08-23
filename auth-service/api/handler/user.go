package handler

import (
	"auth_service/api/token"
	pb "auth_service/genproto/user"
	"auth_service/model"
	"auth_service/service"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticaionHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout(c *gin.Context)
}

type AuthenticaionHandlerImpl struct {
	UserManage service.AuthServiceI
	Logger     *slog.Logger
}

func NewAuthenticaionHandlerImpl(userManage service.AuthServiceI, logger *slog.Logger) *AuthenticaionHandlerImpl {
	return &AuthenticaionHandlerImpl{
		UserManage: userManage,
		Logger:     logger,
	}
}

// @Summary      Login a new user
// @Description  Log in a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     model.LoginRequest  true  "User registration request"
// @Success      200   {object}  model.Tokens
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /auth/login [post]
func (h *AuthenticaionHandlerImpl) Login(c *gin.Context) {
	req := model.LoginRequest{}
	fmt.Println(4)
	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.UserManage.Login(c.Request.Context(), &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := token.GenerateJWT(resp)

	c.JSON(http.StatusOK, res)
}


// @Summary      Register a new user
// @Description  Register a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     user.RegisterRequest  true  "User registration request"
// @Success      200   {object}  user.RegisterResponse
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /auth/register [post]
func (h *AuthenticaionHandlerImpl) Register(c *gin.Context) {
	req := pb.RegisterRequest{}
	fmt.Println(1)
	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.UserManage.Register(c.Request.Context(), &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
	
}


// @Summary      Logout a user
// @Description  Log out a user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     model.LogoutRequest  true  "User logout request"
// @Success      200   {object}  model.LogoutResponse
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /auth/logout [post]
func (h *AuthenticaionHandlerImpl) Logout(c *gin.Context) {
	req := model.LogoutRequest{}

	resp, err := h.UserManage.Logout(c.Request.Context(), &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
