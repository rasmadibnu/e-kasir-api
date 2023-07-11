package entity

import (
	"time"
)

type Cart struct {
	ID           int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	ProdukID     int       `gorm:"column:produk_id" json:"produk_id"`
	Produk       Produk    `json:"produk"`
	Count        int       `json:"count"`
	UserCreateID int       `gorm:"column:created_by;type:int(11)" json:"created_by"`
	UserCreate   User      `gorm:"foreignKey:UserCreateID" json:"user_create"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

func (m *Cart) TableName() string {
	return "cart"
}
