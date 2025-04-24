package handler

import (
	"net/http"
	"strconv"

	"github.com/abhiraj-ku/health_app/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PatientService interface {
	Create(*model.Patient) error
	GetAll() ([]model.Patient, error)
	Update(*model.Patient) error
	Delete(int64) error
	GetById(int64) (*model.Patient, error)
}

type PatientHandler struct {
	Service PatientService
}

func NewPatientHandler(s PatientService) *PatientHandler {
	return &PatientHandler{
		Service: s,
	}

}

func (h *PatientHandler) RegisterRoutes(r *gin.Engine, authMiddleware gin.HandlerFunc, requireRole func(...string) gin.HandlerFunc) {
	patients := r.Group("/patients", authMiddleware)
	{
		patients.POST("", requireRole("receptionist"), h.Create)
	}
}

func (h *PatientHandler) Create(c *gin.Context) {
	var p model.Patient

	if err := c.ShouldBindJSON(&p); err != nil {
		if validationError, ok := err.(validator.ValidationErrors); ok {

			errors := make(map[string]string)
			for _, fieldErr := range validationError {

				errors[fieldErr.Field()] = fieldErr.Error()
			}
			c.JSON(http.StatusBadRequest, gin.H{"validation_error": errors})
		} else {

			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		}
		return
	}

	userID := c.GetInt64("userID")
	p.UpdatedBy = userID

	if err := h.Service.Create(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create the user record"})
		return
	}

}
func (h *PatientHandler) GetAll(c *gin.Context) {
	patients, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get patients"})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func (h *PatientHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var p model.Patient
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	p.ID = id
	p.UpdatedBy = c.GetInt64("userID")
	if err := h.Service.Update(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *PatientHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.Status(http.StatusNoContent)
}
