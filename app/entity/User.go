package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int64          `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Username     string         `gorm:"column:username;type:varchar(100);NOT NULL" json:"username"`
	Name         string         `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`
	PhoneNumber  string         `gorm:"column:phone_number;type:varchar(100);NOT NULL" json:"phone_number"`
	JenisKelamin string         `gorm:"column:jenis_kelamin;type:varchar(1);NOT NULL" json:"jenis_kelamin"`
	Password     string         `gorm:"column:password;type:varchar(100);NOT NULL" json:"password"`
	Role         string         `gorm:"column:role;type:varchar(100);NOT NULL" json:"role"`
	CreatedBy    int            `gorm:"column:created_by;type:int(11)" json:"created_by"`
	Cart         []Cart         `gorm:"foreignKey:UserCreateID" json:"cart"`
	Transaksi    []Transaksi    `gorm:"foreignKey:KasirID" json:"transaksi"`
	UserCreate   *User          `gorm:"foreignKey:CreatedBy" json:"user_create"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func (m *User) TableName() string {
	return "users"
}
