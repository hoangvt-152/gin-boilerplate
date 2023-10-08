package models

import (
	"gorm.io/gorm"
	"time"
)

type ProviderRule struct {
	gorm.Model
	Id          int        `gorm:"primary_key" json:"id,omitempty"`
	ProviderId  int        `json:"provider_id,omitempty"`
	Name        string     `gorm:"type:varchar(20);not null" json:"name"`
	Description string     `gorm:"type:varchar(20);not null" json:"description"`
	SocialType  string     `gorm:"type:varchar(20);not null" json:"social_type"`
	StepNumber  int        `json:"step_number"`
	CreatedAt   *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at_at,string,omitempty"`
	DealRules   []DealRule  `gorm:"foreignKey:RuleId;references:Id"`  //`gorm:"foreignKey:RuleId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (e *ProviderRule) TableName() string {
	return "provider_rules"
}
