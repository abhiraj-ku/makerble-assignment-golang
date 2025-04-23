package model

import "time"

type Role string

// enums for Role
const (
	RoleReceptionist Role = "receptionist"
	RoleDoctor       Role = "doctor"
)

type User struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"password" db:"password"`
	Role      Role      `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
