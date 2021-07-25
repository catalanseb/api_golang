package users

import (
	"api-go/internal/entity"
	"api-go/internal/mappers"
	"api-go/internal/repositories"
	"fmt"
)

// interface
type ServiceUsers interface {
	GetAll() (*entity.ResponseUsers, error)
}

type serviceUsers struct {
	repository repositories.UserRepository
}

func NewServiceUsers(repository repositories.UserRepository) ServiceUsers {
	return &serviceUsers{
		repository: repository,
	}
}

func (s *serviceUsers) GetAll() (*entity.ResponseUsers, error) {
	res, err := s.repository.GetAllUsers()
	if err != nil {
		fmt.Printf("Error: %v \n\n", err)
		return nil, err
	}

	mte := mappers.NewModelToEntity()
	users := mte.MapUsersResponse(res)

	return users, nil
}
