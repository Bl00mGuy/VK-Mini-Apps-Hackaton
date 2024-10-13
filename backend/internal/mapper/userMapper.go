package mapper

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
)

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (a *UserMapper) ConvertToSliceDTO(users []entity.User) []dto.UserDTO {
	arr := make([]dto.UserDTO, 0, len(users))
	for _, user := range users {
		UserDTO := a.ConvertToDTO(&user)
		arr = append(arr, *UserDTO)
	}
	return arr
}

func (a *UserMapper) ConvertToDTO(user *entity.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:       user.ID,
		Dep:      user.Dep,
		Lvl:      user.Lvl,
		Course:   user.Course,
		Sport:    user.Sport,
		Club:     user.Club,
		Mer:      user.MER,
		AvatarID: user.AvatarID,
	}
}
