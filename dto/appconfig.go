package dto

import "time"

var (
	NilAppConfig = AppConfig{}
)

type AppConfig struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Name      string     `json:"name"`
}

func (d AppConfig) IsNil() bool {
	return d.ID == ""
}
