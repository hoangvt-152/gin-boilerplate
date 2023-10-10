package models

import (
	"time"
)

type Provider struct {
	Id            int                  `gorm:"primary_key" json:"id,omitempty"`
	CategoryId    int                  `json:"category_id"`
	Name          string               `gorm:"type:varchar(255);not null" json:"name"`
	Poster        string               `gorm:"type:text;not null" json:"poster"`
	Avatar        string               `gorm:"type:text;not null" json:"avatar"`
	HouseNumber   string               `gorm:"type:varchar(255);not null" json:"house_number"`
	Street        string               `gorm:"type:varchar(255);not null" json:"street"`
	PhoneNumber   string               `gorm:"type:varchar(255);not null" json:"phone_number"`
	LatLocation   float64              `json:"lat_location"`
	LonLocation   float64              `json:"lon_location"`
	InstagramName string               `gorm:"type:varchar(255);not null" json:"instagram_name"`
	Description   string               `gorm:"type:text;not null" json:"description"`
	OtherMedia    string               `gorm:"type:text;not null" json:"other_media"`
	CreatedAt     *time.Time           `json:"created_at,string,omitempty"`
	UpdatedAt     *time.Time           `json:"updated_at_at,string,omitempty"`
	ProviderRules []ProviderRule       `gorm:"foreignKey:ProviderId;references:Id"`
	Productions   []ProviderProduction `gorm:"foreignKey:ProviderId;references:Id"`
}

func (e *Provider) TableName() string {
	return "providers"
}
