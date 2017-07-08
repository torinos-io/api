package model

// Notification represents user's notification settings.
type Notification struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	ProjectUUID string `json:"project_uuid"`

	Email     string `json:"email"`
	DeletedAt Time   `json:"deleted_at"`
}
