package models

import (
	"time"
)

type Category struct {
	Id          int        `gorm:"primary_key" json:"id,omitempty"`
	Name        string     `gorm:"type:varchar(20);not null" json:"name"`
	Description string     `gorm:"type:varchar(20)" json:"description"`
	Providers   []Provider `gorm:"foreignKey:CategoryId;references:Id"`
	CreatedAt   *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at_at,string,omitempty"`
}

type CategoryOutput struct {
	Id          int        `gorm:"primary_key" json:"id,omitempty"`
	Name        string     `gorm:"type:varchar(20);not null" json:"name"`
	Description string     `gorm:"type:varchar(20)" json:"description"`
	CreatedAt   *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at_at,string,omitempty"`
}

func (e *Category) TableName() string {
	return "categories"
}

func (e *CategoryOutput) TableName() string {
	return "categories"
}
