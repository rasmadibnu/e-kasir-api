package entity

import (
	"time"
)

type Stok struct {
	ID        int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	ProdukID  string    `gorm:"column:produk_id;type:varchar(255)" json:"produk_id"`
	Stok      int       `gorm:"column:stok;type:int(11)" json:"stok"`
	LastStok  int       `gorm:"column:last_stok;type:int(11)" json:"last_stok"`
	Value     int       `gorm:"column:value;type:int(11)" json:"value"`
	Type      string    `gorm:"column:type;type:varchar(255)" json:"type"`
	CreatedBy int       `gorm:"column:created_by;type:int(11)" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
}

func (m *Stok) TableName() string {
	return "stok"
}
