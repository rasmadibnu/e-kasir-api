package entity

import (
	"fmt"
	"time"

	"github.com/yudapc/go-rupiah"
	"gorm.io/gorm"
)

type Transaksi struct {
	ID              int               `gorm:"column:id;type:int(11);primary_key" json:"id"`
	NoTransaksi     string            `gorm:"-" json:"no_transaksi"`
	Diskon          int               `gorm:"column:diskon;type:int(11)" json:"diskon"`
	Ppn             int               `gorm:"column:ppn;type:int(11)" json:"ppn"`
	TotalBelanja    int               `gorm:"column:total_belanja;type:int(11)" json:"total_belanja"`
	TotalBelanjaRp  string            `gorm:"-" json:"total_belanja_rp"`
	KasirID         int               `gorm:"column:kasir_id;type:int(11)" json:"kasir_id"`
	UserCreate      *User             `gorm:"foreignKey:KasirID" json:"user_create"`
	DetailTransaksi []DetailTransaksi `gorm:"foreginKey:TransaksiID" json:"detail"`
	CreatedAt       time.Time         `gorm:"column:tanggal;type:timestamp" json:"tanggal"`
}

func (m *Transaksi) TableName() string {
	return "transaksi"
}

func (u *Transaksi) AfterFind(tx *gorm.DB) (err error) {
	u.NoTransaksi = u.CreatedAt.Format("02012006") + fmt.Sprintf("%06d", u.ID)
	amount := u.TotalBelanja
	amountFloat := float64(amount)
	u.TotalBelanjaRp = rupiah.FormatRupiah(amountFloat)
	return
}

type DetailTransaksi struct {
	ID          int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	TransaksiID int    `gorm:"column:transaksi_id;type:int(11)" json:"transaksi_id"`
	ProdukID    int    `gorm:"column:produk_id;type:int(11)" json:"produk_id"`
	Produk      Produk `json:"produk"`
	JumlahBeli  int    `gorm:"column:jumlah_beli;type:int(11)" json:"jumlah_beli"`
}

func (m *DetailTransaksi) TableName() string {
	return "detail_transaksi"
}
