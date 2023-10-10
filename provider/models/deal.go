package models

import (
	"time"
)

type ProductDeal struct {
	Id             int        `gorm:"primary_key" json:"id,omitempty"`
	ProductId      int        `json:"product_id"`
	State          int        `json:"state"`
	DiscountNumber float64    `gorm:"type:Numeric(5,0);not null" json:"discount_number"`
	DiscountType   string     `gorm:"type:varchar(10);not null" json:"discount_type"`
	Count          int        `gorm:"not null" json:"count"`
	ExpiredTime    int        `gorm:"not null" json:"expired_time"`
	VoucherImage   string     `gorm:"type:varchar(255);not null" json:"voucher_image"`
	CreatedAt      *time.Time `json:"created_at,string,omitempty"`
	DealRules      []DealRule `gorm:"foreignKey:PDealId;references:Id"` //`gorm:"foreignKey:PDealId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (e *ProductDeal) TableName() string {
	return "product_deals"
}
