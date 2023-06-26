package service

import (
	"kasir-cepat-api/app/entity"
	"kasir-cepat-api/app/repository"
	"kasir-cepat-api/security"
)

type AuthService struct {
	repository repository.UserRepository
}

func NewAuthService(r repository.UserRepository) AuthService {
	return AuthService{
		repository: r,
	}
}

func (s *AuthService) Login(req entity.User) (entity.User, error) {
	reqUsername := req.Username
	reqPassword := req.Password

	user, err := s.repository.FindByUsername(reqUsername)

	if err != nil {
		return user, err
	}

	checkPassword := security.CheckPassword(user.Password, reqPassword)

	if checkPassword == nil {
		return user, nil
	}

	return user, checkPassword
}
