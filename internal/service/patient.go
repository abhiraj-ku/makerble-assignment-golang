package service

import (
	"github.com/abhiraj-ku/health_app/internal/model"
	"github.com/abhiraj-ku/health_app/internal/repository"
)

type PatientService struct {
	Repo repository.PatientRepository
}

func NewPatientService(r repository.PatientRepository) *PatientService {
	return &PatientService{
		Repo: r,
	}
}

func (s *PatientService) Create(p *model.Patient) error {
	return s.Repo.Create(p)
}

func (s *PatientService) GetAll() ([]model.Patient, error) {
	return s.Repo.GetAll()
}

func (s *PatientService) Update(p *model.Patient) error {
	return s.Repo.Update(p)
}

func (s *PatientService) Delete(id int64) error {
	return s.Repo.Delete(id)
}

func (s *PatientService) GetById(id int64) (*model.Patient, error) {
	return s.Repo.GetById(id)
}
