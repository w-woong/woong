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

	ShortNoticeList   ShortNoticeList   `json:"short_notice_list"`
	MainPromotionList MainPromotionList `json:"main_promotion_list"`
	// MainProducts      ProductList       `json:"main_products"`
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

type MainPromotion struct {
	ID          string     `json:"id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	HomeID      string     `json:"home_id"`
	ImgUrl      string     `json:"img_url"`
	Name        string     `json:"name"`
	Description string     `json:"description"` // without size = text
	Tags        Tags       `json:"tags"`
}

func (e *MainPromotion) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type MainPromotionList []MainPromotion

func (e *MainPromotionList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type Tag struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	OwnerType string     `json:"owner_type"`
	OwnerID   string     `json:"owner_id"`
	Name      string     `json:"name"`
}

func (e *Tag) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type Tags []Tag

func (e *Tags) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
