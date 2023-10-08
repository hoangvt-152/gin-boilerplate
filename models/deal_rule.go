package models

import (
	"gorm.io/gorm"
	"time"
)

type DealRule struct {
	gorm.Model
	Id        int        `gorm:"primary_key" json:"id,omitempty"`
	RuleId    int        `gorm:"not null" json:"rule_id"`
	PDealId   int        `gorm:"not null" json:"p_deal_id"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
}

func (e *DealRule) TableName() string {
	return "deal_rules"
}
