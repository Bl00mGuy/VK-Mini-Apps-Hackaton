package services

import "github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"

type AvatarServiceInterface interface {
	CreateAvatar(createAvatarDTO *dto.CreateAvatarDTO) error
	FindAvatarByID(id uint) (*dto.AvatarDTO, error)
	FindAllAvatars(userID uint) ([]dto.AvatarDTO, error)
	UpdateAvatar(updateAvatarDTO *dto.UpdateAvatarDTO) error
	DeleteAvatar(id uint) error
}
