package models
import (
	"time"
)

type Category struct {
	Id        		int        `json:"id"`
	Name			string     `json:"name"`
	Description		string     `json:"description"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}

func (e *Category) TableName() string {
	return "categories"
}