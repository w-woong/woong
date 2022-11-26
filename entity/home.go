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

	AppConfigID string     `json:"app_config_id" gorm:"column:app_config_id;type:string;size:64;unique;not null"`
	AppConfig   *AppConfig `json:"app_config" gorm:"foreignKey:AppConfigID;references:ID"`

	ShortNoticeList   ShortNoticeList   `json:"short_notice_list" gorm:"foreignKey:HomeID;references:ID"`
	MainPromotionList MainPromotionList `json:"main_promotion_list" gorm:"foreignKey:HomeID;references:ID"`

	// MainProducts ProductList `json:"main_products" gorm:"many2many:main_products;foreignKey:ID;joinForeignKey:HomeID;references:ID;joinReferences:ProductID"`
}

func (e *Home) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type ShortNotice struct {
	ID          string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt   *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"<-:update"`
	HomeID      string     `json:"home_id" gorm:"column:home_id;type:string;size:64;not null"`
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

type MainPromotion struct {
	ID          string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt   *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"<-:update"`
	HomeID      string     `json:"home_id" gorm:"column:home_id;type:string;size:64;not null"`
	ImgUrl      string     `json:"img_url" gorm:"type:string;size:2048"`
	Name        string     `json:"name" gorm:"column:name;type:string;size:256"`
	Description string     `json:"description" gorm:"column:description;type:string"` // without size = text
	// Tags
	Tags Tags `json:"tags" gorm:"polymorphic:Owner"`
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
	ID        string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"<-:update"`
	OwnerType string     `json:"owner_type" gorm:"type:string;size:128;not null"`
	OwnerID   string     `json:"owner_id" gorm:"type:string;size:64;not null"`
	Name      string     `json:"name" gorm:"column:name;type:string;size:128"`
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

type MainProduct struct {
	HomeID    string `json:"home_id" gorm:"column:home_id;primaryKey;type:string;size:64"`
	ProductID string `json:"product_id" gorm:"column:product_id;primaryKey;type:string;size:64"`
}
