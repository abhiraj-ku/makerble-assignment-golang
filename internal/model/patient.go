package model

import "time"

type Patient struct {
	ID              int64     `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Age             int       `json:"age ,required" db:"age" `
	Gender          string    `json:"gender" db:"gender"`
	Contact         string    `json:"contact" db:"contact"`
	Address         string    `json:"address" db:"address"`
	Disease         string    `json:"disease" db:"disease"`
	HandledByDoctor int64     `json:"handled_by_doctor" db:"handled_by_doctor"`
	UpdatedBy       int64     `json:"updated_by" db:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
