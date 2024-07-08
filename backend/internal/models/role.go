package models

import "github.com/lib/pq"

type Role struct {
	ID   string   `json:"id" db:"id"`
	Name string   `json:"name" db:"name"`
	Menu []string `json:"menu"`
}
type RoleWithMenuDTO struct {
	ID      string         `json:"id" db:"id"`
	Name    string         `json:"name" db:"name"`
	Extends pq.StringArray `db:"extends"`
	Menu    pq.StringArray `db:"menu"`
	// Menu    string         `db:"menu"`
}

type RoleFull struct {
	ID          string   `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Level       int      `json:"level" db:"level"`
	Extends     []string `json:"extends" db:"extends"`
	Description string   `json:"description" db:"description"`
}
type RoleFullDTO struct {
	ID          string         `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Level       int            `json:"level" db:"level"`
	Extends     pq.StringArray `json:"extends" db:"extends"`
	Description string         `json:"description" db:"description"`
}

type RoleWithApi struct{}

type GetRolesDTO struct{}

type RoleDTO struct {
	ID          string   `json:"id" db:"id"`
	Name        string   `json:"name" db:"name" binding:"required"`
	Level       int      `json:"level" db:"level" binding:"required"`
	Extends     []string `json:"extends" db:"extends"`
	Description string   `json:"description" db:"description"`
}
