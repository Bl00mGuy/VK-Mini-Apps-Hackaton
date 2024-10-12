package dto

type AchievementDTO struct {
	AchievementID uint   `json:"achievement_id" binding:"required"`
	UserID        uint   `json:"user_id" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description" binding:"required"`
}

type CreateAchievementDTO struct {
	UserID      uint   `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type FindAchievementDTO struct {
	AchievementID uint `json:"achievement_id" binding:"required"`
}

type FindAllAchievementsDTO struct {
	UserID uint `json:"user_id" binding:"required"`
}

type UpdateAchievementDTO struct {
	AchievementID uint   `json:"achievement_id" binding:"required"`
	UserID        uint   `json:"user_id" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description" binding:"required"`
}

type DeleteAchievementDTO struct {
	AchievementID uint `json:"achievement_id" binding:"required"`
}
