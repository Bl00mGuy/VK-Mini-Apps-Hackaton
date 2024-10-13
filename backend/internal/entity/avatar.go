package entity

type Avatar struct {
	ID        uint     `gorm:"primaryKey"`
	Name      string   `gorm:"size:255"`
	ImageURL  string   `gorm:"not null;size:255"`
	Interests []string `gorm:"type:text[]"`
}

func (Avatar) TableName() string {
	return "avatars"
}
