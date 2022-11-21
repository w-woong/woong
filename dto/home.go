package dto

import (
	"encoding/json"
	"time"
)

var (
	NilHome        = Home{}
	NilShortNotice = ShortNotice{}
)

type Home struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Name      string     `json:"name"`

	AppConfigID string     `json:"app_config_id"`
	AppConfig   *AppConfig `json:"app_config,omitempty"`

	ShortNoticeList ShortNoticeList `json:"short_notice_list"`
}

func (e *Home) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type ShortNotice struct {
	ID          string     `json:"id"`
	HomeID      string     `json:"home_id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	ImgUrl      string     `json:"img_url"`
	Name        string     `json:"name"`
	Description string     `json:"description"` // without size = text
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
