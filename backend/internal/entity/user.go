package entity

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Dep      string `gorm:"not null"`
	Lvl      string `gorm:"not null"`
	Course   int    `gorm:"not null"`
	Sport    string
	Club     string
	MER      string
	AvatarID uint   `gorm:"index"`
	Avatar   Avatar `gorm:"foreignKey:AvatarID;references:ID;constraint:OnDelete:SET NULL"`
}

func (User) TableName() string {
	return "users"
}
