package mapper

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
)

func ConvertToAvatarDTOs(avatars []entity.Avatar) []dto.AvatarDTO {
	var avatarDTOs []dto.AvatarDTO

	for i, avatar := range avatars {
		avatarDTOs[i] = ConvertToAvatarDTO(avatar)
	}
	return avatarDTOs
}

func ConvertToAvatarDTO(avatar entity.Avatar) dto.AvatarDTO {
	return dto.AvatarDTO{
		AvatarID:  avatar.ID,
		UserID:    avatar.UserID,
		Name:      avatar.Name,
		ImageURL:  avatar.ImageURL,
		Interests: avatar.Interests,
	}
}
