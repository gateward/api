package models

import "time"

type User struct {
	UserID    string    `json:"id"`
	OrgsIDs   string    `json:"orgs_ids,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
