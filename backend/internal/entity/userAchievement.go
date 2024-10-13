package entity

type UserAchievement struct {
	UserID        uint `gorm:"primaryKey;autoIncrement:false"`
	AchievementID uint `gorm:"primaryKey;autoIncrement:false"`

	User        User        `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Achievement Achievement `gorm:"foreignKey:AchievementID;references:ID;constraint:OnDelete:CASCADE"`
}

func (UserAchievement) TableName() string {
	return "users_achievements"
}
