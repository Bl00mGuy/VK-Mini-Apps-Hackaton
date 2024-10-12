package entity

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	IsCompleted bool   `gorm:"not null"`
}
