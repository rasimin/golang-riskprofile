package handler

import (
	"fmt"
	"net/http"
	model "project-riskprofile/Model"
	"project-riskprofile/entity"
	"project-riskprofile/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// IUserHandler mendefinisikan interface untuk handler user
type IUserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	userService service.IUserService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser menghandle permintaan untuk membuat user baru
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := err.Error()
		errMsg = convertUserMandatoryFieldErrorString(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	createdUser, err := h.userService.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//c.JSON(http.StatusCreated, createdUser)
	c.JSON(http.StatusCreated, MessageCreate{Message: fmt.Sprintf("success create user with ID %d", createdUser.ID)})
}

// GetUser menghandle permintaan untuk mendapatkan user berdasarkan ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	user.Definition = model.GetDefinitionByCategory(model.ProfileRiskCategory(user.RiskCategory))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUser menghandle permintaan untuk mendapatkan user berdasarkan ID
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.userService.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser menghandle permintaan untuk mengupdate informasi user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := err.Error()
		errMsg = convertUserMandatoryFieldErrorString(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	updatedUser, err := h.userService.UpdateUser(c.Request.Context(), id, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser menghandle permintaan untuk menghapus user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// GetAllUsers menghandle permintaan untuk mendapatkan semua user
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func convertUserMandatoryFieldErrorString(oldErrorMsg string) string {
	switch {
	case strings.Contains(oldErrorMsg, "'Name' failed on the 'required' tag"):
		return "name is mandatory"
	case strings.Contains(oldErrorMsg, "'Email' failed on the 'required' tag"):
		return "email is mandatory"
	}
	return oldErrorMsg
}
