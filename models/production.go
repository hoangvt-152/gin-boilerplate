package models

import (
	"gorm.io/gorm"
	"time"
)

type ProviderProduction struct {
	gorm.Model
	Id           int           `gorm:"primary_key" json:"id,omitempty"`
	ProviderId   int           `json:"provider_id"`
	Name         string        `gorm:"type:varchar(255);not null" json:"name"`
	Poster       string        `gorm:"type:text;not null" json:"poster"`
	Avatar       string        `gorm:"type:text;not null" json:"avatar"`
	MediaVideo   string        `gorm:"type:text;not null" json:"house_number"`
	MediaImage   string        `gorm:"type:text;not null" json:"street"`
	Description  string        `gorm:"type:text;not null" json:"description"`
	CreatedAt    *time.Time    `json:"created_at,string,omitempty"`
	UpdatedAt    *time.Time    `json:"updated_at,string,omitempty"`
	ProductDeals []ProductDeal `gorm:"foreignKey:ProductId;references:Id"`
}

func (e *ProviderProduction) TableName() string {
	return "provider_productions"
}
