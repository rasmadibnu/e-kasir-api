package service

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/repository"
)

type KategoriService struct {
	repository repository.KategoriRepository
}

func NewKategoriService(r repository.KategoriRepository) KategoriService {
	return KategoriService{
		repository: r,
	}
}

// @Summary : List Kategori
// @Description : Get Kategori from repository
// @Author : rasmadibnu
func (s *KategoriService) List(param map[string]interface{}) ([]entity.Kategori, error) {
	Kategori, err := s.repository.FindAll(param)

	if err != nil {
		return Kategori, err
	}

	return Kategori, nil
}

// @Summary : Insert Kategori
// @Description : insert Kategori to repository
// @Author : rasmadibnu
func (s *KategoriService) Insert(Kategori entity.Kategori) (entity.Kategori, error) {
	newKategori, err := s.repository.Insert(Kategori)

	if err != nil {
		return newKategori, err
	}

	return newKategori, nil
}

// @Summary : Find Kategori
// @Description : Find Kategori by id from repository
// @Author : rasmadibnu
func (s *KategoriService) FindById(ID int) (entity.Kategori, error) {
	Kategori, err := s.repository.FindById(ID)

	if err != nil {
		return Kategori, err
	}

	return Kategori, nil
}

// @Summary : Update Kategori
// @Description : Update Kategori by id from repository
// @Author : rasmadibnu
func (s *KategoriService) Update(Kategori entity.Kategori, ID int) (entity.Kategori, error) {

	updateKategori, err := s.repository.Update(Kategori, ID)

	if err != nil {
		return updateKategori, err
	}

	return updateKategori, nil
}

// @Summary : Delete Kategori
// @Description : Delete Kategori from repository
// @Author : rasmadibnu
func (s *KategoriService) Delete(ID int) (bool, error) {
	Kategori, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return Kategori, nil
}
