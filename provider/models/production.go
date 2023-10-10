package models

import (
	"time"
)

type ProviderProduction struct {
	Id           int           `gorm:"primary_key" json:"id,omitempty"`
	ProviderId   int           `json:"provider_id"`
	Name         string        `gorm:"type:varchar(255);not null" json:"name"`
	Poster       string        `gorm:"type:text;not null" json:"poster"`
	Avatar       string        `gorm:"type:text;not null" json:"avatar"`
	MediaVideo   string        `gorm:"type:text;not null" json:"media_video"`
	MediaImage   string        `gorm:"type:text;not null" json:"media_image"`
	Description  string        `gorm:"type:text;not null" json:"description"`
	CreatedAt    *time.Time    `json:"created_at,string,omitempty"`
	UpdatedAt    *time.Time    `json:"updated_at,string,omitempty"`
	ProductDeals []ProductDeal `gorm:"foreignKey:ProductId;references:Id"`
}

func (e *ProviderProduction) TableName() string {
	return "provider_productions"
}
