package model

import (
	"github.com/guregu/null"
)

// Subscription represents user's notification settings.
type Subscription struct {
	ID          int      `json:"id" gorm:"primary_key"`
	UserID      null.Int `json:"user_id" gorm:"ForeignKey:ID"`
	ProjectUUID string   `json:"project_uuid" gorm:"ForeignKey:UUID"`

	Email     string `json:"email"`
	DeletedAt Time   `json:"deleted_at"`
}
