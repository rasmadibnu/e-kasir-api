package service

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/repository"
)

type StokService struct {
	repository repository.StokRepository
}

func NewStokService(r repository.StokRepository) StokService {
	return StokService{
		repository: r,
	}
}

// @Summary : List Stok
// @Description : Get Stok from repository
// @Author : rasmadibnu
func (s *StokService) List(param map[string]interface{}) ([]entity.Stok, error) {
	Stok, err := s.repository.FindAll(param)

	if err != nil {
		return Stok, err
	}

	return Stok, nil
}

// @Summary : Insert Stok
// @Description : insert Stok to repository
// @Author : rasmadibnu
func (s *StokService) Insert(Stok entity.Stok) (entity.Stok, error) {
	newStok, err := s.repository.Insert(Stok)

	if err != nil {
		return newStok, err
	}

	return newStok, nil
}

// @Summary : Find Stok
// @Description : Find Stok by id from repository
// @Author : rasmadibnu
func (s *StokService) FindById(ID int) (entity.Stok, error) {
	Stok, err := s.repository.FindById(ID)

	if err != nil {
		return Stok, err
	}

	return Stok, nil
}

// @Summary : Update Stok
// @Description : Update Stok by id from repository
// @Author : rasmadibnu
func (s *StokService) Update(Stok entity.Stok, ID int) (entity.Stok, error) {

	updateStok, err := s.repository.Update(Stok, ID)

	if err != nil {
		return updateStok, err
	}

	return updateStok, nil
}

// @Summary : Delete Stok
// @Description : Delete Stok from repository
// @Author : rasmadibnu
func (s *StokService) Delete(ID int) (bool, error) {
	Stok, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return Stok, nil
}
