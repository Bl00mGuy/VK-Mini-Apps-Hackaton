package services

import "github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"

type UserServiceInterface interface {
	CreateUser(createUserDTO *dto.CreateUserDTO) error
	FindUserByID(id uint) (*dto.UserDTO, error)
	FindAllUsers() ([]dto.UserDTO, error)
	UpdateUser(updateUserDTO *dto.UpdateUserDTO) error
	DeleteUser(id uint) error
}
