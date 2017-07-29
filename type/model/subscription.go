package model

import (
	"github.com/guregu/null"
)

// Subscription represents user's notification settings.
type Subscription struct {
	ID          int      `json:"id"`
	UserID      null.Int `json:"user_id"`
	ProjectUUID string   `json:"project_uuid" gorm:"ForeignKey:UUID"`

	Email     string `json:"email"`
	DeletedAt Time   `json:"deleted_at"`
}
