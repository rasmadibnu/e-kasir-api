package repository

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/config"

	"gorm.io/gorm"
)

type TransaksiRepository struct {
	config config.Database
}

func NewTransaksiRepository(database config.Database) TransaksiRepository {
	return TransaksiRepository{
		config: database,
	}
}

// @Summary : Insert Transaksi
// @Description : Insert Transaksi to database
// @Author : rasmadibbnu
func (r *TransaksiRepository) Insert(Transaksi entity.Transaksi) (entity.Transaksi, error) {
	err := r.config.DB.Create(&Transaksi).Error

	if err != nil {
		return Transaksi, err
	}

	return Transaksi, nil
}

// @Summary : Get Transaksis
// @Description : -
// @Author : rasmadibbnu
func (r *TransaksiRepository) FindAll(param map[string]interface{}) ([]entity.Transaksi, error) {
	var Transaksis []entity.Transaksi

	err := r.config.DB.Where(param).Preload("UserCreate").Preload("DetailTransaksi.Produk", func(db *gorm.DB) *gorm.DB {
		return db.Preload("UserCreate").Preload("Stok.UserCreate").Preload("Supplier.UserCreate")
	}).Order("id desc").Find(&Transaksis).Error

	if err != nil {
		return Transaksis, err
	}

	return Transaksis, nil
}

// @Summary : Get Transaksi
// @Description : Find Transaksi by ID
// @Author : rasmadibbnu
func (r *TransaksiRepository) FindById(ID int) (entity.Transaksi, error) {
	var Transaksi entity.Transaksi

	err := r.config.DB.Preload("DetailTransaksi.Produk", func(db *gorm.DB) *gorm.DB {
		return db.Preload("UserCreate").Preload("Stok.UserCreate").Preload("Supplier.UserCreate")
	}).First(&Transaksi).Error

	if err != nil {
		return Transaksi, err
	}

	return Transaksi, nil
}

// @Summary : Update Transaksi
// @Description : Update Transaksi by ID
// @Author : rasmadibbnu
func (r *TransaksiRepository) Update(Transaksi entity.Transaksi, ID int) (entity.Transaksi, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&Transaksi).Error

	if err != nil {
		return Transaksi, err
	}

	return Transaksi, nil
}

// @Summary : Delete Transaksi
// @Description : Delete Transaksi temporary
// @Author : rasmadibbnu
func (r *TransaksiRepository) Delete(ID int) (bool, error) {
	var Transaksi entity.Transaksi

	err := r.config.DB.Where("id = ?", ID).Delete(&Transaksi).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
