package service

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/repository"
)

type ProdukService struct {
	repository repository.ProdukRepository
}

func NewProdukService(r repository.ProdukRepository) ProdukService {
	return ProdukService{
		repository: r,
	}
}

// @Summary : List Produk
// @Description : Get Produk from repository
// @Author : rasmadibnu
func (s *ProdukService) List(param map[string]interface{}) ([]entity.Produk, error) {
	Produk, err := s.repository.FindAll(param)

	if err != nil {
		return Produk, err
	}

	return Produk, nil
}

// @Summary : Insert Produk
// @Description : insert Produk to repository
// @Author : rasmadibnu
func (s *ProdukService) Insert(Produk entity.Produk) (entity.Produk, error) {
	newProduk, err := s.repository.Insert(Produk)

	if err != nil {
		return newProduk, err
	}

	return newProduk, nil
}

// @Summary : Find Produk
// @Description : Find Produk by id from repository
// @Author : rasmadibnu
func (s *ProdukService) FindById(ID int) (entity.Produk, error) {
	Produk, err := s.repository.FindById(ID)

	if err != nil {
		return Produk, err
	}

	return Produk, nil
}

// @Summary : Update Produk
// @Description : Update Produk by id from repository
// @Author : rasmadibnu
func (s *ProdukService) Update(Produk entity.Produk, ID int) (entity.Produk, error) {

	updateProduk, err := s.repository.Update(Produk, ID)

	if err != nil {
		return updateProduk, err
	}

	return updateProduk, nil
}

// @Summary : Delete Produk
// @Description : Delete Produk from repository
// @Author : rasmadibnu
func (s *ProdukService) Delete(ID int) (bool, error) {
	Produk, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return Produk, nil
}
