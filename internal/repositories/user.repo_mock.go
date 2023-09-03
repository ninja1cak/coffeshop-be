package repositories

import (
	"ninja1cak/coffeshop-be/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoUserMock struct {
	mock.Mock
}

func (r *RepoUserMock) CreateUser(data *models.User) (string, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}

func (r *RepoUserMock) GetUser(user_id string) (*[]models.User, error) {
	args := r.Mock.Called(user_id)
	return args.Get(0).(*[]models.User), args.Error(1)
}

func (r *RepoUserMock) UpdateUser(data *models.User) (string, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}

func (r *RepoUserMock) DeleteUser(data *models.User) (string, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}

func (r *RepoUserMock) GetAuthData(email string) (*models.User, error) {
	args := r.Mock.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (r *RepoUserMock) UpdateStatusUser(email string) (string, error) {
	args := r.Mock.Called(email)
	return args.Get(0).(string), args.Error(1)
}
