package repository

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/config"
)

type ProdukRepository struct {
	config config.Database
}

func NewProdukRepository(database config.Database) ProdukRepository {
	return ProdukRepository{
		config: database,
	}
}

// @Summary : Insert Produk
// @Description : Insert Produk to database
// @Author : rasmadibbnu
func (r *ProdukRepository) Insert(Produk entity.Produk) (entity.Produk, error) {
	err := r.config.DB.Debug().Create(&Produk).Error

	if err != nil {
		return Produk, err
	}

	return Produk, nil
}

// @Summary : Get Produks
// @Description : -
// @Author : rasmadibbnu
func (r *ProdukRepository) FindAll(param map[string]interface{}) ([]entity.Produk, error) {
	var Produks []entity.Produk

	err := r.config.DB.Where(param).Find(&Produks).Error

	if err != nil {
		return Produks, err
	}

	return Produks, nil
}

// @Summary : Get Produk
// @Description : Find Produk by ID
// @Author : rasmadibbnu
func (r *ProdukRepository) FindById(ID int) (entity.Produk, error) {
	var Produk entity.Produk

	err := r.config.DB.First(&Produk).Error

	if err != nil {
		return Produk, err
	}

	return Produk, nil
}

// @Summary : Update Produk
// @Description : Update Produk by ID
// @Author : rasmadibbnu
func (r *ProdukRepository) Update(Produk entity.Produk, ID int) (entity.Produk, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&Produk).Error

	if err != nil {
		return Produk, err
	}

	return Produk, nil
}

// @Summary : Delete Produk
// @Description : Delete Produk temporary
// @Author : rasmadibbnu
func (r *ProdukRepository) Delete(ID int) (bool, error) {
	var Produk entity.Produk

	err := r.config.DB.Where("id = ?", ID).Delete(&Produk).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
