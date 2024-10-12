package impl

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository"
	"gorm.io/gorm"
)

type achievementRepository struct {
	db *gorm.DB
}

func NewAchievementRepository(db *gorm.DB) repository.AchievementRepository {
	return &achievementRepository{db}
}

func (r *achievementRepository) Create(createAchievementDTO *dto.CreateAchievementDTO) error {
	achievement := &entity.Achievement{
		UserID:      createAchievementDTO.UserID,
		Title:       createAchievementDTO.Title,
		Description: createAchievementDTO.Description,
	}
	return r.db.Create(achievement).Error
}

func (r *achievementRepository) FindByID(achievementID uint) (*dto.AchievementDTO, error) {
	var achievement entity.Achievement
	if err := r.db.First(&achievement, achievementID).Error; err != nil {
		return nil, err
	}

	achievementDTO := mapper.ConvertToAchievementDTO(achievement)
	return &achievementDTO, nil
}

func (r *achievementRepository) FindAll(userID uint) ([]dto.AchievementDTO, error) {
	var achievement []entity.Achievement
	if err := r.db.Where("user_id = ?", userID).Find(&achievement).Error; err != nil {
		return nil, err
	}

	achievementDTOs := mapper.ConvertToAchievementDTOs(achievement)
	return achievementDTOs, nil
}

func (r *achievementRepository) Update(updateAchievementDTO *dto.UpdateAchievementDTO) error {
	var achievement entity.Achievement
	if err := r.db.First(&achievement, updateAchievementDTO.AchievementID).Error; err != nil {
		return err
	}

	achievement.Title = updateAchievementDTO.Title
	achievement.Description = updateAchievementDTO.Description
	achievement.UserID = updateAchievementDTO.UserID

	return r.db.Save(achievement).Error
}

func (r *achievementRepository) Delete(deleteAchievementDTO *dto.DeleteAchievementDTO) error {
	return r.db.Delete(&entity.Task{}, deleteAchievementDTO.AchievementID).Error
}