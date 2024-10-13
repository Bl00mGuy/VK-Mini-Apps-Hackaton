package serviceImpl

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(UserRepository repository.UserRepository) *UserService {
	return &UserService{
		repo: UserRepository,
	}
}

func (s *UserService) CreateUser(createUserDTO *dto.CreateUserDTO) error {
	return s.repo.Create(createUserDTO)
}

func (s *UserService) FindUserByID(id uint) (*dto.UserDTO, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) FindAllUsers() ([]dto.UserDTO, error) {
	return s.repo.FindAll()
}

func (s *UserService) UpdateUser(updateUserDTO *dto.UpdateUserDTO) error {
	return s.repo.Update(updateUserDTO)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
