package mapper

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
)

type AvatarMapper struct{}

func NewAvatarMapper() *AvatarMapper {
	return &AvatarMapper{}
}

func (a *AvatarMapper) ConvertToAvatarDTOs(avatars []entity.Avatar) []dto.AvatarDTO {
	var avatarDTOs []dto.AvatarDTO

	for i, avatar := range avatars {
		avatarDTOs[i] = a.ConvertToAvatarDTO(avatar)
	}
	return avatarDTOs
}

func (a *AvatarMapper) ConvertToAvatarDTO(avatar entity.Avatar) dto.AvatarDTO {
	return dto.AvatarDTO{
		AvatarID:  avatar.ID,
		Name:      avatar.Name,
		ImageURL:  avatar.ImageURL,
		Interests: avatar.Interests,
	}
}
