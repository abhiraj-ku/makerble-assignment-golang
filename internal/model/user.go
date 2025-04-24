package model

import "time"

type Role string

// enums for Role
const (
	RoleReceptionist Role = "receptionist"
	RoleDoctor       Role = "doctor"
)

type User struct {
	ID        int64     `json:"id" db:"id" `
	Name      string    `json:"name" db:"name" binding:"required,min=2,max=100"`
	Password  string    `json:"-" db:"password" `
	Role      Role      `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
