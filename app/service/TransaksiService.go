package service

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/repository"
)

type TransaksiService struct {
	repository repository.TransaksiRepository
	stokRepo   repository.StokRepository
	cartRepo   repository.CartRepository
	prodRepo   repository.ProdukRepository
}

func NewTransaksiService(r repository.TransaksiRepository, s repository.StokRepository, c repository.CartRepository, p repository.ProdukRepository) TransaksiService {
	return TransaksiService{
		repository: r,
		stokRepo:   s,
		cartRepo:   c,
		prodRepo:   p,
	}
}

// @Summary : List Transaksi
// @Description : Get Transaksi from repository
// @Author : rasmadibnu
func (s *TransaksiService) List(param map[string]interface{}) ([]entity.Transaksi, error) {
	Transaksi, err := s.repository.FindAll(param)

	if err != nil {
		return Transaksi, err
	}

	return Transaksi, nil
}

// @Summary : Insert Transaksi
// @Description : insert Transaksi to repository
// @Author : rasmadibnu
func (s *TransaksiService) Insert(Transaksi entity.Transaksi) (entity.Transaksi, error) {
	newTransaksi, err := s.repository.Insert(Transaksi)

	if err != nil {
		return newTransaksi, err
	}

	for _, detail := range Transaksi.DetailTransaksi {
		produk, err := s.prodRepo.FindById(detail.ProdukID)

		if err != nil {
			return newTransaksi, err
		}
		_, err = s.stokRepo.Insert(entity.Stok{
			ProdukID:     detail.ProdukID,
			Type:         "Sold",
			Value:        detail.JumlahBeli,
			Stok:         produk.Stok.Stok - detail.JumlahBeli,
			UserCreateID: Transaksi.KasirID,
		})

		if err != nil {
			return newTransaksi, err
		}
	}

	_, err = s.cartRepo.ClearCart(newTransaksi.KasirID)

	if err != nil {
		return newTransaksi, err
	}

	return newTransaksi, nil
}

// @Summary : Find Transaksi
// @Description : Find Transaksi by id from repository
// @Author : rasmadibnu
func (s *TransaksiService) FindById(ID int) (entity.Transaksi, error) {
	Transaksi, err := s.repository.FindById(ID)

	if err != nil {
		return Transaksi, err
	}

	return Transaksi, nil
}

// @Summary : Update Transaksi
// @Description : Update Transaksi by id from repository
// @Author : rasmadibnu
func (s *TransaksiService) Update(Transaksi entity.Transaksi, ID int) (entity.Transaksi, error) {

	updateTransaksi, err := s.repository.Update(Transaksi, ID)

	if err != nil {
		return updateTransaksi, err
	}

	return updateTransaksi, nil
}

// @Summary : Delete Transaksi
// @Description : Delete Transaksi from repository
// @Author : rasmadibnu
func (s *TransaksiService) Delete(ID int) (bool, error) {
	Transaksi, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return Transaksi, nil
}
