package service

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/repository"
)

type CartService struct {
	repository     repository.CartRepository
	produkRepo     repository.ProdukRepository
	stokRepository repository.StokRepository
}

func NewCartService(r repository.CartRepository, s repository.StokRepository, p repository.ProdukRepository) CartService {
	return CartService{
		repository:     r,
		stokRepository: s,
		produkRepo:     p,
	}
}

// @Summary : List Cart
// @Description : Get Cart from repository
// @Author : rasmadibnu
func (s *CartService) List(param map[string]interface{}) ([]entity.Cart, error) {
	Cart, err := s.repository.FindAll(param)

	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

// @Summary : Insert Cart
// @Description : insert Cart to repository
// @Author : rasmadibnu
func (s *CartService) Insert(Cart entity.Cart, types string) (entity.Cart, error) {
	findCartProd, err := s.repository.FindByProduk(Cart.ProdukID, Cart.UserCreateID)

	if err != nil {
		newCart, err := s.repository.Insert(Cart)
		if err != nil {

			return newCart, nil
		}
		findCartProd = newCart

	} else {
		if types == "add" {
			updateCount, err := s.repository.AddCount(findCartProd.ID)
			if err != nil {

				return updateCount, err
			}

			findCartProd = updateCount
		} else {
			updateCount, err := s.repository.MinCount(findCartProd.ID)
			if err != nil {
				return updateCount, err
			}

			if updateCount.Count == 0 {
				_, err := s.repository.Delete(findCartProd.ID)
				if err != nil {
					return updateCount, err
				}
			}

			findCartProd = updateCount
		}

	}

	produk, err := s.produkRepo.FindById(Cart.ProdukID)

	if err != nil {

		return findCartProd, nil
	}

	if types == "add" {
		_, err = s.stokRepository.Insert(entity.Stok{
			ProdukID:     Cart.ProdukID,
			Type:         "Add Cart",
			Value:        Cart.Count,
			Stok:         produk.Stok.Stok - Cart.Count,
			UserCreateID: Cart.UserCreateID,
		})

		if err != nil {

			return findCartProd, err
		}
	} else {
		_, err = s.stokRepository.Insert(entity.Stok{
			ProdukID:     Cart.ProdukID,
			Type:         "Min Cart",
			Value:        Cart.Count,
			Stok:         produk.Stok.Stok + Cart.Count,
			UserCreateID: Cart.UserCreateID,
		})

		if err != nil {

			return findCartProd, err
		}
	}

	if err != nil {

		return findCartProd, err
	}

	return findCartProd, nil

}

// @Summary : Find Cart
// @Description : Find Cart by id from repository
// @Author : rasmadibnu
func (s *CartService) FindById(ID int) (entity.Cart, error) {
	Cart, err := s.repository.FindById(ID)

	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

// @Summary : Update Cart
// @Description : Update Cart by id from repository
// @Author : rasmadibnu
func (s *CartService) Update(Cart entity.Cart, ID int) (entity.Cart, error) {

	updateCart, err := s.repository.Update(Cart, ID)

	if err != nil {
		return updateCart, err
	}

	return updateCart, nil
}

// @Summary : Delete Cart
// @Description : Delete Cart from repository
// @Author : rasmadibnu
func (s *CartService) Delete(ID int) (bool, error) {
	Cart, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return Cart, nil
}
