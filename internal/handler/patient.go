package handler

import (
	"log"
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
		// Receptionist routes
		patients.POST("", requireRole("receptionist"), h.Create)
		patients.DELETE("/:id", requireRole("receptionist"), h.Delete)

		// Both Parties can view patients details
		patients.GET("", requireRole("receptionist", "doctor"), h.GetAll)
		// Shared update route (both receptionist and doctor can update)
		patients.PUT("/:id", requireRole("receptionist", "doctor"), h.Update)
	}
}

func (h *PatientHandler) Create(c *gin.Context) {
	var patientRecord model.Patient

	if err := c.ShouldBindJSON(&patientRecord); err != nil {
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
	patientRecord.UpdatedBy = userID

	if err := h.Service.Create(&patientRecord); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create the user record"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "user registered successfully",
		"data":    patientRecord,
	})

}
func (h *PatientHandler) GetAll(c *gin.Context) {
	patients, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get patients"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Fetched all patient record successfully",
		"data":    patients,
	})

}

func (h *PatientHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Println(id)
	var patientRecord model.Patient
	if err := c.ShouldBindJSON(&patientRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	patientRecord.ID = id
	patientRecord.UpdatedBy = c.GetInt64("userID")
	log.Println(c.Get("userID"))
	log.Println(h.Service.Update(&patientRecord))
	if err := h.Service.Update(&patientRecord); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"message":    "patient record updated successfully",
		"updated_by": patientRecord.UpdatedBy,
		"data":       patientRecord,
	})
}

func (h *PatientHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Println(id)
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.Status(http.StatusNoContent)
	c.JSON(http.StatusNoContent, gin.H{
		"success": true,
		"message": "patient record deleted successfully",
		"data":    "",
	})
}
