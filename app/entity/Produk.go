package entity

import (
	"time"

	"github.com/yudapc/go-rupiah"
	"gorm.io/gorm"
)

type Produk struct {
	ID           int            `gorm:"column:id;type:int(11);primary_key" json:"id"`
	SupplierID   int            `gorm:"column:supplier_id;type:int(11)" json:"supplier_id"`
	Image        string         `gorm:"column:image;type:varchar(255)" json:"image"`
	Name         string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Deskripsi    string         `gorm:"column:deskripsi;type:varchar(255)" json:"deskripsi"`
	KategoriID   int            `gorm:"column:kategori_id;type:varchar(255)" json:"kategori_id"`
	Harga        int            `gorm:"column:harga;type:int(11)" json:"harga"`
	HargaRp      string         `gorm:"-" json:"harga_rp"`
	Supplier     Supplier       `json:"supplier"`
	Kategori     Kategori       `json:"kategori"`
	Stok         Stok           `json:"stok"`
	UserCreateID int            `gorm:"column:created_by;type:int(11)" json:"created_by"`
	UserCreate   *User          `gorm:"foreignKey:UserCreateID" json:"user_create"`
	UserUpdateID int            `gorm:"column:created_by;type:int(11)" json:"updated_by"`
	UserUpdate   *User          `gorm:"foreignKey:UserUpdateID" json:"user_update"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

func (m *Produk) TableName() string {
	return "produk"
}

func (u *Produk) AfterFind(tx *gorm.DB) (err error) {
	amount := u.Harga
	amountFloat := float64(amount)
	u.HargaRp = rupiah.FormatRupiah(amountFloat)
	return
}
