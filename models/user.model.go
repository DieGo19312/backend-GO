package models

import "time"

type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	PatSummary string    `json:"pat_summary"`
	MatSummary string    `json:"mat_summary"`
	Username   string    `json:"username"`
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Phone      int64     `json:"phone"`
	RoleID     *int       `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Active     bool      `json:"active"`
}

