package impl

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository"
)

type AchievementService struct {
	repo repository.AchievementRepository
}

func NewAchievementService(achievementRepository repository.AchievementRepository) *AchievementService {
	return &AchievementService{
		repo: achievementRepository,
	}
}

func (s *AchievementService) CreateAchievement(createAchievementDTO *dto.CreateAchievementDTO) error {
	return s.repo.Create(createAchievementDTO)
}

func (s *AchievementService) FindAchievementByID(id uint) (*dto.AchievementDTO, error) {
	return s.repo.FindByID(id)
}

func (s *AchievementService) FindAllAchievements(userID uint) ([]dto.AchievementDTO, error) {
	return s.repo.FindAll(userID)
}

func (s *AchievementService) UpdateAchievement(updateAchievementDTO *dto.UpdateAchievementDTO) error {
	return s.repo.Update(updateAchievementDTO)
}

func (s *AchievementService) DeleteAchievement(id uint) error {
	return s.repo.Delete(id)
}
