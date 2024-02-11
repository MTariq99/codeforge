package models

type User struct {
	User_Name     *string `json:"user_name" db:"user_name"`
	User_bio      *string `json:"user_bio,omitempty" db:"user_bio,omitempty"`
	User_profile  *string `json:"user_profile,omitempty" db:"user_profile,omitempty"`
	User_email    *string `json:"user_email,omitempty" db:"user_email,omitempty"`
	User_password string  `json:"user_password,omitempty" db:"user_password,omitempty"`
	Gender        *string `json:"gender" db:"gender"`
	Created_at    *int64  `json:"created_at" db:"created_at"`
	Updated_at    *int64  `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	Deleted_at    *int64  `json:"deleted_at,omitempty" db:"deleted_at,omitempty"`
}
