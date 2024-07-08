package models

type User struct {
	ID           string   `json:"id" db:"id"`
	Name         string   `json:"name" db:"name"`
	Role         string   `json:"role"`
	Menu         []string `json:"menu"`
	AccessToken  string   `json:"token"`
	RefreshToken string   `json:"-"`
}
