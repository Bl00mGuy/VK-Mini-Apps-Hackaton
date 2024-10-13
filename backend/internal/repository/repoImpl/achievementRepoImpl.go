package repoImpl

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository"
	"gorm.io/gorm"
)

type achievementRepository struct {
	db *gorm.DB
	mp *mapper.AchievementMapper
}

func NewAchievementRepository(db *gorm.DB, mp *mapper.AchievementMapper) repository.AchievementRepository {
	return &achievementRepository{
		db: db,
		mp: mp,
	}
}

func (r *achievementRepository) Create(createAchievementDTO *dto.CreateAchievementDTO) error {
	achievement := &entity.Achievement{
		Title:       createAchievementDTO.Title,
		Description: createAchievementDTO.Description,
		UserID:      createAchievementDTO.UserID,
	}

	if err := r.db.Create(achievement).Error; err != nil {
		return err
	}

	return r.LinkUserAchievement(createAchievementDTO.UserID, achievement.ID)
}

func (r *achievementRepository) FindByID(id uint) (*dto.AchievementDTO, error) {
	var achievement entity.Achievement
	if err := r.db.First(&achievement, id).Error; err != nil {
		return nil, err
	}
	return r.mp.ConvertToDTO(&achievement), nil
}

func (r *achievementRepository) FindAll(userID uint) ([]dto.AchievementDTO, error) {
	var achievements []entity.Achievement
	if err := r.db.Where("user_id = ?", userID).Find(&achievements).Error; err != nil {
		return nil, err
	}
	return r.mp.ConvertToSliceDTO(achievements), nil
}

func (r *achievementRepository) Update(updateAchievementDTO *dto.UpdateAchievementDTO) error {
	var achievement entity.Achievement
	if err := r.db.First(&achievement, updateAchievementDTO.AchievementID).Error; err != nil {
		return err
	}

	achievement.Title = updateAchievementDTO.Title
	achievement.Description = updateAchievementDTO.Description
	achievement.UserID = updateAchievementDTO.UserID

	return r.db.Save(&achievement).Error
}

func (r *achievementRepository) Delete(id uint) error {
	r.db.Where("achievement_id = ?", id).Delete(&entity.UserAchievement{})
	return r.db.Delete(&entity.Achievement{}, id).Error
}

func (r *achievementRepository) LinkUserAchievement(userID uint, achievementID uint) error {
	userAchievement := entity.UserAchievement{
		UserID:        userID,
		AchievementID: achievementID,
	}
	return r.db.Create(&userAchievement).Error
}
