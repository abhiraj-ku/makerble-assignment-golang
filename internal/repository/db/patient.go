package postresdb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/abhiraj-ku/health_app/internal/model"
)

type Patient struct {
	db *sql.DB
}

func NewPatientRepo(db *sql.DB) *Patient {
	return &Patient{
		db: db,
	}
}

func (r *Patient) Create(p *model.Patient) error {
	const query = `
		INSERT INTO patients (
			name, age, gender, contact, address, disease, 
			handled_by_doctor, updated_by, updated_at, created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		) RETURNING id;
	`

	err := r.db.QueryRow(query,
		p.Name, p.Age, p.Gender, p.Contact, p.Address,
		p.Disease, p.HandledByDoctor, p.UpdatedBy, p.UpdatedAt, p.CreatedAt,
	).Scan(&p.ID)

	if err != nil {
		return fmt.Errorf("failed to insert patient: %w", err)
	}

	return nil
}

func (r *Patient) GetAll() ([]model.Patient, error) {
	const query = `SELECT id, name, age, gender,contact, address,disease,handled_by_doctor, updated_by, updated_at, created_at
	               FROM patients ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed %w", err)
	}

	defer rows.Close()

	var patients []model.Patient
	for rows.Next() {
		var p model.Patient
		err := rows.Scan(&p.Name, &p.Age, &p.Gender, &p.Contact, &p.Address,
			&p.Disease, &p.HandledByDoctor, &p.UpdatedBy, &p.UpdatedAt, &p.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed at getAll %w", err)
		}
		patients = append(patients, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}
	return patients, nil

}

func (r *Patient) Update(p *model.Patient) error {
	query := ` update patients set 
			name = $1, 
			age = $2, 
			gender = $3, 
			contact = $4, 
			address = $5, 
			disease = $6, 
			handled_by_doctor = $7, 
			updated_by = $8, 
			updated_at = $9 
	where id = $10;
	`
	_, err := r.db.Exec(query, p.Name, p.Age, p.Gender, p.Contact, p.Address,
		p.Disease, p.HandledByDoctor, p.UpdatedBy, time.Now(), p.ID)
	if err != nil {
		return fmt.Errorf("failed to update patient: %w", err)
	}

	return nil
}

func (r *Patient) Delete(id int64) error {
	query := `delete from patients where id= $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete patient with id %d: %w", id, err)
	}

	return nil
}

func (r *Patient) GetById(id int64) (*model.Patient, error) {
	query := `SELECT id, name, age, gender, contact, address, disease, 
		       handled_by_doctor, updated_by, updated_at, created_at
		FROM patients WHERE id = $1; `

	var p model.Patient

	err := r.db.QueryRow(query).Scan(
		&p.ID, &p.Name, &p.Age, &p.Gender, &p.Contact, &p.Address,
		&p.Disease, &p.HandledByDoctor, &p.UpdatedBy, &p.UpdatedAt, &p.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch patient by id: %w", err)
	}
	return &p, nil
}
