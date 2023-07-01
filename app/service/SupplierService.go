package service

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/repository"
)

type SupplierService struct {
	repository repository.SupplierRepository
}

func NewSupplierService(r repository.SupplierRepository) SupplierService {
	return SupplierService{
		repository: r,
	}
}

// @Summary : List supplier
// @Description : Get supplier from repository
// @Author : rasmadibnu
func (s *SupplierService) List(param map[string]interface{}) ([]entity.Supplier, error) {
	supplier, err := s.repository.FindAll(param)

	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

// @Summary : Insert supplier
// @Description : insert supplier to repository
// @Author : rasmadibnu
func (s *SupplierService) Insert(supplier entity.Supplier) (entity.Supplier, error) {
	newSupplier, err := s.repository.Insert(supplier)

	if err != nil {
		return newSupplier, err
	}

	return newSupplier, nil
}

// @Summary : Find supplier
// @Description : Find supplier by id from repository
// @Author : rasmadibnu
func (s *SupplierService) FindById(ID int) (entity.Supplier, error) {
	supplier, err := s.repository.FindById(ID)

	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

// @Summary : Update supplier
// @Description : Update supplier by id from repository
// @Author : rasmadibnu
func (s *SupplierService) Update(supplier entity.Supplier, ID int) (entity.Supplier, error) {

	updateSupplier, err := s.repository.Update(supplier, ID)

	if err != nil {
		return updateSupplier, err
	}

	return updateSupplier, nil
}

// @Summary : Delete supplier
// @Description : Delete supplier from repository
// @Author : rasmadibnu
func (s *SupplierService) Delete(ID int) (bool, error) {
	supplier, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return supplier, nil
}
