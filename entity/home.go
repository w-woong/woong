package entity

import (
	"encoding/json"
	"time"
)

var (
	NilHome        = Home{}
	NilShortNotice = ShortNotice{}
)

type Home struct {
	ID        string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"<-:update"`
	Name      string     `json:"name" gorm:"column:name;type:string;size:256"`

	AppConfigID string     `json:"app_config_id"`
	AppConfig   *AppConfig `json:"app_config" gorm:"foreignKey:AppConfigID;references:ID"`

	ShortNoticeList ShortNoticeList `json:"short_notice_list" gorm:"foreignKey:HomeID;references:ID"`
}

func (e *Home) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type ShortNotice struct {
	ID          string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	HomeID      string     `json:"home_id" gorm:"column:home_id;type:string;size:64"`
	CreatedAt   *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"<-:update"`
	ImgUrl      string     `json:"img_url" gorm:"type:string;size:2048"`
	Name        string     `json:"name" gorm:"column:name;type:string;size:256"`
	Description string     `json:"description" gorm:"column:description;type:string"` // without size = text
}

func (e *ShortNotice) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type ShortNoticeList []ShortNotice

func (e *ShortNoticeList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
