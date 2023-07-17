package repository

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/config"
)

type DashboardRepository struct {
	config config.Database
}

func NewDashboardRepository(database config.Database) DashboardRepository {
	return DashboardRepository{
		config: database,
	}
}

// @Summary : Get Dahsboard
// @Description : -
// @Author : rasmadibnu
func (r *DashboardRepository) FindAll(param map[string]interface{}) (map[string]interface{}, error) {

	m := make(map[string]interface{})

	var countKasir int64
	kasir := r.config.DB.Model(&entity.User{}).Where("role = ?", "Kasir").Count(&countKasir).Error
	m["kasir"] = countKasir
	if kasir != nil {
		return m, kasir
	}

	var countProduk int64
	produk := r.config.DB.Model(&entity.Produk{}).Count(&countProduk).Error
	m["produk"] = countProduk
	if produk != nil {
		return m, produk
	}

	var countSupplier int64
	supplier := r.config.DB.Model(&entity.Supplier{}).Count(&countSupplier).Error
	m["supplier"] = countSupplier
	if supplier != nil {
		return m, supplier
	}

	return m, nil
}
