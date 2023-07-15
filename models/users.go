package models

type (
	UserResponse struct {
		UserID     int64  `json:"user_id" db:"user_id"`
		Email      string `json:"email" db:"email"`
		Name       string `json:"name" db:"name"`
		DataSource string `json:"data_source"`
	}
)
