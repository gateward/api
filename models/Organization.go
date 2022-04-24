package models

import "time"

type Organization struct {
	OrgID     string    `json:"id"`
	Name      string    `json:"name"`
	AdminsIDs string    `json:"admins_ids"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrgInvit struct {
	InvitID    string    `json:"id"`
	AdminID    string    `json:"admin_id"`
	UserID     string    `json:"user_id"`
	Pending    bool      `json:"pending"`
	AcceptedAt string    `json:"accepted_at,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
