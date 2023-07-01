package repository

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/config"
)

type KategoriRepository struct {
	config config.Database
}

func NewKategoriRepository(database config.Database) KategoriRepository {
	return KategoriRepository{
		config: database,
	}
}

// @Summary : Insert Kategori
// @Description : Insert Kategori to database
// @Author : rasmadibbnu
func (r *KategoriRepository) Insert(Kategori entity.Kategori) (entity.Kategori, error) {
	err := r.config.DB.Debug().Create(&Kategori).Error

	if err != nil {
		return Kategori, err
	}

	return Kategori, nil
}

// @Summary : Get Kategoris
// @Description : -
// @Author : rasmadibbnu
func (r *KategoriRepository) FindAll(param map[string]interface{}) ([]entity.Kategori, error) {
	var Kategoris []entity.Kategori

	err := r.config.DB.Where(param).Find(&Kategoris).Error

	if err != nil {
		return Kategoris, err
	}

	return Kategoris, nil
}

// @Summary : Get Kategori
// @Description : Find Kategori by ID
// @Author : rasmadibbnu
func (r *KategoriRepository) FindById(ID int) (entity.Kategori, error) {
	var Kategori entity.Kategori

	err := r.config.DB.First(&Kategori).Error

	if err != nil {
		return Kategori, err
	}

	return Kategori, nil
}

// @Summary : Update Kategori
// @Description : Update Kategori by ID
// @Author : rasmadibbnu
func (r *KategoriRepository) Update(Kategori entity.Kategori, ID int) (entity.Kategori, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&Kategori).Error

	if err != nil {
		return Kategori, err
	}

	return Kategori, nil
}

// @Summary : Delete Kategori
// @Description : Delete Kategori temporary
// @Author : rasmadibbnu
func (r *KategoriRepository) Delete(ID int) (bool, error) {
	var Kategori entity.Kategori

	err := r.config.DB.Where("id = ?", ID).Delete(&Kategori).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
