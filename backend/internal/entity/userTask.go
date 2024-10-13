package entity

type UserTask struct {
	UserID uint `gorm:"primaryKey;autoIncrement:false"`
	TaskID uint `gorm:"primaryKey;autoIncrement:false"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Task Task `gorm:"foreignKey:TaskID;references:ID;constraint:OnDelete:CASCADE"`
}

func (UserTask) TableName() string {
	return "users_tasks"
}
