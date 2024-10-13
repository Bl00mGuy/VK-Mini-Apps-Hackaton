package mocks

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(userDTO *dto.CreateUserDTO) error {
	args := m.Called(userDTO)
	return args.Error(0)
}

func (m *MockUserService) FindUserByID(id uint) (*dto.UserDTO, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.UserDTO), args.Error(1)
}

func (m *MockUserService) FindAllUsers() ([]dto.UserDTO, error) {
	args := m.Called()
	return args.Get(0).([]dto.UserDTO), args.Error(1)
}

func (m *MockUserService) UpdateUser(userDTO *dto.UpdateUserDTO) error {
	args := m.Called(userDTO)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(createUserDTO *dto.CreateUserDTO) error {
	args := m.Called(createUserDTO)
	return args.Error(0)
}

func (m *MockUserRepository) FindByID(id uint) (*dto.UserDTO, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.UserDTO), args.Error(1)
}

func (m *MockUserRepository) FindAll() ([]dto.UserDTO, error) {
	args := m.Called()
	return args.Get(0).([]dto.UserDTO), args.Error(1)
}

func (m *MockUserRepository) Update(updateUserDTO *dto.UpdateUserDTO) error {
	args := m.Called(updateUserDTO)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
