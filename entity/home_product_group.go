package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type HomeGroupProduct struct {
	ID        string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"<-"`

	HomeID  string `json:"home_id" gorm:"uniqueIndex:idx_home_group_product_1;not null;type:string;size:64"`
	GroupID string `json:"group_id" gorm:"uniqueIndex:idx_home_group_product_1;not null;type:string;size:64"`
}

func (e *HomeGroupProduct) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e HomeGroupProduct) CreateID() string {
	return uuid.New().String()
}

func (e *HomeGroupProduct) CreateSetID() {
	e.ID = e.CreateID()
}

type HomeGroupProductList []HomeGroupProduct

func (e *HomeGroupProductList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
