package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// "strconv"
	// "strings"

	model "project-riskprofile/Model"
	"project-riskprofile/entity"
	"project-riskprofile/service"

	"github.com/gin-gonic/gin"
)

// IUserHandler mendefinisikan interface untuk handler user
type IsubmissionHandler interface {
	CreateSubmi(c *gin.Context)
	GetSubmi(c *gin.Context)
	// UpdateSubmi(c *gin.Context)
	// UpdateUser(c *gin.Context)
	GetUSubmiByUserID(c *gin.Context)
	DeleteSubmi(c *gin.Context)
	GetAllSubmi(c *gin.Context)
}

type submissionHandler struct {
	submissionService service.IsubmissionService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewsubmissionHandler(submissionService service.IsubmissionService) IsubmissionHandler {
	return &submissionHandler{
		submissionService: submissionService,
	}
}

type MessageCreate struct {
	Message string `json:"message"`
}

// CreateUser menghandle permintaan untuk membuat user baru
func (h *submissionHandler) CreateSubmi(c *gin.Context) {
	var postsub entity.Post_Submission
	if err := c.ShouldBindJSON(&postsub); err != nil {
		errMsg := err.Error()
		//errMsg = convertUserMandatoryFieldErrorString(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var submission entity.Submission
	answersJSON, err := json.Marshal(postsub.Answers)

	score := model.FindWeight2(postsub.Answers)
	RiskCategory := model.GetRiskProfile(score)
	submission.RiskCategory = string(RiskCategory.Category)

	submission.UserID = postsub.UserID
	submission.Answers = answersJSON

	createdUser, err := h.submissionService.CreateSubmi(c.Request.Context(), &submission)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, MessageCreate{Message: fmt.Sprintf("success create submission with ID %d", createdUser.ID)})
	// c.JSON(http.StatusCreated, interface{})
}

// GetUser menghandle permintaan untuk mendapatkan user berdasarkan ID
func (h *submissionHandler) GetSubmi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := h.submissionService.GetUSubmiByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subs not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *submissionHandler) GetUSubmiByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	print(id)
	user, err := h.submissionService.GetUSubmiByUserID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subs not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// // GetUser menghandle permintaan untuk mendapatkan user berdasarkan ID
// func (h *UserHandler) GetUserByEmail(c *gin.Context) {
// 	email := c.Param("email")
// 	user, err := h.userService.GetUserByEmail(c.Request.Context(), email)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// UpdateUser menghandle permintaan untuk mengupdate informasi user
// func (h *submissionHandler) UpdateSubmi(c *gin.Context) {
// 	var postsub entity.Post_Submission

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 		return
// 	}

// 	var submission entity.Submission
// 	if err := c.ShouldBindJSON(&submission); err != nil {
// 		errMsg := err.Error()
// 		errMsg = convertUserMandatoryFieldErrorString(errMsg)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
// 		return
// 	}

// 	answersJSON, err := json.Marshal(postsub.Answers)

// 	score := model.FindWeight2(postsub.Answers)
// 	RiskCategory := model.GetRiskProfile(score)
// 	submission.RiskCategory = string(RiskCategory.Category)

// 	submission.UserID = postsub.UserID
// 	submission.Answers = answersJSON

// 	updatedUser, err := h.submissionService.UpdateSubmi(c.Request.Context(), id, submission)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, updatedUser)
// }

// DeleteUser menghandle permintaan untuk menghapus user
func (h *submissionHandler) DeleteSubmi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.submissionService.DeleteSubmi(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Submi"})
}

// GetAllUsers menghandle permintaan untuk mendapatkan semua user
func (h *submissionHandler) GetAllSubmi(c *gin.Context) {
	users, err := h.submissionService.GetAllSubmi(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// func convertUserMandatoryFieldErrorString(oldErrorMsg string) string {
// 	switch {
// 	case strings.Contains(oldErrorMsg, "'Name' failed on the 'required' tag"):
// 		return "name is mandatory"
// 	case strings.Contains(oldErrorMsg, "'Email' failed on the 'required' tag"):
// 		return "email is mandatory"
// 	}
// 	return oldErrorMsg
// }
