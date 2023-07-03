package repository

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/config"
)

type SupplierRepository struct {
	config config.Database
}

func NewSupplierRepository(database config.Database) SupplierRepository {
	return SupplierRepository{
		config: database,
	}
}

// @Summary : Insert Supplier
// @Description : Insert Supplier to database
// @Author : rasmadibbnu
func (r *SupplierRepository) Insert(Supplier entity.Supplier) (entity.Supplier, error) {
	err := r.config.DB.Debug().Create(&Supplier).Error

	if err != nil {
		return Supplier, err
	}

	return Supplier, nil
}

// @Summary : Get Suppliers
// @Description : -
// @Author : rasmadibbnu
func (r *SupplierRepository) FindAll(param map[string]interface{}) ([]entity.Supplier, error) {
	var Suppliers []entity.Supplier

	err := r.config.DB.Where(param).Preload("UserCreate").Find(&Suppliers).Error

	if err != nil {
		return Suppliers, err
	}

	return Suppliers, nil
}

// @Summary : Get Supplier
// @Description : Find Supplier by ID
// @Author : rasmadibbnu
func (r *SupplierRepository) FindById(ID int) (entity.Supplier, error) {
	var supplier entity.Supplier

	err := r.config.DB.First(&supplier).Error

	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

// @Summary : Update Supplier
// @Description : Update Supplier by ID
// @Author : rasmadibbnu
func (r *SupplierRepository) Update(Supplier entity.Supplier, ID int) (entity.Supplier, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&Supplier).Error

	if err != nil {
		return Supplier, err
	}

	return Supplier, nil
}

// @Summary : Delete Supplier
// @Description : Delete Supplier temporary
// @Author : rasmadibbnu
func (r *SupplierRepository) Delete(ID int) (bool, error) {
	var Supplier entity.Supplier

	err := r.config.DB.Where("id = ?", ID).Delete(&Supplier).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
