package entity

import (
	"time"
)

type Produk struct {
	ID         int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	SupplierID int       `gorm:"column:supplier_id;type:int(11)" json:"supplier_id"`
	Image      string    `gorm:"column:image;type:varchar(255)" json:"image"`
	Name       string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Deskripsi  string    `gorm:"column:deskripsi;type:varchar(255)" json:"deskripsi"`
	Kategori   string    `gorm:"column:kategori;type:varchar(255)" json:"kategori"`
	CreatedBy  int       `gorm:"column:created_by;type:int(11)" json:"created_by"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

func (m *Produk) TableName() string {
	return "produk"
}
