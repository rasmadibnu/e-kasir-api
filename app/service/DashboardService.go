package service

import (
	"kasir-cepat-api/app/repository"
)

type DashboardService struct {
	repository repository.DashboardRepository
}

func NewDashboardService(r repository.DashboardRepository) DashboardService {
	return DashboardService{
		repository: r,
	}
}

// @Summary : Get dashboard
// @Description : Get dashboard from repository
// @Author : rasmadibnu
func (s *DashboardService) List(param map[string]interface{}) (map[string]interface{}, error) {
	m, err := s.repository.FindAll(param)

	if err != nil {
		return m, err
	}

	return m, nil
}
