package mappers

import (
	"api-go/internal/entity"
	"api-go/internal/models"
)

type ModelToEntity interface {
	MapUsersResponse(in []models.Users) *entity.ResponseUsers
}

type modelToEntity struct{}

func NewModelToEntity() ModelToEntity {
	return &modelToEntity{}
}

func (m *modelToEntity) MapUsersResponse(in []models.Users) *entity.ResponseUsers {
	users := make([]*entity.Users, 0)
	for _, item := range in {
		users = append(users, &entity.Users{
			Id:        item.Id,
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Email:     item.Email,
			CreatedOn: item.CreatedOn,
			UpdatedOn: item.UpdatedOn,
			DeletedOn: item.DeletedOn,
		})
	}

	res := &entity.ResponseUsers{
		Status: "Success",
		Data:   users,
	}

	return res
}
