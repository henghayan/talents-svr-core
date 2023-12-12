package application

import "core/pkg/domain"

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*domain.User, error) {
	return s.repo.GetUserByID(id)
}
