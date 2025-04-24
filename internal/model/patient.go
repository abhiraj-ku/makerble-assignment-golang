package model

import "time"

type Patient struct {
	ID              int64     `json:"id" db:"id"`
	Name            string    `json:"name" db:"name" binding:"required,min=2,max=100"`
	Age             int       `json:"age ,required" db:"age" binding:"required,min=0,max=110" `
	Gender          string    `json:"gender" db:"gender" binding:"required,oneof=male female other"`
	Contact         string    `json:"contact" db:"contact" binding:"required"`
	Address         string    `json:"address" db:"address" binding:"required"`
	Disease         string    `json:"disease" db:"disease" binding:"required"`
	HandledByDoctor int64     `json:"handled_by_doctor" db:"handled_by_doctor" binding:"required"`
	UpdatedBy       int64     `json:"updated_by" db:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
