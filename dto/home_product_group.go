package dto

import (
	"encoding/json"
	"time"
)

type HomeGroupProduct struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	HomeID  string `json:"home_id"`
	GroupID string `json:"group_id"`
}

func (e *HomeGroupProduct) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type HomeGroupProductList []HomeGroupProduct

func (e *HomeGroupProductList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
