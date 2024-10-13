package impl

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository"
	"gorm.io/gorm"
)

type avatarRepository struct {
	db *gorm.DB
	mp *mapper.AvatarMapper
}

func NewAvatarRepository(db *gorm.DB, mp *mapper.AvatarMapper) repository.AvatarRepository {
	return &avatarRepository{db, mp}
}

func (r *avatarRepository) Create(createAvatarDTO *dto.CreateAvatarDTO) error {
	avatar := &entity.Avatar{
		UserID:    createAvatarDTO.UserID,
		Name:      createAvatarDTO.Name,
		ImageURL:  createAvatarDTO.ImageURL,
		Interests: createAvatarDTO.Interests,
	}
	return r.db.Create(avatar).Error
}

func (r *avatarRepository) FindByID(avatarID uint) (*dto.AvatarDTO, error) {
	var avatar entity.Avatar
	if err := r.db.First(&avatar, avatarID).Error; err != nil {
		return nil, err
	}

	avatarDTO := r.mp.ConvertToAvatarDTO(avatar)
	return &avatarDTO, nil
}

func (r *avatarRepository) FindAll(userID uint) ([]dto.AvatarDTO, error) {
	var avatars []entity.Avatar
	if err := r.db.Where("user_id = ?", userID).Find(&avatars).Error; err != nil {
		return nil, err
	}

	avatarDTOs := r.mp.ConvertToAvatarDTOs(avatars)
	return avatarDTOs, nil
}

func (r *avatarRepository) Update(updateAvatarDTO *dto.UpdateAvatarDTO) error {
	var avatar entity.Avatar
	if err := r.db.First(&avatar, updateAvatarDTO.AvatarID).Error; err != nil {
		return err
	}

	avatar.UserID = updateAvatarDTO.UserID
	avatar.Name = updateAvatarDTO.Name
	avatar.ImageURL = updateAvatarDTO.ImageURL
	avatar.Interests = updateAvatarDTO.Interests

	return r.db.Save(avatar).Error
}

func (r *avatarRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Task{}, id).Error
}
