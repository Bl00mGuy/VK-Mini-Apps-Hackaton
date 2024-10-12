package repository

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
)

type AvatarRepository interface {
	Create(createAvatarID *dto.CreateAvatarDTO) error
	FindByID(avatarID uint) (*dto.AvatarDTO, error)
	FindAll(userID uint) ([]dto.AvatarDTO, error)
	Update(updateAvatarDTO *dto.UpdateAvatarDTO) error
	Delete(deleteAvatar *dto.DeleteAvatarDTO) error
}
