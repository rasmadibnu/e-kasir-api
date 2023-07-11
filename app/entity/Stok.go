package entity

import (
	"time"
)

type Stok struct {
	ID           int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	ProdukID     int       `gorm:"column:produk_id;type:int" json:"produk_id"`
	Stok         int       `gorm:"column:stok;type:int(11)" json:"stok"`
	Value        int       `gorm:"column:value;type:int(11)" json:"value"`
	Type         string    `gorm:"column:type;type:varchar(255)" json:"type"`
	UserCreateID int       `gorm:"column:created_by;type:int(11)" json:"created_by"`
	UserCreate   *User     `gorm:"foreignKey:UserCreateID" json:"user_create"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
}

func (m *Stok) TableName() string {
	return "stok"
}
