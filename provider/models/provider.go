package models
import (
	"time"
)

type Provider struct {
	Id        		int        `json:"id"`
	CategoryId      int        `json:"category_id"`
	Name			string     `json:"name"`
	Poster          string     `json:"poster"`
	Avatar          string     `json:"avatar"`
	HouseNumber     string     `json:"house_number"`
	Street          string     `json:"street"`
	PhoneNumber     string     `json:"phone_number"`
	LatLocation     float64    `json:"lat_location"`
	LonLocation     float64    `json:"lon_location"` 
    InstagramName   string     `json:"instagram_name"`  
	Description		string     `json:"description"`
	OtherMedia      string     `json:"other_media"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}

func (e *Provider) TableName() string {
	return "providers"
}