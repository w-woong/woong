package entity

import (
	"time"

	"github.com/google/uuid"
)

var (
	NilAppConfig = AppConfig{}
)

type AppConfig struct {
	ID        string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"<-"`
	Name      string     `json:"name" gorm:"column:name;type:string;size:256"`
}

func (e AppConfig) IsNil() bool {
	return e.ID == ""
}

func (e *AppConfig) CreateSetID() {
	e.ID = e.CreateID()
}

func (e AppConfig) CreateID() string {
	return uuid.New().String()
}
