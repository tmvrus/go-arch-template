package model

type User struct {
	Name         string `json:"name"   db:"name"`
	Password     string `json:"pass"   db:"pass"`
	EnableNotify bool   `json:"enable" db:"enable"`
}
