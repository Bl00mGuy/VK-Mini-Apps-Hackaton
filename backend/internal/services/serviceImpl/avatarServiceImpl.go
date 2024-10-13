package serviceImpl

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository"
)

type AvatarService struct {
	repo repository.AvatarRepository
}

func NewAvatarService(AvatarRepository repository.AvatarRepository) *AvatarService {
	return &AvatarService{
		repo: AvatarRepository,
	}
}

func (s *AvatarService) CreateAvatar(createAvatarDTO *dto.CreateAvatarDTO) error {
	return s.repo.Create(createAvatarDTO)
}

func (s *AvatarService) FindAvatarByID(id uint) (*dto.AvatarDTO, error) {
	return s.repo.FindByID(id)
}

func (s *AvatarService) FindAllAvatars(userID uint) ([]dto.AvatarDTO, error) {
	return s.repo.FindAll(userID)
}

func (s *AvatarService) UpdateAvatar(updateAvatarDTO *dto.UpdateAvatarDTO) error {
	return s.repo.Update(updateAvatarDTO)
}

func (s *AvatarService) DeleteAvatar(id uint) error {
	return s.repo.Delete(id)
}
