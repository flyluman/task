package service

import (
	"task/internal/repository"
	"task/model"
)

type UserService interface {
	GetUserRestaurants(userID int) ([]model.Restaurant, error)
	PurchaseMenuItem(userID, menuItemID int) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetUserRestaurants(userID int) ([]model.Restaurant, error) {
	restaurants, err := s.repo.QueryRestaurants(userID)

	if err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (s *userService) PurchaseMenuItem(userID, menuItemID int) error {
	return s.repo.PurchaseTX(userID, menuItemID)
}
