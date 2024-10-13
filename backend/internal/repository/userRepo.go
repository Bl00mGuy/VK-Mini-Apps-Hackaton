package repository

import "github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"

type UserRepository interface {
	Create(createUserDTO *dto.CreateUserDTO) error
	FindByID(id uint) (*dto.UserDTO, error)
	FindAll() ([]dto.UserDTO, error)
	Update(updateUserDTO *dto.UpdateUserDTO) error
	Delete(id uint) error
}
