package entity

import (
	"time"

	"gorm.io/gorm"
)

type Kategori struct {
	ID           int            `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Name         string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Deskripsi    *string        `gorm:"column:deskripsi;type:varchar(255)" json:"deskripsi"`
	Produk       []Produk       `json:"produk"`
	UserCreateID int            `gorm:"column:created_by;type:int(11)" json:"created_by"`
	UserCreate   User           `gorm:"foreignKey:UserCreateID" json:"user_create"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

func (m *Kategori) TableName() string {
	return "kategori"
}
