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
	UpdatedAt *time.Time `json:"updated_at" gorm:"<-:update"`
	Name      string     `json:"name" gorm:"column:name;type:string;size:256"`
}

func (e *AppConfig) GenerateAndSetID() {
	e.ID = e.generateID()
}

func (e AppConfig) generateID() string {
	return uuid.New().String()
}
