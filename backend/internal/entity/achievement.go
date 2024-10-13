package entity

type Achievement struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
}

func (Achievement) TableName() string {
	return "achievements"
}
