package repository

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/config"

	"gorm.io/gorm"
)

type CartRepository struct {
	config config.Database
}

func NewCartRepository(database config.Database) CartRepository {
	return CartRepository{
		config: database,
	}
}

// @Summary : Insert Cart
// @Description : Insert Cart to database
// @Author : rasmadibbnu
func (r *CartRepository) Insert(Cart entity.Cart) (entity.Cart, error) {
	err := r.config.DB.Debug().Create(&Cart).Error

	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

// @Summary : Get Carts
// @Description : -
// @Author : rasmadibbnu
func (r *CartRepository) FindAll(param map[string]interface{}) ([]entity.Cart, error) {
	var Carts []entity.Cart

	err := r.config.DB.Where(param).Where("count != 0").Preload("Produk", func(db *gorm.DB) *gorm.DB {
		return db.Preload("UserCreate").Preload("Stok.UserCreate").Preload("Supplier.UserCreate")
	}).Preload("UserCreate").Find(&Carts).Error

	if err != nil {
		return Carts, err
	}

	return Carts, nil
}

// @Summary : Get Cart
// @Description : Find Cart by ID
// @Author : rasmadibbnu
func (r *CartRepository) FindById(ID int) (entity.Cart, error) {
	var Cart entity.Cart

	err := r.config.DB.First(&Cart).Error

	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

// @Summary : Get Cart
// @Description : Find Cart by Produk
// @Author : rasmadibbnu
func (r *CartRepository) FindByProduk(ID int, UserID int) (entity.Cart, error) {
	var Cart entity.Cart

	err := r.config.DB.Preload("Produk.Stok").Where("produk_id = ?", ID).Where("created_by = ?", UserID).First(&Cart).Error

	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

// @Summary : Get update count
// @Description : Find  update count
// @Author : rasmadibbnu
func (r *CartRepository) AddCount(ID int) (entity.Cart, error) {
	var Cart entity.Cart

	err := r.config.DB.Model(&Cart).Where("id = ?", ID).UpdateColumn("count", gorm.Expr("count + ?", 1)).Error

	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

// @Summary : Get update count
// @Description : Find  update count
// @Author : rasmadibbnu
func (r *CartRepository) MinCount(ID int) (entity.Cart, error) {
	var Cart entity.Cart

	err := r.config.DB.Model(&Cart).Where("id = ?", ID).UpdateColumn("count", gorm.Expr("count - ?", 1)).Error

	if err != nil {
		return Cart, err
	}

	findcart, err := r.FindById(ID)

	if err != nil {
		return Cart, err
	}

	return findcart, nil
}

// @Summary : Update Cart
// @Description : Update Cart by ID
// @Author : rasmadibbnu
func (r *CartRepository) Update(Cart entity.Cart, ID int) (entity.Cart, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&Cart).Error

	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

// @Summary : Delete Cart
// @Description : Delete Cart temporary
// @Author : rasmadibbnu
func (r *CartRepository) Delete(ID int) (bool, error) {
	var Cart entity.Cart

	err := r.config.DB.Where("id = ?", ID).Delete(&Cart).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// @Summary : Clear Cart
// @Description : Clear Cart
// @Author : rasmadibbnu
func (r *CartRepository) ClearCart(ID int) (bool, error) {
	var Cart entity.Cart

	err := r.config.DB.Debug().Unscoped().Where("created_by = ?", ID).Delete(&Cart).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
