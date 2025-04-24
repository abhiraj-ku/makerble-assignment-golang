package repository

import "github.com/abhiraj-ku/health_app/internal/model"

type PatientRepository interface {
	Create(patient *model.Patient) error
	GetAll() ([]model.Patient, error)
	Update(patient *model.Patient) error
	Delete(id int64) error
	GetById(id int64) (*model.Patient, error)
}
