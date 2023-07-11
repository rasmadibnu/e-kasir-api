package repository

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/config"
)

type StokRepository struct {
	config config.Database
}

func NewStokRepository(database config.Database) StokRepository {
	return StokRepository{
		config: database,
	}
}

// @Summary : Insert Stok
// @Description : Insert Stok to database
// @Author : rasmadibbnu
func (r *StokRepository) BacthInsert(Stok []entity.Stok) ([]entity.Stok, error) {
	err := r.config.DB.Create(&Stok).Error

	if err != nil {
		return Stok, err
	}

	return Stok, nil
}

// @Summary : Insert Stok
// @Description : Insert Stok to database
// @Author : rasmadibbnu
func (r *StokRepository) Insert(Stok entity.Stok) (entity.Stok, error) {
	err := r.config.DB.Debug().Create(&Stok).Error

	if err != nil {
		return Stok, err
	}

	return Stok, nil
}

// @Summary : Get Stoks
// @Description : -
// @Author : rasmadibbnu
func (r *StokRepository) FindAll(param map[string]interface{}) ([]entity.Stok, error) {
	var Stoks []entity.Stok

	err := r.config.DB.Where(param).Find(&Stoks).Error

	if err != nil {
		return Stoks, err
	}

	return Stoks, nil
}

// @Summary : Get Stok
// @Description : Find Stok by ID
// @Author : rasmadibbnu
func (r *StokRepository) FindById(ID int) (entity.Stok, error) {
	var Stok entity.Stok

	err := r.config.DB.First(&Stok).Error

	if err != nil {
		return Stok, err
	}

	return Stok, nil
}

// @Summary : Update Stok
// @Description : Update Stok by ID
// @Author : rasmadibbnu
func (r *StokRepository) Update(Stok entity.Stok, ID int) (entity.Stok, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&Stok).Error

	if err != nil {
		return Stok, err
	}

	return Stok, nil
}

// @Summary : Delete Stok
// @Description : Delete Stok temporary
// @Author : rasmadibbnu
func (r *StokRepository) Delete(ID int) (bool, error) {
	var Stok entity.Stok

	err := r.config.DB.Where("id = ?", ID).Delete(&Stok).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
